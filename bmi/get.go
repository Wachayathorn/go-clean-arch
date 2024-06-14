package bmi

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/sirupsen/logrus"
)

func (b *bmi) Get(ctx context.Context) (domain.BmisResponse, error) {
	bmis, err := b.bmiRepo.Get(ctx)
	if err != nil {
		logrus.Errorf("get bmis error:%s", err.Error())
		return nil, err
	}

	result := make([]domain.BmiResponse, len(bmis))
	for i, bmi := range bmis {
		res, err := bmi.ToResponse()
		if err != nil {
			logrus.Errorf("build bmi res:%v error:%s", bmi, err.Error())
			return nil, err
		}
		result[i] = res
	}

	return result, nil
}
