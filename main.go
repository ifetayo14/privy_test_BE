package main

import (
	"github.com/gin-gonic/gin"
	"privy_cake_store/database"
	"privy_cake_store/router"
	log "privy_cake_store/utils"
)

func main() {
	log.Info("starting application")

	db, err := database.StartDB()
	log.Info("starting DB")
	if err != nil {
		log.FatalError("error on starting DB")
	}

	log.Info("initiating rest api")
	gin.SetMode(gin.ReleaseMode)

	rest := router.NewCakeHandler(db, gin.New())
	r, err := rest.InitRest()
	if err != nil {
		log.FatalError("error initiating rest")
	}

	err = r.Run(":8080")
	log.Info("server running on :8080")
	if err != nil {
		log.FatalError("error running rest")
	}

	err = db.Close()
	if err != nil {
		log.FatalError("fail to close DB")
	}

	log.Info("successfully close DB")
}
