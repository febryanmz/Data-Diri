//package delivery
//
//import (
//	"github.com/febryanmz/Data-Diri/features/profile"
//	"github.com/labstack/echo/v4"
//)
//
//func NewEcho(service profile.ServiceInterface, e *echo.Echo) {
//	handler := &ProfileDelivery{
//		profileService: service,
//	}
//	e.POST("/profile", handler.CreateEcho)
//	e.GET("/profile/:id", handler.ReadEcho)
//	e.PUT("/profile/:id", handler.UpdateEcho)
//	e.DELETE("/profile/:id", handler.DeleteEcho)
//}
//
//func (delivery *ProfileDelivery) CreateEcho(c *echo.Context) {
//
//}
//func (delivery *ProfileDelivery) ReadEcho(c *echo.Context) {
//
//}
//func (delivery *ProfileDelivery) UpdateEcho(c *echo.Context) {
//
//}
//func (delivery *ProfileDelivery) DeleteEcho(c *echo.Context) {
//
//}
