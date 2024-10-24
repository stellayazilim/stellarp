package authHandler

import "github.com/gin-gonic/gin"

func UseAuthHandlers(r *gin.RouterGroup) {

	r.POST("/login", loginHandler)
}
