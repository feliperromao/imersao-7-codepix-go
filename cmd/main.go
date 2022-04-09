package main

import (
	"os"

	"github.com/feliperromao/imersao-7-codepix-go/application/grpc"
	"github.com/feliperromao/imersao-7-codepix-go/infra/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}