package server

import (
	"iwf-playground/controller"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	controller := controller.NewOrderController()

	router.GET("/order/:id", controller.GetOrderByID)
	router.POST("/order/checkout/:id", controller.CheckoutOrder)
	router.POST("/order/reject/:id", controller.CheckoutOrder)
	router.POST("/order/accept/:id", controller.CheckoutOrder)
	router.POST(iwf.WorkflowStateStartApi, controller.OrderInvokeStartHandler)
	router.POST(iwf.WorkflowStateDecideApi, controller.OrderInvokeDecideHandler)
	return router
}
