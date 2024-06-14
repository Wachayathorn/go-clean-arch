package bmi

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (b *bmi) GetByID(ctx context.Context, id int) (domain.BmiResponse, error) {
	bmi, err := b.bmiRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Errorf("get by id:%d error:%s", id, err.Error())
		return domain.BmiResponse{}, err
	}
	res, err := bmi.ToResponse()
	if err != nil {
		logrus.Errorf("build bmi res:%v error:%s", bmi, err.Error())
		return domain.BmiResponse{}, err
	}
	return res, nil
}
