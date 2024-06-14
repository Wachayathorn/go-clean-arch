package bmi

import (
	"context"
	"errors"
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/repository/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bmi_Get(t *testing.T) {
	type fields struct {
		bmiRepo mysql.Bmi
	}
	type args struct {
		ctx context.Context
	}

	validArgs := args{
		ctx: context.Background(),
	}

	id := 1
	bmiRepoRes := domain.BmiRepository{
		ID:     id,
		Name:   "name",
		Weight: "70",
		Height: "70",
		Bmi:    "142.86",
	}
	res, _ := bmiRepoRes.ToResponse()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.BmisResponse
		wantErr bool
	}{
		{
			name: "get_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().Get(mock.Anything).Return([]domain.BmiRepository{}, errors.New("error")).Once()
					return m
				}(),
			},
			wantErr: true,
		},
		{
			name: "get_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().Get(mock.Anything).Return([]domain.BmiRepository{
						{},
					}, nil).Once()
					return m
				}(),
			},
			wantErr: true,
		},
		{
			name: "succeed",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().Get(mock.Anything).Return([]domain.BmiRepository{
						bmiRepoRes,
					}, nil).Once()
					return m
				}(),
			},
			want: domain.BmisResponse{
				res,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bmi{
				bmiRepo: tt.fields.bmiRepo,
			}
			got, err := b.Get(tt.args.ctx)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}
