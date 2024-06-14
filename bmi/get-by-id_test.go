package bmi

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/repository/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bmi_GetByID(t *testing.T) {
	timeNow := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	type fields struct {
		bmiRepo mysql.Bmi
	}
	type args struct {
		ctx context.Context
		id  int
	}

	id := 1
	validArgs := args{
		ctx: context.Background(),
		id:  id,
	}

	bmiRepoRes := domain.BmiRepository{
		ID:        id,
		Name:      "name",
		Weight:    "70",
		Height:    "70",
		Bmi:       "142.86",
		CreatedAt: &timeNow,
		UpdatedAt: &timeNow,
	}
	res, _ := bmiRepoRes.ToResponse()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.BmiResponse
		wantErr bool
	}{
		{
			name: "get_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().GetByID(mock.Anything, id).Return(domain.BmiRepository{}, errors.New("error")).Once()
					return m
				}(),
			},
			want:    domain.BmiResponse{},
			wantErr: true,
		},
		{
			name: "map_response_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().GetByID(mock.Anything, id).Return(domain.BmiRepository{}, nil).Once()
					return m
				}(),
			},
			want:    domain.BmiResponse{},
			wantErr: true,
		},
		{
			name: "succeed",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().GetByID(mock.Anything, id).Return(bmiRepoRes, nil).Once()
					return m
				}(),
			},
			want: res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bmi{
				bmiRepo: tt.fields.bmiRepo,
			}
			got, err := b.GetByID(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}
