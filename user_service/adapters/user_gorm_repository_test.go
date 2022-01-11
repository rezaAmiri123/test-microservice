package adapters_test
//
//import (
//	"context"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/google/uuid"
//	"github.com/jinzhu/gorm"
//	"github.com/rezaAmiri123/test-microservice/user_service/adapters"
//	"github.com/stretchr/testify/require"
//	"regexp"
//	"testing"
//	"time"
//)
//
//var createQuery = `INSERT INTO "gorm_user_models" ("created_at","updated_at","deleted_at","uuid","username","password","email","bio","image") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
//var create = `INSERT INTO "gorm_user_models" ("created_at","updated_at","deleted_at","uuid","username","password","email","bio","image") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "gorm_user_models"."id"`
//var create2= `INSERT INTO "gorm_user_models" ("created_at","updated_at","deleted_at","uuid","username","password","email","bio","image") VALUES ('2022-01-10 16:44:48','2022-01-10 16:44:48',NULL,'7f090540-e177-4ab5-a3c3-00904cfbcfd2','username','password','email@example.com','','') RETURNING "gorm_user_models"."id"`
//func TestGORMUserRepository_Create(t *testing.T) {
//	//t.Parallel()
//
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
//	require.NoError(t, err)
//	defer db.Close()
//
//	gormDB, err := gorm.Open("postgres", db)
//	require.NoError(t, err)
//	defer gormDB.Close()
//
//	newRepo := &adapters.GORMUserRepository{gormDB}
//	now := time.Now()
//	gormUser := &adapters.GORMUserModel{}
//	gormUser.CreatedAt = now
//	gormUser.UpdatedAt = now
//	gormUser.DeletedAt = nil
//	gormUser.UUID = uuid.NewString()
//	gormUser.Username = "username"
//	gormUser.Password = "password"
//	gormUser.Email = "email@example.com"
//	gormUser.Bio = ""
//	gormUser.Image = ""
//
//	//rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "uuid", "username", "password", "email", "bio", "image"}).AddRow(
//	//	1,
//	//	gormUser.CreatedAt.String(),
//	//	gormUser.UpdatedAt.String(),
//	//	gormUser.DeletedAt,
//	//	gormUser.UUID,
//	//	gormUser.Username,
//	//	gormUser.Password,
//	//	gormUser.Email,
//	//	gormUser.Bio,
//	//	gormUser.Image,
//	//)
//	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
//
//	//mock.ExpectBegin()
//	//mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
//	//mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
//	//mock.ExpectCommit()
//	mock.ExpectBegin()
//	//mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
//	//mock.ExpectExec("INSERT INTO gorm_user_models").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
//	// ("created_at","updated_at","deleted_at","uuid","username","password","email","bio","image")
//	mock.ExpectQuery(regexp.QuoteMeta(create)).WithArgs(
//		gormUser.CreatedAt.UTC().Format("2006-01-02 15:04:05"),
//		gormUser.UpdatedAt.UTC().Format("2006-01-02 15:04:05"),
//		nil,
//		gormUser.UUID,
//		gormUser.Username,
//		gormUser.Password,
//		gormUser.Email,
//		gormUser.Bio,
//		gormUser.Image,
//		).WillReturnRows(rows)
//	//mock.ExpectQuery(create2).WillReturnRows(rows)
//	mock.ExpectExec("INSERT INTO gorm_user_models").WithArgs(
//		"2006-01-02 15:04:05",
//		"2006-01-02 15:04:05",
//		nil,
//		gormUser.UUID,
//		gormUser.Username,
//		gormUser.Password,
//		gormUser.Email,
//		gormUser.Bio,
//		gormUser.Image,).WillReturnResult(sqlmock.NewResult(mock.EeoID, 1))
//	mock.ExpectCommit()
//	err = newRepo.Create(context.Background(), gormUser.ProtoDomainUser())
//	require.NoError(t, err)
//
//}
