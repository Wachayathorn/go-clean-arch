package bmi

import (
	"context"
	"fmt"
	"math"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/repository/mysql"
	"github.com/bxcodec/go-clean-arch/util"
)

type Bmi interface {
	Create(ctx context.Context, bmi domain.BmiRequest) (domain.BmiResponse, error)
	Get(ctx context.Context) (domain.BmisResponse, error)
	GetByID(ctx context.Context, id int) (domain.BmiResponse, error)
	UpdateByID(ctx context.Context, id int, req domain.BmiRequest) (domain.BmiResponse, error)
	DeleteByID(ctx context.Context, id int) error
}

type bmi struct {
	bmiRepo mysql.Bmi
	time    util.TimeUtil
}

func New(bmiRepo mysql.Bmi) Bmi {
	return &bmi{
		bmiRepo,
		util.NewTimeUtil(),
	}
}

func (b *bmi) CalculateBmi(w, h float64) string {
	heightM := h / 100.0
	bmi := w / math.Pow(heightM, 2)
	return fmt.Sprintf("%.2f", bmi)
}
