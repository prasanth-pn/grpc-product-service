package domain




type Products struct{
	Id int  `gorm:"serial primaryKey;autoIncrement;unique"`
	Product_Name string `json:"product_name"`
	Stock int `json:"stock"`
	Price int `json:"price"`
	StockDecreaseLogs StockDecreaseLog `gorm:"foreignKey:ProductRefer"`

}

type StockDecreaseLog struct{
	Id int `json:"id" gorm:"primayKey"`
	OrderId int `json:"order_id"`
	ProductRefer int `json:"product_id"`
}