package router

import (
	authController "todolist/application/authentification/controller"
	"todolist/application/authentification/service"
	repoRedis "todolist/application/redis/repository"
	"todolist/application/simple/controller"
	repoUser "todolist/application/user/repository"
)

func SimpleInit() *controller.Controller {
	return controller.NewController()
}

func AuthInit() *authRouter {
	repoUser := repoUser.NewUserRepo()
	repoRedis := repoRedis.NewRedisRepository()
	serv := service.NewAuthService(repoUser, repoRedis)
	controll := authController.NewAuthController(serv)
	return NewAuthRouter(controll)
}
