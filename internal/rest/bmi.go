package rest

import (
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type BmiHandler struct {
	Service bmi.Bmi
}

func NewBmiHandler(e *echo.Echo, svc bmi.Bmi) {
	handler := &BmiHandler{
		Service: svc,
	}
	bmiApi := e.Group("/api/bmi")
	bmiApi.POST("", handler.Create)
	bmiApi.GET("", handler.Get)
	bmiApi.GET("/:id", handler.GetByID)
	bmiApi.PUT("/:id", handler.UpdateByID)
	bmiApi.DELETE("/:id", handler.DeleteByID)
}

// Create godoc
// @Summary Create BMI
// @Description Create BMI
// @Accept json
// @Produce json
// @Param body body domain.BmiRequest true "BMI request object"
// @Success 201 {object} domain.BmiResponse
// @Failure 400 {object} ResponseError
// @Router /api/bmi [post]
func (b *BmiHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var bmiRequest domain.BmiRequest
	if err := c.Bind(&bmiRequest); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	res, err := b.Service.Create(ctx, bmiRequest)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

// Get godoc
// @Summary Get BMIs
// @Description Get BMIs
// @Produce json
// @Success 200 {object} []domain.BmiResponse
// @Router /api/bmi [get]
func (b *BmiHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := b.Service.Get(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

// GetByID godoc
// @Summary Get BMI by ID
// @Description Get BMI by ID
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} domain.BmiResponse
// @Failure 400 {object} ResponseError
// @Router /api/bmi/{id} [get]
func (b *BmiHandler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "id type invalid"})
	}

	res, err := b.Service.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

// UpdateByID godoc
// @Summary Update BMI by ID
// @Description Update BMI by ID
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param body body domain.BmiRequest true "BMI request object"
// @Success 200 {object} domain.BmiResponse
// @Failure 400 {object} ResponseError
// @Router /api/bmi/{id} [put]
func (b *BmiHandler) UpdateByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "id type invalid"})
	}

	var bmiRequest domain.BmiRequest
	if err := c.Bind(&bmiRequest); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	res, err := b.Service.UpdateByID(ctx, id, bmiRequest)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteByID godoc
// @Summary Delete BMI by ID
// @Description Delete BMI by ID
// @Param id path int true "ID"
// @Success 204 "No Content"
// @Failure 400 {object} ResponseError
// @Router /api/bmi/{id} [delete]
func (b *BmiHandler) DeleteByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "id type invalid"})
	}

	if err := b.Service.DeleteByID(ctx, id); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
