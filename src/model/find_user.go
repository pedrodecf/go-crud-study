package model

import (
	"github.com/pedrodecf/go-crud-study/src/configuration/logger"
	"github.com/pedrodecf/go-crud-study/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) FindUser(id string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("init findUser model", zap.String("journey", "findUser"))
	
	return nil, nil
}
