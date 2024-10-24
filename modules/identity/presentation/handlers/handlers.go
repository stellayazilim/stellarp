package handlers

import (
	"StellaRP/modules/identity/presentation/handlers/authHandler"
	"StellaRP/modules/identity/presentation/handlers/realmHandler"
	"github.com/gin-gonic/gin"
)

func UseHandlers(g *gin.Engine) {

	authHandler.UseAuthHandlers(g.Group("/auth"))
	realmHandler.UseRealmHandlers(g.Group("/realm"))

}
