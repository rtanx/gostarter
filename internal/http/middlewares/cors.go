package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func EnableCors() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AllowMethods = []string{"POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	conf.AllowHeaders = []string{"Authorization", "Access-Control-Allow-Headers", "Origin", "Accept", "X-Requested-With", "Content-Type", "Content-Length", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	// conf.AllowCredentials = true
	return cors.New(conf)
}
