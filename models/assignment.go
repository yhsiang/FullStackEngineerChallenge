package models

import (
	"time"

	"github.com/yhsiang/review360/database"
)

type Assignment struct {
	ID        int64     `db:"id"`
	Reviewee  int64     `db:"reviewee" json:"reviewee"`
	Reviewer  int64     `db:"reviewer" json:"reviewer"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (a Assignment) FindAssignID(db database.DB) (int64, error) {
	sqlStatement := `SELECT id FROM review_assignments WHERE reviewee = ? AND reviewer = ?`
	err := db.QueryRow(sqlStatement, a.Reviewee, a.Reviewer).Scan(&a.ID)
	if err != nil {
		return 0, err
	}
	return a.ID, nil
}

func (a Assignment) Save(db database.DB) (Assignment, error) {
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	sqlStatement := `INSERT INTO review_assignments (reviewee, reviewer, updated_at, created_at) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(sqlStatement, a.Reviewee, a.Reviewer, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return a, err
	}
	a.ID, err = result.LastInsertId()
	if err != nil {
		return a, err
	}

	return a, nil
}

func (a Assignment) Remove(db database.DB) error {
	sqlStatement := `DELETE FROM review_assignments WHERE reviewee = ? AND reviewer = ?`
	_, err := db.Exec(sqlStatement, a.Reviewee, a.Reviewer)
	if err != nil {
		return err
	}

	return nil
}
