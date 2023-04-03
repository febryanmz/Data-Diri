package delivery

import (
	"github.com/febryanmz/Data-Diri/features/profile"
	"github.com/gin-gonic/gin"
)

func NewEcho(service profile.ServiceInterface, e *gin.Engine) {
	handler := &ProfileDelivery{
		profileService: service,
	}
	e.POST("/profile", handler.Create)

}
