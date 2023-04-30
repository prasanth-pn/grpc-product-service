package main

import (
	"fmt"
	"log"

	"github.com/prasanth-pn/grpc-product-service/pkg/config"
)

func main() {
	c, err := config.LoadConfig()
	fmt.Println(c,"port")
	if err!=nil{
		log.Fatalln("failed at config",err)
	}
	

}
