package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// errDotenv := godotenv.Load()
	// if errDotenv != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// dbUser := os.Getenv("DB_USER")
	// dbName := os.Getenv("DB_NAME")
	// dbPassword := os.Getenv("DB_PASSWORD")

	// dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("DB connect error")
	// }

	// foodRepository := food.NewRepository(db)
	// food = new Food()
	// food.
	// foodService := food.NewService(foodRepository)
	// foodHandler := handler.NewFoodHandler(foodService)

	// orderRepository := order.NewRepository(db)
	// orderService := order.NewService(orderRepository)
	// orderHandler := handler.NewOrderHandler(orderService)

	// merchantRepo := merchant.NewRepository(db)
	// merchantServ := merchant.NewService(merchantRepo)
	// merchantHandler := handler.NewMerchantHandler(merchantServ)
	// 1 order punya 1 bill
	// bill: belum tau cara menggunakannya
	// billRepo := bill.NewRepository(db)
	// billService := bill.NewService(billRepo)
	// billHandler := handler.NewBillHandler(billService)

	// custRepository := customer.NewRepository(db)
	// custService := customer.NewService(db)
	// custHandler := customer.NewRepository(db)

	// buat migrasi kode entity to table
	// db.AutoMigrate(&food.Food{})

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "Ranufrozen",
			"panduan": "silahkan buka base_url/v1/url",
		})
	})

	// v1 := router.Group("/v1")

	// v1.GET("/food/:id", foodHandler.Show)
	// v1.GET("/foods/order_by", foodHandler.OrderBy)
	// v1.POST("/order", orderHandler.PostorderHandler)

	// fitur order, halaman apa saja yang mungkin muncul?
	// 1. as a pelanggan?
	// apakah sudah ada template html nya?
	//
	// 2. as A Customer Service (CS)
	// ini untuk dashboard CS nantinya
	// siapa tau ada masalah dengan transaksi
	// jadi CS Punya data

	// v1.GET("/order/:id", orderHandler.Show)
	// v1.GET("/orders/order_by", orderHandler.OrderBy)
	// v1.POST("/order", orderHandler.PostorderHandler)

	// v1.GET("/", merchantHandler.RootHandler)
	// v1.GET("/merchant/:id", merchantHandler.Show)
	// v1.GET("/merchants/order_by", merchantHandler.OrderBy)

	// ========================= V2 =====================
	// v2 := router.Group("/v2")
	// apakah handler nya dipisahkan saja?
	// v2.GET("/food/:id", foodHandler.Show)

	router.Run()
}
