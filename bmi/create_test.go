package bmi

import (
	"context"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/repository/mysql"
	"github.com/bxcodec/go-clean-arch/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bmi_Create(t *testing.T) {
	timeNow := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)

	type fields struct {
		bmiRepo mysql.Bmi
	}
	type args struct {
		ctx context.Context
		bmi domain.BmiRequest
	}

	name := "name"
	w := float64(70)
	h := float64(70)

	validArgs := args{
		ctx: context.Background(),
		bmi: domain.BmiRequest{
			Name:   name,
			Weight: w,
			Height: h,
		},
	}

	bmiRepoReq := domain.BmiRepository{
		Name:      name,
		Weight:    strconv.FormatFloat(w, 'f', -1, 64),
		Height:    strconv.FormatFloat(h, 'f', -1, 64),
		Bmi:       "142.86",
		CreatedAt: &timeNow,
	}

	bmiRepoRes := domain.BmiRepository{
		ID:        1,
		Name:      name,
		Weight:    strconv.FormatFloat(w, 'f', -1, 64),
		Height:    strconv.FormatFloat(h, 'f', -1, 64),
		Bmi:       "142.86",
		CreatedAt: &timeNow,
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
			name: "create_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().Create(mock.Anything, bmiRepoReq).Return(domain.BmiRepository{}, errors.New("error")).Once()
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
					m.EXPECT().Create(mock.Anything, bmiRepoReq).Return(domain.BmiRepository{}, nil).Once()
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
					m.EXPECT().Create(mock.Anything, bmiRepoReq).Return(bmiRepoRes, nil).Once()
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
				time: func() util.TimeUtil {
					m := util.NewMockTimeUtil(t)
					m.EXPECT().Now().Return(timeNow).Once()
					return m
				}(),
			}
			got, err := b.Create(tt.args.ctx, tt.args.bmi)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}
