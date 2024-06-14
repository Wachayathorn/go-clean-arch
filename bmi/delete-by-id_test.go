package bmi

import (
	"context"
	"errors"
	"testing"

	"github.com/bxcodec/go-clean-arch/internal/repository/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bmi_DeleteByID(t *testing.T) {
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

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete_fail",
			args: validArgs,
			fields: fields{
				bmiRepo: func() mysql.Bmi {
					m := mysql.NewMockBmi(t)
					m.EXPECT().DeleteByID(mock.Anything, id).Return(errors.New("error")).Once()
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
					m.EXPECT().DeleteByID(mock.Anything, id).Return(nil).Once()
					return m
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bmi{
				bmiRepo: tt.fields.bmiRepo,
			}
			if err := b.DeleteByID(tt.args.ctx, tt.args.id); err != nil {
				assert.Error(t, err)
			}
		})
	}
}
