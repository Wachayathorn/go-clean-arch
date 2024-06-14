package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		e := echo.New()
		mockService := bmi.NewMockBmi(t)
		bmiRequest := domain.BmiRequest{Weight: 70, Height: 175}
		bmiResponse := domain.BmiResponse{ID: 1, Weight: 70, Height: 175, Bmi: 22.86}
		mockService.EXPECT().Create(mock.Anything, bmiRequest).Return(bmiResponse, nil).Once()

		reqBody, _ := json.Marshal(bmiRequest)
		req := httptest.NewRequest(http.MethodPost, "/bmi", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := &BmiHandler{Service: mockService}
		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		var res domain.BmiResponse
		json.Unmarshal(rec.Body.Bytes(), &res)
		assert.Equal(t, bmiResponse, res)
	})
}

func TestGet(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		e := echo.New()
		mockService := bmi.NewMockBmi(t)
		bmiResponses := []domain.BmiResponse{
			{ID: 1, Weight: 70.0, Height: 175, Bmi: 22.86},
		}
		mockService.EXPECT().Get(mock.Anything).Return(bmiResponses, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/bmi", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := &BmiHandler{Service: mockService}
		err := handler.Get(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var res []domain.BmiResponse
		json.Unmarshal(rec.Body.Bytes(), &res)
		assert.Equal(t, bmiResponses, res)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		e := echo.New()
		mockService := bmi.NewMockBmi(t)
		bmiResponse := domain.BmiResponse{ID: 1, Weight: 70.0, Height: 175, Bmi: 22.86}
		mockService.EXPECT().GetByID(mock.Anything, 1).Return(bmiResponse, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/bmi/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		handler := &BmiHandler{Service: mockService}
		err := handler.GetByID(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var res domain.BmiResponse
		json.Unmarshal(rec.Body.Bytes(), &res)
		assert.Equal(t, bmiResponse, res)
	})
}

func TestUpdateByID(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		e := echo.New()
		mockService := bmi.NewMockBmi(t)
		bmiRequest := domain.BmiRequest{Weight: 70.0, Height: 175}
		bmiResponse := domain.BmiResponse{ID: 1, Weight: 70.0, Height: 175, Bmi: 22.86}
		mockService.EXPECT().UpdateByID(mock.Anything, 1, bmiRequest).Return(bmiResponse, nil).Once()

		reqBody, _ := json.Marshal(bmiRequest)
		req := httptest.NewRequest(http.MethodPut, "/bmi/1", bytes.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		handler := &BmiHandler{Service: mockService}
		err := handler.UpdateByID(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		var res domain.BmiResponse
		json.Unmarshal(rec.Body.Bytes(), &res)
		assert.Equal(t, bmiResponse, res)
	})
}

func TestDeleteByID(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		e := echo.New()
		mockService := bmi.NewMockBmi(t)
		mockService.EXPECT().DeleteByID(mock.Anything, 1).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/bmi/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		handler := &BmiHandler{Service: mockService}
		err := handler.DeleteByID(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})
}
