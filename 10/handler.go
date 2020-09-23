package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateEmployee(c *gin.Context) {
	var employeee Employee
	if err := c.BindJSON(&employeee); err != nil {
		fmt.Printf("failed to bind empployee: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	h.storage.Insert(&employeee)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": employeee.ID,
	})
}

func (h *Handler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	var employee Employee
	if err != c.BindJSON(&employee); err != nil {
		fmt.Printf("failed to bind employee: %s\n", err.Error())
	}
}
