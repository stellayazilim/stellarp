package realmHandler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRealmShouldCreateNewRealm(t *testing.T) {

	g := NewGomegaWithT(t)
	r := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(r)
	req, _ := http.NewRequest("POST", "/realm", bytes.NewBuffer(
		[]byte("{\"name\": \"stella\"}")))
	ctx.Request = req

	createRealm(ctx)

	g.Expect(r.Code).To(Equal(http.StatusOK))
}

func TestCreateRealmShouldFailToCreateRealmWithInvalidBody(t *testing.T) {
	g := NewGomegaWithT(t)
	bodies := []string{
		"{}",
		"",
		"{\"name\": \"\"}",
		"{\"name\": \"ss\"}",
		"{\"name\": \"ssssssssssssssssssssssssssssssss\"}",
	}

	for _, body := range bodies {
		r := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(r)
		req, _ := http.NewRequest("POST", "/realm", bytes.NewBuffer([]byte(body)))
		ctx.Request = req
		createRealm(ctx)
		g.Expect(r.Code).To(Equal(http.StatusBadRequest))

	}
}
