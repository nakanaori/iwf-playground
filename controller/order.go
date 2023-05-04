package controller

import (
	"iwf-playground/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
)

type orderController struct {
	OrderService service.OrderService
}

func NewOrderController() *orderController {
	return &orderController{
		OrderService: service.NewOrderService(),
	}
}

func (ths *orderController) GetOrderByID(c *gin.Context) {
	orderId := c.Param("id")
	data, err := ths.OrderService.GetOrderByID(c.Request.Context(), orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, data)
}

func (ths *orderController) CheckoutOrder(c *gin.Context) {
	orderId := c.Param("id")
	runId, err := ths.OrderService.CheckoutOrder(c.Request.Context(), orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, "runId = ", runId)
}

func (ths *orderController) RejectOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	err := ths.OrderService.RejectOrder(c.Request.Context(), orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, "Hello %s", orderId)
}

func (ths *orderController) AcceptOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	err := ths.OrderService.AcceptOrder(c.Request.Context(), orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, "Hello %s", orderId)
}

func (ths *orderController) OrderInvokeStartHandler(c *gin.Context) {
	var req iwfidl.WorkflowStateStartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ths.OrderService.OrderInvokeStartHandler(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ths *orderController) OrderInvokeDecideHandler(c *gin.Context) {
	var req iwfidl.WorkflowStateDecideRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ths.OrderService.OrderInvokeDecideHandler(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
