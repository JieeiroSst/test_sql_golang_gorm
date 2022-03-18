package orm_test 

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"test_sql/app/model"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

var u = &model.User{
	PhoneNumber: "982721224",
	Password:    "123456",
}

func NewMock() (IUser, *sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gdb, err := gorm.Open("postgres", db) // open gorm db
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	userOrm := InitUserOrm(gdb)
	return userOrm, db, mock
}

func TestGetByPhoneNumber(t *testing.T) {
	userOrm, db, mock := NewMock()

	defer func() {
		db.Close()
	}()

	const sqlSelectOne = `SELECT * FROM "users" WHERE (phone_number = $1) ORDER BY id DESC LIMIT 1`
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
		AddRow(u.ID, u.PhoneNumber, u.Password, u.CreatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).WithArgs(u.PhoneNumber).WillReturnRows(rows)

	user, err := userOrm.GetByPhoneNumber(u.PhoneNumber)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestGetByPhoneNumberNotFound(t *testing.T) {
	userOrm, db, mock := NewMock()

	defer func() {
		db.Close()
	}()

	mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))

	user, err := userOrm.GetByPhoneNumber("123456789")
	assert.Empty(t, user)
	assert.Empty(t, err)
}