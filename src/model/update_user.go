package model

import (
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) UpdateUser(id string) *rest_err.RestErr {
	logger.Info("init updateUser model", zap.String("journey", "updateUser"))

	return nil
}