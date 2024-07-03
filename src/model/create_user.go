package model

import (
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("init createUser model", zap.String("journey", "createUser"))
	ud.EncryptPassword()
	return nil
}
