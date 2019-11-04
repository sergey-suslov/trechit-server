package routing

import (
	gin "github.com/gin-gonic/gin"
)


// InitUsersControllerPublic returns controller for users public api
func InitUsersControllerPublic(r *gin.RouterGroup) {
	r.GET("", func (c *gin.Context) {
		// Implement routes
	})
}

// InitUsersControllerPrivate returns controller for users private api
func InitUsersControllerPrivate(r *gin.RouterGroup) {
	r.GET("", func (c *gin.Context) {
		// Implement routes
	})
}