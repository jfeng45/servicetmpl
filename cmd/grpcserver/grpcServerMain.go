package main

import (
	"context"
	"github.com/jfeng45/servicetmpl/adapter/userclient"
	uspb "github.com/jfeng45/servicetmpl/adapter/userclient/generatedclient"
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/container/servicecontainer"
	"github.com/jfeng45/servicetmpl/usecase"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

const (
	DEV_CONFIG  string = "../../config/appConfigDev.yaml"
	PROD_CONFIG string = "../../config/appConfigProd.yaml"
)

type UserService struct {
	container container.Container
}

func catchPanic() {
	if p := recover(); p != nil {
		logger.Log.Errorf("%+v\n", p)
	}
}

func (uss *UserService) RegisterUser(ctx context.Context, req *uspb.RegisterUserReq) (*uspb.RegisterUserResp, error) {
	defer catchPanic()
	logger.Log.Debug("RegisterUser called")

	ruci, err := getRegistrationUseCase(uss.container)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	mu, err := userclient.GrpcToUser(req.User)

	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("mu:", mu)
	resultUser, err := ruci.RegisterUser(mu)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("resultUser:", resultUser)
	gu, err := userclient.UserToGrpc(resultUser)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}

	logger.Log.Debug("user registered: ", gu)

	return &uspb.RegisterUserResp{User: gu}, nil

}

func (uss *UserService) ListUser(ctx context.Context, in *uspb.ListUserReq) (*uspb.ListUserResp, error) {
	defer catchPanic()
	logger.Log.Debug("ListUser called")

	luci, err := getListUserUseCase(uss.container)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}

	lu, err := luci.ListUser()
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	gu, err := userclient.UserListToGrpc(lu)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}

	logger.Log.Debug("user list: ", gu)

	return &uspb.ListUserResp{User: gu}, nil

}
func runServer(sc *servicecontainer.ServiceContainer) error {
	logger.Log.Debug("start runserver")

	srv := grpc.NewServer()

	cs := &UserService{sc}
	uspb.RegisterUserServiceServer(srv, cs)
	//l, err:=net.Listen(GRPC_NETWORK, GRPC_ADDRESS)
	ugc := sc.AppConfig.UserGrpcConfig
	logger.Log.Debugf("userGrpcConfig: %+v\n", ugc)
	l, err := net.Listen(ugc.DriverName, ugc.UrlAddress)
	if err != nil {
		return errors.Wrap(err, "")
	} else {
		logger.Log.Debug("server listening")
	}
	return srv.Serve(l)
}

func main() {
	filename := DEV_CONFIG
	//filename := PROD_CONFIG
	container, err := buildContainer(filename)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		panic(err)
	}
	if err := runServer(container); err != nil {
		logger.Log.Errorf("Failed to run user server: %+v\n", err)
		panic(err)
	} else {
		logger.Log.Info("server started")
	}
}
func buildContainer(filename string) (*servicecontainer.ServiceContainer, error) {

	factoryMap := make(map[string]interface{})
	config := config.AppConfig{}
	container := servicecontainer.ServiceContainer{factoryMap, &config}

	err := container.InitApp(filename)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &container, nil
}

func getListUserUseCase(c container.Container) (usecase.ListUserUseCaseInterface, error) {
	key := container.LIST_USER
	value, err := c.BuildUseCase(key)
	if err != nil {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.ListUserUseCaseInterface), nil
}

func getListCourseUseCase(c container.Container) (usecase.ListCourseUseCaseInterface, error) {
	key := container.LIST_COURSE
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.ListCourseUseCaseInterface), nil

}

func getRegistrationUseCase(c container.Container) (usecase.RegistrationUseCaseInterface, error) {
	key := container.REGISTRATION
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.RegistrationUseCaseInterface), nil

}
