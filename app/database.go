package app

import (
	"database/sql"
	"golang-restful-api/model/helper"
	"time"
)

func NewDB() *sql.DB {
	dbName := helper.LoadEnv("DB_NAME")
	port := helper.LoadEnv("PORT")
	dbUser := helper.LoadEnv("DB_USER")

	db, err := sql.Open("mysql", dbUser+"@tcp(localhost:"+port+")/"+dbName)
	helper.PanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// BISA PAKE LEVEL (OPSIONAL)

// migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migrations" -path db/migrations up

// ROLLBACK
// migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migrations" -path db/migrations down

// FORCE KALO ADA DIRTY (KE VERSI SEBELUMNYA)
// migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migrations" -path db/migrations force xxxxxxxx

// BUAT MIGRATION BARU
// migrate create -ext sql -dir db/migrations migration_name                                           
