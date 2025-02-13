package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.PUT("/v1/:key", KeyValuePutHandler)
	r.GET("/v1/:key", KeyValueGetHandler)
	r.DELETE("/v1/:key", KeyValueDeleteHandler)
	r.GET("v2/:key", KeyValueGetHandler)

	r.Run(":3009")
}
