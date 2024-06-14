package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/bxcodec/go-clean-arch/domain"
)

type Bmi interface {
	Create(ctx context.Context, bmi domain.BmiRepository) (domain.BmiRepository, error)
	Get(ctx context.Context) ([]domain.BmiRepository, error)
	GetByID(ctx context.Context, id int) (domain.BmiRepository, error)
	UpdateByID(ctx context.Context, bmi domain.BmiRepository) (domain.BmiRepository, error)
	DeleteByID(ctx context.Context, id int) error
}

type bmi struct {
	insertStmt     *sql.Stmt
	getStmt        *sql.Stmt
	getByIDStmt    *sql.Stmt
	updateByIDStmt *sql.Stmt
	deleteByIDStmt *sql.Stmt
}

func NewBmiRepository(db *sql.DB) Bmi {
	insertStmt, err := db.Prepare(insertBmiSql)
	if err != nil {
		log.Fatalf("Error preparing statement:%s error:%v", insertBmiSql, err)
	}

	getStmt, err := db.Prepare(getBmisSql)
	if err != nil {
		log.Fatalf("Error preparing statement:%s error:%v", getBmisSql, err)
	}

	getByIDStmt, err := db.Prepare(getBmisByIDSql)
	if err != nil {
		log.Fatalf("Error preparing statement:%s error:%v", getBmisByIDSql, err)
	}

	updateByIDStmt, err := db.Prepare(updateBmiByIDsql)
	if err != nil {
		log.Fatalf("Error preparing statement:%s error:%v", updateBmiByIDsql, err)
	}

	deleteByIDStmt, err := db.Prepare(deleteBmiByIDsql)
	if err != nil {
		log.Fatalf("Error preparing statement:%s error:%v", deleteBmiByIDsql, err)
	}

	return &bmi{
		insertStmt,
		getStmt,
		getByIDStmt,
		updateByIDStmt,
		deleteByIDStmt,
	}
}

func (b *bmi) Create(ctx context.Context, bmi domain.BmiRepository) (domain.BmiRepository, error) {
	res, err := b.insertStmt.ExecContext(ctx, bmi.Name, bmi.Weight, bmi.Height, bmi.Bmi, bmi.CreatedAt)
	if err != nil {
		return domain.BmiRepository{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return domain.BmiRepository{}, err
	}
	bmi.ID = int(lastID)

	return bmi, nil
}

func (b *bmi) Get(ctx context.Context) ([]domain.BmiRepository, error) {
	bmis := []domain.BmiRepository{}

	rows, err := b.getStmt.QueryContext(ctx)
	if err != nil {
		return bmis, err
	}
	defer rows.Close()

	for rows.Next() {
		bmi := domain.BmiRepository{}
		if err := rows.Scan(&bmi.ID, &bmi.Name, &bmi.Weight, &bmi.Height, &bmi.Bmi, &bmi.CreatedAt, &bmi.UpdatedAt); err != nil {
			return nil, err
		}
		bmis = append(bmis, bmi)
	}

	return bmis, nil
}

func (b *bmi) GetByID(ctx context.Context, id int) (domain.BmiRepository, error) {
	rows, err := b.getByIDStmt.QueryContext(ctx, id)
	if err != nil {
		return domain.BmiRepository{}, err
	}
	defer rows.Close()

	bmi := domain.BmiRepository{}
	if rows.Next() {
		if err := rows.Scan(&bmi.ID, &bmi.Name, &bmi.Weight, &bmi.Height, &bmi.Bmi, &bmi.CreatedAt, &bmi.UpdatedAt); err != nil {
			return domain.BmiRepository{}, err
		}
	} else {
		return domain.BmiRepository{}, sql.ErrNoRows
	}

	return bmi, nil
}

func (b *bmi) UpdateByID(ctx context.Context, bmi domain.BmiRepository) (domain.BmiRepository, error) {
	res, err := b.updateByIDStmt.ExecContext(ctx, bmi.Name, bmi.Weight, bmi.Height, bmi.Bmi, bmi.UpdatedAt, bmi.ID)
	if err != nil {
		return domain.BmiRepository{}, nil
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return domain.BmiRepository{}, nil
	}

	if affect != 1 {
		return domain.BmiRepository{}, fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
	}

	return bmi, nil
}

func (b *bmi) DeleteByID(ctx context.Context, id int) error {
	res, err := b.deleteByIDStmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		return fmt.Errorf("weird  Behavior. Total Affected: %d", rowsAfected)
	}

	return nil
}
