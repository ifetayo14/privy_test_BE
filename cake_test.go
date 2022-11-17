package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"privy_cake_store/controller"
	"privy_cake_store/database"
	log "privy_cake_store/utils"
	"testing"
)

func Router() *gin.Engine {
	route := gin.New()
	gin.SetMode(gin.ReleaseMode)
	db, err := database.StartDB()
	if err != nil {
		log.Error("error start db")
		return nil
	}

	cakeController := controller.NewCakeController(db)

	cakeRoute := route.Group("/cakes")
	{
		cakeRoute.POST("/", cakeController.Create)
		cakeRoute.GET("/", cakeController.ListAll)
		cakeRoute.GET("/:id", cakeController.Detail)
		cakeRoute.PATCH("/:id", cakeController.Update)
		cakeRoute.DELETE("/:id", cakeController.Delete)
	}

	return route
}

func TestCreate(t *testing.T) {
	testBody := `{"title": "The Most Simple Cake", "description": "very very good", "rating": 9, "image": "yourimage.com"}`
	request, _ := http.NewRequest("POST", "/cakes/", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code, "201 response is expected")

	testBody = `{"title": "Strawberry Cake", "description": "very very good", "image": "yourimage.com"}`
	request, _ = http.NewRequest("POST", "/cakes/", bytes.NewBufferString(testBody))
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "400 response is expected")

	db, err := database.StartDB()
	if err != nil {
		return
	}

	_, err = db.Query("UPDATE cake SET id = 909 WHERE title = 'The Most Simple Cake'")
}

func TestListAll(t *testing.T) {
	request, _ := http.NewRequest("GET", "/cakes/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")
}

func TestDetail(t *testing.T) {
	request, _ := http.NewRequest("GET", "/cakes/:909", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 response is expected")

	request, _ = http.NewRequest("GET", "/cakes/:9999", nil)
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "400 response is expected")
}

func TestUpdate(t *testing.T) {
	testBody := `{"id": 909, "title": "Peanut Cake", "description": "meh", "rating": 3, "image": "yourimage.com"}`
	request, _ := http.NewRequest("PATCH", "/cakes/:909", bytes.NewBufferString(testBody))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "202 response is expected")

	testBody = `{"id": 909, "title": "Peanut Cake", "rating": -1, "image": "yourimage.com"}`
	request, _ = http.NewRequest("PATCH", "/cakes/:909", bytes.NewBufferString(testBody))
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "400 response is expected")
}

func TestDelete(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/cakes/:99999", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code, "400 response is expected")

	request, _ = http.NewRequest("DELETE", "/cakes/:909", nil)
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 202, response.Code, "OK response is expected")
}
