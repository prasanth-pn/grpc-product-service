package repository

import (
	"database/sql"

	"github.com/prasanth-pn/grpc-product-service/pkg/repository/interfaces"
)

type ProductDatabase struct {
	DB *sql.DB
}
func NewAuthRepository(db *sql.DB)interfaces.ProductRepository{
	return &ProductDatabase{db}
}
