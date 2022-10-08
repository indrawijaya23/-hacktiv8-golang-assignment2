package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

// GetAllOrders godoc
// @Summary List orders
// @Description List all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Router /orders [get]
func GetAllOrders(ctx *gin.Context) {
	db := database.GetDB()
	var orders []models.Order

	err = db.Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// CreaterOrder godoc
// @Summary Create order
// @Description Create new order
// @Tags orders
// @Param order body models.Order true "Order Detail"
// @Success 200 {object} models.Order
// @Accept json
// @Produce json
// @Router /orders [POST]
func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var order models.Order

	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//newOrder.OrdererdAt = time.Now().Local()

	err = db.Create(&order).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

// UpdateOrderById godoc
// @Summary Update order
// @Description Update order by orderId
// @Tags orders
// @Param orderId path int true "Order ID"
// @Param order body models.Order true "Order Detail"
// @Success 200
// @Accept json
// @Produce json
// @Router /orders/{orderId} [PUT]
func UpdateOrderById(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	db := database.GetDB()
	order := models.Order{}
	item := models.Item{}

	orderIDint, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	order.ID = orderIDint

	err = db.Where("order_id = ?", orderID).Delete(&item).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Model(&order).Where("id = ?", orderID).Updates(order).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orderId": orderIDint,
	})
}

// DeleteOrderById godoc
// @Summary Delete order
// @Description Delete order by orderId
// @Tags orders
// @Param orderId path int true "Order ID"
// @Success 200
// @Accept json
// @Produce json
// @Router /orders/{orderId} [delete]
func DeleteOrderById(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	db := database.GetDB()
	order := models.Order{}
	item := models.Item{}

	orderIDint, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Where("order_id = ?", orderID).Delete(&item).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Where("id = ?", orderID).Delete(&order).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orderId": orderIDint,
	})
}
