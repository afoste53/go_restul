package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type album struct {
	ID	string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

// seeder
var albums = []album{
	{ ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99 },
	{ ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99 },
	{ ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99 },
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

// getAlbums fetches list of all albums jsonified
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums -> add an album from JSON
func postAlbums(c *gin.Context){
	var newAlbum album

	// Call bindjson to bind received json to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to albums slice
	albums = append(albums, newAlbum)

	c.IndentedJSON(http:StatusCreated, newAlbum)
}
