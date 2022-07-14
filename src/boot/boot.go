package boot

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/outerwolf/http-server-gin-go/src/application/controller"
)

func CreateHTTPGinServer() *gin.Engine {
	// Create a new mux.
	return gin.Default()
}
func RegisterHttpRoutesGin(mux *gin.Engine) {
	mux.GET("/api/health", controller.Health2())
}
func StartHttpGinServer(mux *gin.Engine) {
	http.ListenAndServe(":5000", mux)
}

// mux server
func CreateHttpServer() *http.ServeMux {
	// Create a new mux.
	return http.NewServeMux()
}
func RegisterHttpRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", controller.Health)
}
func StartHttpServer(mux *http.ServeMux) {
	http.ListenAndServe(":8080", mux)
}
