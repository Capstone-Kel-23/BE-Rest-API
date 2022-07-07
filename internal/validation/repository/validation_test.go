package repository_test

import (
	"database/sql"
	"database/sql/driver"
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/validation/repository"
	sqlMock "github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func SetupDBMock(dbMock *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		DSN:                       "sqlmock_db_0",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{PrepareStmt: false})
	if err != nil {
		panic(err)
	}
	return gormDB
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var dummyValidation = []domain.Validation{
	domain.Validation{
		ID:        uuid.FromStringOrNil("1"),
		UserEmail: "test@email.com",
		Type:      "verify",
		Code:      "12345",
		ExpiredAt: time.Time{},
	},
	domain.Validation{
		ID:        uuid.FromStringOrNil("2"),
		UserEmail: "test2@email.com",
		Type:      "verify",
		Code:      "12345",
		ExpiredAt: time.Time{},
	},
}

func TestValidationRepository_Save(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `validations` (`id`,`user_email`,`type`,`code`,`expired_at`) VALUES (?,?,?,?,?)").
		WithArgs(dummyValidation[0].ID, dummyValidation[0].UserEmail, dummyValidation[0].Type, dummyValidation[0].Code, AnyTime{}).WillReturnResult(sqlMock.NewErrorResult(nil))
	mock.ExpectCommit()

	validationRepository := repository.NewValidationRepository(db)
	save, err := validationRepository.Save(&dummyValidation[0])
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, save)
}

func TestValidationRepository_Delete(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `validations` WHERE user_email = ? AND type = ?").
		WithArgs(dummyValidation[0].UserEmail, dummyValidation[0].Type).WillReturnResult(sqlMock.NewErrorResult(nil))
	mock.ExpectCommit()

	validationRepository := repository.NewValidationRepository(db)
	err = validationRepository.Delete(&dummyValidation[0])
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
}

func TestValidationRepository_FindByCode(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectQuery("SELECT * FROM `validations` WHERE code = ?").
		WithArgs(dummyValidation[0].Code).
		WillReturnRows(sqlMock.NewRows([]string{"id", "user_email", "type", "code", "expired_at"}).
			AddRow(dummyValidation[0].ID, dummyValidation[0].UserEmail, dummyValidation[0].Type, dummyValidation[0].Code, dummyValidation[0].ExpiredAt))

	validationRepository := repository.NewValidationRepository(db)
	val, err := validationRepository.FindByCode(dummyValidation[0].Code)
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, val)
}

func TestValidationRepository_FindByEmailAndType(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectQuery("SELECT * FROM `validations` WHERE user_email = ? AND type = ?").
		WithArgs(dummyValidation[0].UserEmail, "verify").
		WillReturnRows(sqlMock.NewRows([]string{"id", "user_email", "type", "code", "expired_at"}).
			AddRow(dummyValidation[0].ID, dummyValidation[0].UserEmail, dummyValidation[0].Type, dummyValidation[0].Code, dummyValidation[0].ExpiredAt))

	validationRepository := repository.NewValidationRepository(db)
	val, err := validationRepository.FindByEmailAndType(dummyValidation[0].UserEmail, "verify")
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, val)
}
