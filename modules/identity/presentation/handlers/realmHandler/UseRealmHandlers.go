package realmHandler

import (
	"StellaRP/modules/identity/presentation/contracts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UseRealmHandlers(r *gin.RouterGroup) {
	r.POST("", createRealm)
}

func createRealm(ctx *gin.Context) {

	body := &contracts.RealmCreateRequest{}

	if err := ctx.BindJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
