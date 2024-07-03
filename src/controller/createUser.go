package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/validation"
	"github.com/pedrodecf/go-crud-study/src/controller/model/request"
	"github.com/pedrodecf/go-crud-study/src/model"
	"github.com/pedrodecf/go-crud-study/src/model/service"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("user created successfully",
		zap.String("journey", "createUser"))

		c.String(http.StatusOK, "")

}
