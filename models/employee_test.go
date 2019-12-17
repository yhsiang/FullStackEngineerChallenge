package models

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/review360/database"
)

func mockReviewers(mock sqlmock.Sqlmock) {
	mockReviewers := sqlmock.NewRows([]string{"id", "name", "id"}).
		AddRow("2", "test2", "15")

	mock.ExpectQuery("SELECT E.id, E.name, R.id FROM employees AS E, review_assignments AS R WHERE E.id = R.reviewer AND R.reviewee = ?").
		WithArgs(1).
		WillReturnRows(mockReviewers)
}

func mockReviewees(mock sqlmock.Sqlmock) {
	mockReviewees := sqlmock.NewRows([]string{"id", "name", "review_id"}).
		AddRow("2", "test3", "3")

	mock.ExpectQuery(`SELECT employees.id AS id, employees.name AS name, performance_reviews.id AS review_id FROM employees
	INNER JOIN review_assignments on employees.id = review_assignments.reviewee
	LEFT JOIN performance_reviews on review_assignments.id = performance_reviews.assign_id
	WHERE review_assignments.reviewer = ?`).
		WithArgs(1).
		WillReturnRows(mockReviewees)
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestFindAllEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	mockEmployee := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "test1")

	mock.ExpectQuery("SELECT id, name FROM employees").
		WillReturnRows(mockEmployee)

	mockReviewers(mock)
	mockReviewees(mock)

	mockDB := database.NewWithDB(db)
	var em Employee
	employees, err := em.FindAll(mockDB)
	assert.NoError(t, err)
	assert.Len(t, employees, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	mockEmployee := sqlmock.NewRows([]string{"name"}).
		AddRow("test1")

	mock.ExpectQuery("SELECT name FROM employees WHERE id = ?").
		WithArgs(1).
		WillReturnRows(mockEmployee)

	mockReviewers(mock)
	mockReviewees(mock)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	em := &Employee{
		ID: 1,
	}
	mockDB := database.NewWithDB(db)
	employee, err := em.Find(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, "test1", employee.Name)
	assert.Equal(t, "test2", employee.Reviewers[0].Name)
	assert.Equal(t, "test3", employee.Reviewees[0].Name)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveWithInsertEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	em := &Employee{
		Name: "test2",
	}
	mock.ExpectExec("^INSERT INTO employees").
		WithArgs(em.Name, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mockReviewers(mock)
	mockReviewees(mock)
	defer db.Close()

	mockDB := database.NewWithDB(db)
	employee, err := em.Save(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, em.Name, employee.Name)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveWithUpdateEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	em := &Employee{
		ID:   1,
		Name: "test2",
	}
	mock.ExpectExec("^UPDATE employees SET name").
		WithArgs(em.Name, AnyTime{}, em.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mockReviewers(mock)
	mockReviewees(mock)
	defer db.Close()

	mockDB := database.NewWithDB(db)
	employee, err := em.Save(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, em.Name, employee.Name)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	em := &Employee{
		ID: 1,
	}

	mock.ExpectExec("^DELETE FROM performance_reviews").
		WithArgs(em.ID, em.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("^DELETE FROM review_assignments").
		WithArgs(em.ID, em.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("^DELETE FROM employees").
		WithArgs(em.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mockDB := database.NewWithDB(db)
	err = em.Remove(mockDB)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
