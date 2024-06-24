package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/validation"
	"github.com/pedrodecf/go-crud-study/src/controller/model/request"
	"github.com/pedrodecf/go-crud-study/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("init createUser controller",
		zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("error trying to validate user info", err,
			zap.String("journey", "createUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("user created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
