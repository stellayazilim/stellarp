package realmHandler

import (
	realmfeature "StellaRP/modules/identity/application/features/realm"
	"StellaRP/modules/identity/presentation/contracts"
	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
	"net/http"
)

func UseRealmHandlers(r *gin.RouterGroup) {
	r.POST("", createRealm)
}

func createRealm(ctx *gin.Context) {

	body := &contracts.RealmCreateRequest{}

	if err := ctx.ShouldBindBodyWithJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	command := &realmfeature.CreateRealmCommand{
		Name: body.Name,
	}

	_, _ = mediatr.Send[*realmfeature.CreateRealmCommand, *realmfeature.CreateRealmResponse](ctx, command)

}
