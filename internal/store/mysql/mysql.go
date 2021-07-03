package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/sew810i9/opus/configs"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(s *configs.Config) (*sqlx.DB, error) {
	return sqlx.Open("mysql", s.MysqlConnection)
}
