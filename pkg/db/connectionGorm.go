package db

import (
	"fmt"
"github.com/prasanth-pn/grpc-product-service/pkg/domain"
	"github.com/prasanth-pn/grpc-product-service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	fmt.Println("gorm connected")
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.StockDecreaseLog)
	return db, err
}
