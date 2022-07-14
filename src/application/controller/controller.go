package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/// controller handler
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
func Health2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	}
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	// validate request

}
