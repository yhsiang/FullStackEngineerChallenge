package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/review360/database"
)

func TestFindAssignID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	as := &Assignment{
		ID:       3,
		Reviewee: 1,
		Reviewer: 2,
	}

	mockAssign := sqlmock.NewRows([]string{"id"}).
		AddRow(as.ID)

	mock.ExpectQuery("^SELECT id FROM review_assignments WHERE reviewee").
		WithArgs(as.Reviewee, as.Reviewer).
		WillReturnRows(mockAssign)
	defer db.Close()

	mockDB := database.NewWithDB(db)
	assignID, err := as.FindAssignID(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, as.ID, assignID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveAssignment(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	as := &Assignment{
		Reviewee: 1,
		Reviewer: 2,
	}

	mock.ExpectExec("^INSERT INTO review_assignments").
		WithArgs(as.Reviewee, as.Reviewer, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	defer db.Close()

	mockDB := database.NewWithDB(db)
	assign, err := as.Save(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, as.Reviewee, assign.Reviewee)
	assert.Equal(t, as.Reviewer, assign.Reviewer)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveAssignment(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	as := &Assignment{
		Reviewee: 1,
		Reviewer: 2,
	}

	mock.ExpectExec("^DELETE FROM review_assignments").
		WithArgs(as.Reviewee, as.Reviewer).
		WillReturnResult(sqlmock.NewResult(0, 1))
	defer db.Close()

	mockDB := database.NewWithDB(db)
	err = as.Remove(mockDB)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
