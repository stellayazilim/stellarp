package handlers

import (
	"StellaRP/modules/identity/presentation/handlers/realmHandler"
	"github.com/gin-gonic/gin"
)

func UseHandlers(g *gin.Engine) {
	r := g.Group("")

	r.POST("/login", loginWithPasswordHandler)

	realmHandler.RealmHandler(g.Group("/realm"))

}
