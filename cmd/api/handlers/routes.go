package handlers

import (
	"somo/service/internal/logic"
	"somo/service/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func API(db *sqlx.DB, engine *gin.Engine) {
	productHandler := Product{
		DB:        db,
		ProductDB: product.PostgresDB{DB: db},
	}
	engine.POST("/v1/product", productHandler.CreateNewProduct)
	engine.GET("/v1/product/:id", productHandler.GetProductByID)
	engine.GET("/v1/product", productHandler.GetProducts)
	engine.PUT("/v1/product/:id", productHandler.UpdateProduct)
	engine.DELETE("/v1/product/:id", productHandler.DeleteProductByID)
}

func SomoAPI(db *gorm.DB, engine *gin.Engine) {
	customerHandler := Customer{
		CustomerDB: logic.PostgresDB{DB: db},
	}

	engine.POST("/v1/customer", customerHandler.Create)
	engine.PUT("/v1/customer/:id", customerHandler.Update)
	engine.DELETE("/v1/customer/:id", customerHandler.Delete)
	engine.GET("/v1/customer/:id", customerHandler.Get)
	engine.GET("/v1/customer", customerHandler.List)

	reservationHandler := Reservation{
		ReservationDB: logic.PostgresDB{DB: db},
	}
	engine.POST("/v1/reservation", reservationHandler.Create)
	engine.PUT("/v1/reservation/:id", reservationHandler.Update)
	engine.DELETE("/v1/reservation/:id", reservationHandler.Delete)
	engine.GET("/v1/reservation/:id", reservationHandler.GetOne)
	engine.GET("/v1/reservation/search", reservationHandler.Search)
	engine.PUT("/v1/reservation/cancel/:id", reservationHandler.Cancel)

}
