package bmi

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (b *bmi) DeleteByID(ctx context.Context, id int) error {
	if err := b.bmiRepo.DeleteByID(ctx, id); err != nil {
		logrus.Errorf("delete by id:%d error:%s", id, err.Error())
		return err
	}
	return nil
}
