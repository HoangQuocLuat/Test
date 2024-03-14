package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Sql struct {
	Db       *sqlx.DB
	UserName string
	Password string
	DbName   string
}

func (s *Sql) Connect() {
	s.Db = sqlx.MustConnect("postgres", "user="+s.UserName+" password="+s.Password+" dbname="+s.DbName+" sslmode=disable")

	if err := s.Db.Ping(); err != nil {
		panic(err.Error())
	}
	log.Println("Successfully Connected")

}

func (s *Sql) Close() {
	s.Db.Close()
}
