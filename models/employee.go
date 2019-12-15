package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/yhsiang/review360/database"
)

type Employee struct {
	ID        int64      `json:"id" uri:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Reviewers []Reviewer `json:"reviewers"`
	Reviewees []Reviewee `json:"reviewees"`
	UpdatedAt time.Time  `db:"updated_at" json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"-"`
}

type Reviewer struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	AssignID int64  `json:"assign_id"`
}

type Reviewee struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	NullReviewID sql.NullInt64 `json:"-"`
	ReviewID     int64         `json:"review_id"`
}

type Employees []Employee

func (ems Employees) Type() ResponseType {
	return EmployeesType
}

func (e Employee) Type() ResponseType {
	return EmployeeType
}

func (e Employee) FindAll(db database.DB) (Employees, error) {
	var employees = make(Employees, 0)
	rows, err := db.Query("SELECT id, name FROM employees")
	if err != nil {
		return employees, err
	}
	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.ID, &employee.Name)
		if err != nil {
			return employees, err
		}
		reviewers, err := employee.getReviewers(db)
		if err != nil {
			return employees, err
		}
		employee.Reviewers = reviewers
		reviewees, err := employee.getReviewees(db)
		if err != nil {
			return employees, err
		}
		employee.Reviewees = reviewees
		employees = append(employees, employee)
	}
	return employees, nil
}

func (e Employee) Find(db database.DB) (Employee, error) {
	err := db.QueryRow("SELECT name FROM employees WHERE id = ?", e.ID).Scan(&e.Name)
	if err != nil {
		return e, err
	}
	reviewers, err := e.getReviewers(db)
	if err != nil {
		return e, err
	}
	e.Reviewers = reviewers
	reviewees, err := e.getReviewees(db)
	if err != nil {
		return e, err
	}
	e.Reviewees = reviewees
	return e, nil
}

func (e Employee) getReviewers(db database.DB) ([]Reviewer, error) {
	var reviewers = make([]Reviewer, 0)
	rows, err := db.Query("SELECT E.id, E.name, R.id FROM employees AS E, review_assignments AS R WHERE E.id = R.reviewer AND R.reviewee = ?", e.ID)
	if err != nil {
		return reviewers, err
	}
	for rows.Next() {
		var reviewer Reviewer
		err := rows.Scan(&reviewer.ID, &reviewer.Name, &reviewer.AssignID)
		if err != nil {
			return reviewers, err
		}
		reviewers = append(reviewers, reviewer)
	}
	return reviewers, nil
}

func (e Employee) getReviewees(db database.DB) ([]Reviewee, error) {
	var reviewees = make([]Reviewee, 0)
	rows, err := db.Query(`
SELECT employees.id AS id, employees.name AS name, performance_reviews.id AS review_id FROM employees
INNER JOIN review_assignments on employees.id = review_assignments.reviewee
LEFT JOIN performance_reviews on review_assignments.id = performance_reviews.assign_id
WHERE review_assignments.reviewer = ?`, e.ID)
	if err != nil {
		return reviewees, err
	}
	for rows.Next() {
		var reviewee Reviewee
		err := rows.Scan(&reviewee.ID, &reviewee.Name, &reviewee.NullReviewID)
		if err != nil {
			return reviewees, err
		}
		reviewee.ReviewID = reviewee.NullReviewID.Int64
		reviewees = append(reviewees, reviewee)
	}
	return reviewees, nil
}

func (em Employee) Save(db database.DB) (Employee, error) {
	if em.Name == "" {
		return em, errors.New("employee name can not be empty")
	}
	if em.ID > 0 {
		em.UpdatedAt = time.Now()
		sqlStatement := `UPDATE employees SET name = ?, updated_at = ? WHERE id = ?`
		_, err := db.Exec(sqlStatement, em.Name, em.UpdatedAt, em.ID)
		if err != nil {
			return em, err
		}
	} else {
		em.UpdatedAt = time.Now()
		em.CreatedAt = time.Now()
		sqlStatement := `INSERT INTO employees (name, updated_at, created_at) VALUES (?, ?, ?)`
		result, err := db.Exec(sqlStatement, em.Name, em.UpdatedAt, em.CreatedAt)
		if err != nil {
			return em, err
		}
		em.ID, err = result.LastInsertId()
		if err != nil {
			return em, err
		}
	}

	return em, nil
}

func (em Employee) Remove(db database.DB) error {

	if em.ID == 0 {
		return errors.New("employee id should not be zero")
	}

	sqlStatement := `DELETE FROM employees WHERE id = ?`
	_, err := db.Exec(sqlStatement, em.ID)
	if err != nil {
		return err
	}

	return nil
}
