package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"github.com/prasanth-pn/grpc-product-service/pkg/config"
)

func ConnectDB(cfg config.Config) (*sql.DB) {
	db, err := sql.Open("postgres", cfg.DBSource)
	if err != nil {
		log.Fatalf(" the data base is not connected : %s", err)
	}
	err=db.Ping()
	if err!=nil{
		log.Fatalf("error in ping : %s",err)
	}
	return db
}
