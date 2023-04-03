package delivery

import (
	"github.com/febryanmz/Data-Diri/features/profile"
	"github.com/gin-gonic/gin"
)

type ProfileDelivery struct {
	profileService profile.ServiceInterface
}

func New(service profile.ServiceInterface, e *gin.Engine) {
	handler := &ProfileDelivery{
		profileService: service,
	}
	e.POST("/profile", handler.Create)

}

func (delivery *ProfileDelivery) Create(c *gin.Context) {

}
