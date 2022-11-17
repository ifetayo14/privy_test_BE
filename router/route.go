package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"privy_cake_store/controller"
	"privy_cake_store/database"
)

type (
	CakeHandler struct {
		db *sql.DB
		r  *gin.Engine
	}
)

func NewCakeHandler(db *sql.DB, engine *gin.Engine) *CakeHandler {
	return &CakeHandler{
		r:  engine,
		db: db,
	}
}

func (c *CakeHandler) InitRest() (*gin.Engine, error) {
	db := c.db
	route := c.r

	cakeController := controller.NewCakeController(db)
	migration := database.NewMigration(db)

	cakeRoute := route.Group("/cakes")
	{
		cakeRoute.POST("/", cakeController.Create)
		cakeRoute.GET("/", cakeController.ListAll)
		cakeRoute.GET("/:id", cakeController.Detail)
		cakeRoute.PATCH("/:id", cakeController.Update)
		cakeRoute.DELETE("/:id", cakeController.Delete)
	}

	migrate := route.Group("/migrate")
	{
		migrate.GET("/", migration.Migrate)
	}

	return route, nil
}
