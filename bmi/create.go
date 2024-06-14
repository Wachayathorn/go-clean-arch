package bmi

import (
	"context"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (b *bmi) Create(ctx context.Context, bmi domain.BmiRequest) (domain.BmiResponse, error) {
	now := b.time.Now()
	createResult, err := b.bmiRepo.Create(ctx, domain.BmiRepository{
		Name:      bmi.Name,
		Weight:    strconv.FormatFloat(bmi.Weight, 'f', -1, 64),
		Height:    strconv.FormatFloat(bmi.Height, 'f', -1, 64),
		Bmi:       b.CalculateBmi(bmi.Weight, bmi.Height),
		CreatedAt: &now,
	})
	if err != nil {
		logrus.Errorf("create bmi req:%v error:%s", bmi, err.Error())
		return domain.BmiResponse{}, err
	}
	res, err := createResult.ToResponse()
	if err != nil {
		logrus.Errorf("build bmi res:%v error:%s", createResult, err.Error())
		return domain.BmiResponse{}, err
	}
	return res, nil
}
