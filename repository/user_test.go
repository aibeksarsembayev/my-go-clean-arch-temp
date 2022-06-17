package repository_test

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
	"github.com/pashagolub/pgxmock"
)

// successful case
func TestCreate(t *testing.T) {
	mock, err := pqxmock.NewConn()
	if err != nil{
		t.Fatalf("an error '%s' was not expectd when opening a stuv database connection", err)

	}
	defer mock.Close(context.Background())

	now := time.Now()
	u := &models.User{

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO user").WithArgs("aibek", "123456", "s-bek-k@mail.ru", now).WillReturnResult(pgxmock.NewResult(1, nil))
	mock.ExpectCommit()

	// now we execute our method
	if err = Create(context)


		
	}
}