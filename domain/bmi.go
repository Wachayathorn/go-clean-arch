package domain

import (
	"time"

	"github.com/bxcodec/go-clean-arch/util"
)

type BmiRepository struct {
	ID        int
	Name      string
	Height    string
	Weight    string
	Bmi       string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (b *BmiRepository) ToResponse() (BmiResponse, error) {
	w, e := util.StringToFloat64(b.Weight)
	if e != nil {
		return BmiResponse{}, e
	}
	h, e := util.StringToFloat64(b.Height)
	if e != nil {
		return BmiResponse{}, e
	}
	bmi, e := util.StringToFloat64(b.Bmi)
	if e != nil {
		return BmiResponse{}, e
	}
	return BmiResponse{
		ID:        b.ID,
		Name:      b.Name,
		Weight:    w,
		Height:    h,
		Bmi:       bmi,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}, nil
}

type BmiRequest struct {
	Name   string  `json:"name" binding:"required"`
	Weight float64 `json:"weight" binding:"required"`
	Height float64 `json:"height" binding:"required"`
}

type BmisResponse []BmiResponse

type BmiResponse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Weight    float64    `json:"weight"`
	Height    float64    `json:"height"`
	Bmi       float64    `json:"bmi"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
