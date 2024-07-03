package service


import (
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("init deleteUser model", zap.String("journey", "deleteUser"))

	return nil
}