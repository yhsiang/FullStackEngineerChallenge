package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/review360/database"
)

func TestQuery(t *testing.T) {
	db, mock, err := sqlmock.New()

	mockEmployee := sqlmock.NewRows([]string{"name"}).
		AddRow("test1")

	mock.ExpectQuery("SELECT name FROM employees WHERE id = ?").
		WithArgs(1).
		WillReturnRows(mockEmployee)

	mockReviewers := sqlmock.NewRows([]string{"id", "name", "id"}).
		AddRow("2", "test2", "15")

	mock.ExpectQuery("SELECT E.id, E.name, R.id FROM employees AS E, review_assignments AS R WHERE E.id = R.reviewer AND R.reviewee = ?").
		WithArgs(1).
		WillReturnRows(mockReviewers)

	mockReviewees := sqlmock.NewRows([]string{"id", "name", "review_id"}).
		AddRow("2", "test3", "3")
	mock.ExpectQuery(`SELECT employees.id AS id, employees.name AS name, performance_reviews.id AS review_id FROM employees
	INNER JOIN review_assignments on employees.id = review_assignments.reviewee
	LEFT JOIN performance_reviews on review_assignments.id = performance_reviews.assign_id
	WHERE review_assignments.reviewer = ?`).
		WithArgs(1).
		WillReturnRows(mockReviewees)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	em := &Employee{
		ID: 1,
	}
	mockDB := database.NewWithDB(db)
	// now we execute our method
	employee, err := em.Find(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, "test1", employee.Name)
	assert.Equal(t, "test2", employee.Reviewers[0].Name)
	assert.Equal(t, "test3", employee.Reviewees[0].Name)
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
