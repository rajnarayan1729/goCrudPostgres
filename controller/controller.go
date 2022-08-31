package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajnarayan1729/goCrudPostgres/model"
	"github.com/rajnarayan1729/goCrudPostgres/service"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewController(productService service.ProductService) ProductController {
	return ProductController{

		ProductService: productService,
	}
}

func (pc *ProductController) GetProduct(c *gin.Context) {

	id := c.Param("code")
	p, err := pc.ProductService.GetProduct(&id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)

}

func (pc *ProductController) CreateProduct(c *gin.Context) {

	var p model.Product
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	result, err := pc.ProductService.CreateProduct(&p)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "ID": result})

}

func (pc *ProductController) UpdatePrice(c *gin.Context) {

	var p model.Product
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	if err := pc.ProductService.UpdatePrice(&p); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})

}

func (pc *ProductController) GetAllProduct(c *gin.Context) {

	p, err := pc.ProductService.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {

	code := c.Param("code")
	if err := pc.ProductService.DeleteProduct(&code); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (pc *ProductController) RegisterRoutes(rg *gin.RouterGroup) {

	productRoute := rg.Group("/product")
	productRoute.GET("/getProduct/:code", pc.GetProduct)
	productRoute.POST("/addProduct", pc.CreateProduct)
	productRoute.PUT("/updateProduct", pc.UpdatePrice)
	productRoute.GET("/getProduct", pc.GetAllProduct)
	productRoute.DELETE("/deleteProduct/:code", pc.DeleteProduct)

}
