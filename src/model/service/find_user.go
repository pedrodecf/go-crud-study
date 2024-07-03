package service

import (
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/rest_err"
	"github.com/pedrodecf/go-crud-study/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(id string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init findUser model", zap.String("journey", "findUser"))

	return nil, nil
}
