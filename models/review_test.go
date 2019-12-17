package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yhsiang/review360/database"
)

func TestSaveWithInsertReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	re := &Review{
		Content:  "test-content",
		AssignID: 1,
	}

	mock.ExpectExec("^INSERT INTO performance_reviews").
		WithArgs(re.Content, re.AssignID, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	defer db.Close()

	mockDB := database.NewWithDB(db)
	review, err := re.Save(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, re.Content, review.Content)
	assert.Equal(t, re.AssignID, review.AssignID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveWithUpdateReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	re := &Review{
		ID:       1,
		Content:  "test-content",
		AssignID: 1,
	}

	mock.ExpectExec("^UPDATE performance_reviews SET content").
		WithArgs(re.Content, AnyTime{}, re.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	defer db.Close()

	mockDB := database.NewWithDB(db)
	review, err := re.Save(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, re.Content, review.Content)
	assert.Equal(t, re.AssignID, review.AssignID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindAllReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mockReview := sqlmock.NewRows([]string{"id", "content", "assign_id"}).
		AddRow(1, "test-content-1", 2).
		AddRow(1, "test-content-2", 3)

	mock.ExpectQuery("select id, content, assign_id from performance_reviews").
		WillReturnRows(mockReview)

	defer db.Close()

	mockDB := database.NewWithDB(db)
	var re Review
	reviews, err := re.FindAll(mockDB)
	assert.NoError(t, err)
	assert.Len(t, reviews, 2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestFindReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	re := &Review{
		ID:       1,
		AssignID: 1,
		Content:  "test-content",
	}

	mockReview := sqlmock.NewRows([]string{"assign_id", "content"}).
		AddRow(re.AssignID, re.Content)

	mock.ExpectQuery("select assign_id, content from performance_reviews WHERE id = ?").
		WithArgs(re.ID).
		WillReturnRows(mockReview)
	defer db.Close()

	mockDB := database.NewWithDB(db)
	review, err := re.Find(mockDB)
	assert.NoError(t, err)
	assert.Equal(t, re.Content, review.Content)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
