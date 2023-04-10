package delivery

import (
	"github.com/febryanmz/Data-Diri/features/profile"
	"github.com/febryanmz/Data-Diri/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileDelivery struct {
	profileService profile.ServiceInterface
}

func NewGin(service profile.ServiceInterface, e *gin.Engine) {
	handler := &ProfileDelivery{
		profileService: service,
	}
	e.POST("/profile", handler.CreateGin)
	e.GET("/profile/:id", handler.ReadGin)
	e.PUT("/profile/:id", handler.UpdateGin)
	e.DELETE("/profile/:id", handler.DeleteGin)

}

func (delivery *ProfileDelivery) CreateGin(c *gin.Context) {
	profileInput := profile.Core{}
	errBind := c.Bind(&profileInput)
	if errBind != nil {
		c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	err := delivery.profileService.Create(profileInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	c.IndentedJSON(http.StatusOK, helper.SuccessResponse("success create new users"))
}

//func (delivery *ProfileDelivery) ReadGin(c *gin.Context) {
//	id, errConv := strconv.Atoi(c.Param("id"))
//	if errConv != nil {
//		c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
//	}
//
//	results, err := delivery.profileService.Read(id)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
//	}
//
//	c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read data profile by id", results))
//}

//func (delivery *ProfileDelivery) UpdateGin(c *gin.Context) {
//	id, errConv := strconv.Atoi(c.Param("id"))
//	if errConv != nil {
//		c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
//	}
//
//	profileInput := profile.Core{}
//	errBind := c.Bind(&profileInput)
//	if errBind != nil {
//		c.JSON(http.StatusBadRequest, helper.FailedResponse("Error bind data "+errBind.Error()))
//	}
//
//	errUpt := delivery.profileService.Update(profileInput, id)
//	if errUpt != nil {
//		c.JSON(http.StatusBadRequest, helper.FailedResponse("Error update database"+errUpt.Error()))
//	}
//	c.JSON(http.StatusOK, helper.SuccessResponse("Success update data"))
//}
//
//func (delivery *ProfileDelivery) DeleteGin(c *gin.Context) {
//	id, errConv := strconv.Atoi(c.Param("id"))
//	if errConv != nil {
//		c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
//	}
//
//	errDel := delivery.profileService.Delete(id)
//	if errDel != nil {
//		c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete user"+errDel.Error()))
//	}
//
//	c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
//
//}
