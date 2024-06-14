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

func Test_bmi_UpdateByID(t *testing.T) {
	timeNow := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	type fields struct {
		bmiRepo mysql.Bmi
		time    util.TimeUtil
	}
	type args struct {
		ctx context.Context
		id  int
		req domain.BmiRequest
	}

	id := 1
	name := "name"
	w := float64(70)
	h := float64(70)

	validArgs := args{
		ctx: context.Background(),
		id:  id,
		req: domain.BmiRequest{
			Name:   name,
			Weight: w,
			Height: h,
		},
	}

	bmiRepoReq := domain.BmiRepository{
		ID:        id,
		Name:      name,
		Weight:    strconv.FormatFloat(w, 'f', -1, 64),
		Height:    strconv.FormatFloat(h, 'f', -1, 64),
		Bmi:       "142.86",
		CreatedAt: &timeNow,
		UpdatedAt: &timeNow,
	}

	bmiRepoRes := domain.BmiRepository{
		ID:        1,
		Name:      name,
		Weight:    strconv.FormatFloat(w, 'f', -1, 64),
		Height:    strconv.FormatFloat(h, 'f', -1, 64),
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
			name: "update_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().GetByID(mock.Anything, id).Return(bmiRepoRes, nil).Once()
					m.EXPECT().UpdateByID(mock.Anything, bmiRepoReq).Return(domain.BmiRepository{}, errors.New("error")).Once()
					return m
				}(),
				time: func() util.TimeUtil {
					m := util.NewMockTimeUtil(t)
					m.EXPECT().Now().Return(timeNow).Once()
					return m
				}(),
			},
			want:    domain.BmiResponse{},
			wantErr: true,
		},
		{
			name: "map_reponse_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().GetByID(mock.Anything, id).Return(bmiRepoRes, nil).Once()
					m.EXPECT().UpdateByID(mock.Anything, bmiRepoReq).Return(domain.BmiRepository{}, nil).Once()
					return m
				}(),
				time: func() util.TimeUtil {
					m := util.NewMockTimeUtil(t)
					m.EXPECT().Now().Return(timeNow).Once()
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
					m.EXPECT().UpdateByID(mock.Anything, bmiRepoReq).Return(bmiRepoRes, nil).Once()
					return m
				}(),
				time: func() util.TimeUtil {
					m := util.NewMockTimeUtil(t)
					m.EXPECT().Now().Return(timeNow).Once()
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
				time:    tt.fields.time,
			}
			got, err := b.UpdateByID(tt.args.ctx, tt.args.id, tt.args.req)
			assert.Equal(t, tt.want, got)
			if err != nil {
				assert.Error(t, err)
			}
		})
	}
}
