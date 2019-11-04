package routing

import (
	gin "github.com/gin-gonic/gin"
)

// Init routing
func Init() {
	r := gin.Default()

	api := r.Group("/api")
	public := r.Group("/api")

	InitUsersControllerPrivate(api.Group("/users"))
	InitUsersControllerPublic(public.Group("/users"))
	
	r.Run()
}