package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/rafabcanedo/CyclePix/application/grpc"
	"github.com/rafabcanedo/CyclePix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
