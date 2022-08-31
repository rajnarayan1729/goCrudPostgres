package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rajnarayan1729/goCrudPostgres/controller"
	"github.com/rajnarayan1729/goCrudPostgres/service"
	"github.com/rajnarayan1729/goCrudPostgres/utils"
	"gorm.io/gorm"
)

var (
	server *gin.Engine
	ps     service.ProductService
	pc     controller.ProductController
	db     *gorm.DB
	err    error
)

func init() {

	db, err = utils.GetDbConn()
	if err != nil {
		log.Fatal("db connection failed")
	}

	ps = service.NewServiceImpl(db)
	pc = controller.NewController(ps)
	server = gin.Default()

}

func main() {

	basepath := server.Group("/v1")
	pc.RegisterRoutes(basepath)

	log.Fatal(server.Run("localhost:9090"))

}
