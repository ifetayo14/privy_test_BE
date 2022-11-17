package controller

import (
	"database/sql"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"privy_cake_store/model"
	log "privy_cake_store/utils"
	"strconv"
	"strings"
)

type (
	CakeController struct {
		db *sql.DB
	}
)

func NewCakeController(db *sql.DB) *CakeController {
	return &CakeController{
		db: db,
	}
}

func (c *CakeController) Create(g *gin.Context) {
	db := c.db
	cake := model.Cake{}

	err := g.ShouldBindJSON(&cake)
	if err != nil {
		log.Error("error on binding json")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	_, err = govalidator.ValidateStruct(cake)
	if err != nil {
		log.Error("validation error")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	_, err = db.Query("INSERT INTO `cake` (title, description, rating, image) VALUES (?, ?, ?, ?);", cake.Title, cake.Description, cake.Rating, cake.Image)
	if err != nil {
		log.Error("error insert cake")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	log.Success("create new cake")
	g.JSON(http.StatusCreated, gin.H{
		"error":   "nil",
		"message": "insert success",
		"value": gin.H{
			"title":       cake.Title,
			"description": cake.Description,
			"rating":      cake.Rating,
			"image":       cake.Image,
		},
	})
}

func (c *CakeController) ListAll(g *gin.Context) {
	var result []model.Cake
	db := c.db

	query := "SELECT * FROM cake ORDER BY rating DESC, title ASC;"
	rows, err := db.Query(query)
	if err != nil {
		log.Error("error retrieve cake")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	for rows.Next() {
		var eachCake = model.Cake{}
		err = rows.Scan(&eachCake.ID, &eachCake.Title, &eachCake.Description, &eachCake.Rating, &eachCake.Image, &eachCake.CreatedAt, &eachCake.UpdatedAt)
		if err != nil {
			log.Error("error retrieve cake")
			g.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}

		result = append(result, eachCake)
	}

	log.Success("list all cake")
	g.JSON(http.StatusOK, gin.H{
		"error":   "nil",
		"message": "retrieve all success",
		"value":   result,
	})
}

func (c *CakeController) Detail(g *gin.Context) {
	db := c.db
	cake := model.Cake{}

	url := g.Request.URL.RawQuery
	parse, err := g.Request.URL.Parse(url)
	if err != nil {
		log.Error("error parse url")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	split := strings.Split(parse.Path, ":")
	id, err := strconv.Atoi(split[1])
	if err != nil {
		log.Error("error converting str to int")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.QueryRow("SELECT * FROM cake WHERE id = ?", id).Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err != nil {
		log.Error("error retrieve cake")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	log.Success("list specified cake")
	g.JSON(http.StatusOK, gin.H{
		"error":   "nil",
		"message": "retrieve all success",
		"value":   cake,
	})
}

func (c *CakeController) Update(g *gin.Context) {
	db := c.db
	cake := model.Cake{}

	err := g.ShouldBindJSON(&cake)
	if err != nil {
		log.Error("error on binding json")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	_, err = govalidator.ValidateStruct(cake)
	if err != nil {
		log.Error("validation error")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	url := g.Request.URL.RawQuery
	parse, err := g.Request.URL.Parse(url)
	if err != nil {
		log.Error("error parse url")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	split := strings.Split(parse.Path, ":")
	id, err := strconv.Atoi(split[1])
	if err != nil {
		log.Error("error converting str to int")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	rows, err := db.Exec("UPDATE `cake` SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?", cake.Title, cake.Description, cake.Rating, cake.Image, id)
	if err != nil {
		log.Error("error update")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if aff, _ := rows.RowsAffected(); aff < 1 {
		log.Error("error update")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "data not found",
		})
		return
	}

	log.Success("update cake")
	g.JSON(http.StatusAccepted, gin.H{
		"error":   "nil",
		"message": "update success",
		"value": gin.H{
			"id":          id,
			"title":       cake.Title,
			"description": cake.Description,
			"rating":      cake.Rating,
			"image":       cake.Image,
		},
	})
}

func (c *CakeController) Delete(g *gin.Context) {
	db := c.db

	url := g.Request.URL.RawQuery
	parse, err := g.Request.URL.Parse(url)
	if err != nil {
		log.Error("error parse url")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	split := strings.Split(parse.Path, ":")
	id, err := strconv.Atoi(split[1])
	if err != nil {
		log.Error("error converting str to int")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	rows, err := db.Exec("DELETE FROM cake WHERE id = ?", id)
	if err != nil {
		log.Error("error delete")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if aff, _ := rows.RowsAffected(); aff < 1 {
		log.Error("error delete")
		g.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "data not found",
		})
		return
	}

	log.Success("delete cake")
	g.JSON(http.StatusAccepted, gin.H{
		"error":   "nil",
		"message": "delete success",
	})
}
