package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/sew810i9/opus"
	"github.com/sew810i9/opus/internal/config"
	"log"
)

func Init() {
	srv := new(opus.Server)
	cnf := config.Init()

	connection, err := sqlx.Open("mysql", cnf.MysqlDsn)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer connection.Close()

	log.Fatalln(srv.Run(cnf.Listen, nil))
}
