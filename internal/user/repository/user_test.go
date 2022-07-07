package repository_test

import (
	"database/sql"
	"database/sql/driver"
	"github.com/Capstone-Kel-23/BE-Rest-API/domain"
	"github.com/Capstone-Kel-23/BE-Rest-API/internal/user/repository"
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

var dummyUser = []domain.User{
	domain.User{
		ID:        uuid.FromStringOrNil("1"),
		Fullname:  "testing",
		Email:     "test@email.com",
		Username:  "test",
		Password:  "12345678",
		Verified:  false,
		Roles:     nil,
		Profile:   domain.Profile{},
		Invoices:  nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	domain.User{
		ID:        uuid.FromStringOrNil("2"),
		Fullname:  "testing2",
		Email:     "test2@email.com",
		Username:  "test2",
		Password:  "12345678",
		Verified:  false,
		Roles:     nil,
		Profile:   domain.Profile{},
		Invoices:  nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
}

func TestUserRepository_FindAll(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectQuery("SELECT * FROM `users`").
		WillReturnRows(sqlMock.NewRows([]string{"id", "fullname", "email", "username", "password", "verified", "created_at", "updated_at"}).
			AddRow(dummyUser[0].ID, dummyUser[0].Fullname, dummyUser[0].Email, dummyUser[0].Username, dummyUser[0].Password, dummyUser[0].Verified, dummyUser[0].CreatedAt, dummyUser[0].UpdatedAt).
			AddRow(dummyUser[1].ID, dummyUser[1].Fullname, dummyUser[1].Email, dummyUser[1].Username, dummyUser[1].Password, dummyUser[1].Verified, dummyUser[1].CreatedAt, dummyUser[1].UpdatedAt))

	userRepository := repository.NewUserRepository(db)
	users, err := userRepository.FindAll()
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectQuery("SELECT * FROM `users` WHERE email = ?").
		WithArgs(dummyUser[0].Email).
		WillReturnRows(sqlMock.NewRows([]string{"id", "fullname", "email", "username", "password", "verified", "created_at", "updated_at"}).
			AddRow(dummyUser[0].ID, dummyUser[0].Fullname, dummyUser[0].Email, dummyUser[0].Username, dummyUser[0].Password, dummyUser[0].Verified, dummyUser[0].CreatedAt, dummyUser[0].UpdatedAt))

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindByEmail(dummyUser[0].Email)
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByID(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectQuery("SELECT * FROM `users` WHERE id = ?").
		WithArgs(dummyUser[0].ID.String()).
		WillReturnRows(sqlMock.NewRows([]string{"id", "fullname", "email", "username", "password", "verified", "created_at", "updated_at"}).
			AddRow(dummyUser[0].ID, dummyUser[0].Fullname, dummyUser[0].Email, dummyUser[0].Username, dummyUser[0].Password, dummyUser[0].Verified, dummyUser[0].CreatedAt, dummyUser[0].UpdatedAt))

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindByID(dummyUser[0].ID.String())
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_UpdateVerifiedByEmail(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users` SET `verified`=?,`updated_at`=? WHERE email = ?").
		WithArgs(true, AnyTime{}, dummyUser[0].Email).WillReturnResult(sqlMock.NewErrorResult(nil))
	mock.ExpectCommit()

	userRepository := repository.NewUserRepository(db)
	err = userRepository.UpdateVerifiedByEmail(dummyUser[0].Email, true)
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
}

func TestUserRepository_FindWithProfile(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectQuery("SELECT * FROM `users` WHERE id = ?").
		WithArgs(dummyUser[0].ID.String()).
		WillReturnRows(sqlMock.NewRows([]string{"id", "fullname", "email", "username", "password", "verified", "created_at", "updated_at"}).
			AddRow(dummyUser[0].ID, dummyUser[0].Fullname, dummyUser[0].Email, dummyUser[0].Username, dummyUser[0].Password, dummyUser[0].Verified, dummyUser[0].CreatedAt, dummyUser[0].UpdatedAt))

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindWithProfile(dummyUser[0].ID.String())
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_Save(t *testing.T) {
	dbMock, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db := SetupDBMock(dbMock)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`id`,`fullname`,`email`,`username`,`password`,`verified`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)").
		WithArgs(dummyUser[0].ID, dummyUser[0].Fullname, dummyUser[0].Email, dummyUser[0].Username, dummyUser[0].Password, dummyUser[0].Verified, AnyTime{}, AnyTime{}).WillReturnResult(sqlMock.NewErrorResult(nil))
	mock.ExpectCommit()

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.Save(&dummyUser[0])
	if err != nil {
		assert.Error(t, err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
