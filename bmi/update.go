package bmi

import (
	"context"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (b *bmi) UpdateByID(ctx context.Context, id int, req domain.BmiRequest) (domain.BmiResponse, error) {
	bmi, err := b.bmiRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Errorf("get bmi id:%d error:%s", id, err.Error())
		return domain.BmiResponse{}, err
	}

	now := b.time.Now()
	bmi.Name = req.Name
	bmi.Weight = strconv.FormatFloat(req.Weight, 'f', -1, 64)
	bmi.Height = strconv.FormatFloat(req.Height, 'f', -1, 64)
	bmi.Bmi = b.CalculateBmi(req.Weight, req.Height)
	bmi.UpdatedAt = &now

	updateResult, err := b.bmiRepo.UpdateByID(ctx, bmi)
	if err != nil {
		logrus.Errorf("update bmi id:%d error:%s", id, err.Error())
		return domain.BmiResponse{}, err
	}

	res, err := updateResult.ToResponse()
	if err != nil {
		logrus.Errorf("build bmi res:%v error:%s", updateResult, err.Error())
		return domain.BmiResponse{}, err
	}

	return res, err
}
