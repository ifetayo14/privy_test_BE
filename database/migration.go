package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	log "privy_cake_store/utils"
)

type (
	Migration struct {
		db *sql.DB
	}
)

func NewMigration(db *sql.DB) *Migration {
	return &Migration{
		db: db,
	}
}

func (m *Migration) Migrate(g *gin.Context) {
	db := m.db

	query := `CREATE TABLE IF NOT EXISTS cake(
    			id int PRIMARY KEY auto_increment,
    			title VARCHAR(225),
    			description VARCHAR(225),
    			rating int,
    			image VARCHAR(225),
    			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

	_, err := db.Query(query)
	if err != nil {
		log.FatalError(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "nil",
			"message": "migration fail",
		})
	}

	log.Success("migration success")
	g.JSON(http.StatusOK, gin.H{
		"error":   "nil",
		"message": "migration success",
	})
}
