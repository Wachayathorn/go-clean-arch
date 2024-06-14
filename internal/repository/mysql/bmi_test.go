package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_bmi_Create(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare("INSERT INTO bmi \\(name, weight, height, bmi, created_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\)")
		mock.ExpectExec("INSERT INTO bmi \\(name, weight, height, bmi, created_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\)").
			WithArgs("name", "70.0", "175", "22.86", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		insertStmt, err := db.Prepare("INSERT INTO bmi (name, weight, height, bmi, created_at) VALUES (?, ?, ?, ?, ?)")
		assert.NoError(t, err)

		bmiRepo := &bmi{insertStmt: insertStmt}

		bmiRequest := domain.BmiRepository{
			Name:   "name",
			Weight: "70.0",
			Height: "175",
			Bmi:    "22.86",
		}

		res, err := bmiRepo.Create(context.Background(), bmiRequest)

		assert.NoError(t, err)
		assert.Equal(t, 1, res.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_bmi_Get(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedBmis := []domain.BmiRepository{
			{ID: 1, Name: "name", Weight: "70.0", Height: "175", Bmi: "22.86", CreatedAt: nil, UpdatedAt: nil},
		}
		rows := sqlmock.NewRows([]string{"id", "name", "weight", "height", "bmi", "created_at", "updated_at"}).
			AddRow(expectedBmis[0].ID, expectedBmis[0].Name, expectedBmis[0].Weight, expectedBmis[0].Height, expectedBmis[0].Bmi, nil, nil)
		mock.ExpectQuery("SELECT * FROM bmi").WillReturnRows(rows)
	})
}

func Test_bmi_GetByID(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		expectedBmis := []domain.BmiRepository{
			{ID: 1, Name: "name", Weight: "70.0", Height: "175", Bmi: "22.86"},
		}
		rows := sqlmock.NewRows([]string{"id", "name", "weight", "height", "bmi", "created_at", "updated_at"}).
			AddRow(expectedBmis[0].ID, expectedBmis[0].Name, expectedBmis[0].Weight, expectedBmis[0].Height, expectedBmis[0].Bmi, nil, nil)
		mock.ExpectQuery(getBmisByIDSql).WillReturnRows(rows)
	})
}

func Test_bmi_UpdateByID(t *testing.T) {
	t.Run("succeed", func(t *testing.T) {
		timeNow := time.Now()
		bmiToUpdate := domain.BmiRepository{
			ID:        1,
			Name:      "John Doe",
			Weight:    "75.0",
			Height:    "180",
			Bmi:       "23.15",
			UpdatedAt: &timeNow,
		}

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare("UPDATE bmi SET name=\\?, weight=\\?, height=\\?, bmi=\\?, updated_at=\\? WHERE id=\\?")
		mock.ExpectExec("UPDATE bmi SET name=\\?, weight=\\?, height=\\?, bmi=\\?, updated_at=\\? WHERE id=\\?").
			WithArgs(bmiToUpdate.Name, bmiToUpdate.Weight, bmiToUpdate.Height, bmiToUpdate.Bmi, bmiToUpdate.UpdatedAt, bmiToUpdate.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		updateStmt, err := db.Prepare("UPDATE bmi SET name=?, weight=?, height=?, bmi=?, updated_at=? WHERE id=?")
		assert.NoError(t, err)

		bmiRepo := &bmi{updateByIDStmt: updateStmt}

		updatedBmi, err := bmiRepo.UpdateByID(context.Background(), bmiToUpdate)

		assert.NoError(t, err)
		assert.Equal(t, bmiToUpdate, updatedBmi)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_bmi_DeleteByID(t *testing.T) {
	t.Run("successful deletion", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(deleteBmiByIDsql)
		mock.ExpectExec(deleteBmiByIDsql).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		deleteStmt, err := db.Prepare(deleteBmiByIDsql)
		assert.NoError(t, err)

		bmiRepo := &bmi{deleteByIDStmt: deleteStmt}

		err = bmiRepo.DeleteByID(context.Background(), 1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
