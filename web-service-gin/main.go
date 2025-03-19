package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// writing `json: "artist"` specifies the name of the field
// when the struct is being serialized into JSON
// if not mentioned, it will use the name of the field as in the struct
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with list of all the albums as JSON
// gin.Context carries the request details, validates and serializes JSON
func getAlbums(context *gin.Context) {
	// context.IndentedJSON serializes the struct into JSON and adds it to the response
	// the first argument is the HTTP status code and second is the obj which has to be serialized
	context.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(context *gin.Context) {
	var newAlbum album

	// context.BindJSON binds the recieved JSON with the struct newAlbum
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusAccepted, albums)
}

func getAlbumById(context *gin.Context) {
	// retrieves the id tag mentioned in the params path
	id := context.Param("id")

	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.GET("/album/:id", getAlbumById)

	router.Run("localhost:8080")
}