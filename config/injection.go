package config

import (
	"login-ports/app/controller"
	"login-ports/lib/log"

	userUsecase "login-ports/domain/user/usecases"

	userRepoMysql "login-ports/domain/user/repositories/mysql"
)

type Injection struct {
	Logging log.ILogs

	UserController *controller.UserController
}

func NewInjection() Injection {
	logging := log.NewLog()

	db, err := NewMysql()
	if err != nil {
	}

	_userRepoMysql := userRepoMysql.NewMysqlUserRepo(db, logging)

	_userUsecase := userUsecase.NewUserUsecase(_userRepoMysql, logging)

	UserController := controller.NewUserController(_userUsecase, logging)

	return Injection{
		UserController: UserController,

		Logging: logging,
	}
}
