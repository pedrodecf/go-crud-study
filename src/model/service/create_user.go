package service

import (
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/rest_err"
	"github.com/pedrodecf/go-crud-study/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	return nil
}
