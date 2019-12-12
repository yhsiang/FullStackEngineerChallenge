package models

import (
	"time"

	"github.com/yhsiang/review360/database"
)

type Assignment struct {
	ID        int64     `db:"id"`
	Reviewee  Employee  `db:"reviewee"`
	Reviewer  Employee  `db:"reviewer"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (a Assignment) Save(db database.DB) (Assignment, error) {
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	sqlStatement := `INSERT INTO review_assignments (reviewee, reviewer) VALUES (?, ?)`
	result, err := db.Exec(sqlStatement, a.Reviewee.ID, a.Reviewer.ID)
	if err != nil {
		return a, err
	}
	a.ID, err = result.LastInsertId()
	if err != nil {
		return a, err
	}

	return a, nil
}
