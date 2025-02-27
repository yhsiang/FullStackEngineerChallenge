package models

import (
	"time"

	"github.com/yhsiang/review360/database"
)

type Review struct {
	ID        int64     `db:"id" json:"id"`
	Content   string    `db:"content" json:"content"`
	AssignID  int64     `db:"assign_id" json:"assign_id"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func (r Review) Type() ResponseType {
	return ReviewType
}

type Reviews []Review

func (r Reviews) Type() ResponseType {
	return ReviewsType
}

func (r Review) FindAll(db database.DB) (reviews Reviews, err error) {
	sqlStatement := `select id, content, assign_id from performance_reviews`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return reviews, err
	}

	for rows.Next() {
		var review Review
		err = rows.Scan(&review.ID, &review.Content, &review.AssignID)
		if err != nil {
			return reviews, err
		}
		reviews = append(reviews, review)
	}
	return reviews, err
}

func (r Review) Find(db database.DB) (Review, error) {
	sqlStatement := `select assign_id, content from performance_reviews WHERE id = ?`
	err := db.QueryRow(sqlStatement, r.ID).Scan(&r.AssignID, &r.Content)
	if err != nil {
		return r, err
	}
	return r, nil
}

func (r Review) Save(db database.DB) (Review, error) {
	if r.ID > 0 {
		r.UpdatedAt = time.Now()
		sqlStatement := `UPDATE performance_reviews SET content = ?, updated_at = ? WHERE id = ?`
		_, err := db.Exec(sqlStatement, r.Content, r.UpdatedAt, r.ID)
		if err != nil {
			return r, err
		}
	} else {
		r.UpdatedAt = time.Now()
		r.CreatedAt = time.Now()
		sqlStatement := `INSERT INTO performance_reviews (content, assign_id, updated_at, created_at) VALUES (?, ?, ?, ?)`
		result, err := db.Exec(sqlStatement, r.Content, r.AssignID, r.UpdatedAt, r.CreatedAt)
		if err != nil {
			return r, err
		}
		r.ID, err = result.LastInsertId()
		if err != nil {
			return r, err
		}
	}

	return r, nil
}
