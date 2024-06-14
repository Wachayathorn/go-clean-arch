package mysql

const (
	insertBmiSql     = `INSERT INTO bmi(name, weight, height, bmi, created_at) VALUES(?,?,?,?,?)`
	getBmisSql       = `SELECT * FROM bmi`
	getBmisByIDSql   = `SELECT * FROM bmi WHERE id=?`
	updateBmiByIDsql = `UPDATE bmi set name=?, weight=?, height=?, bmi=?, updated_at=? WHERE id=?`
	deleteBmiByIDsql = `DELETE FROM bmi WHERE id = ?`
)
