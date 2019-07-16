package main

import (
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/container/servicecontainer"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/jfeng45/servicetmpl/tools"
	"github.com/jfeng45/servicetmpl/usecase"
	"github.com/pkg/errors"
	"time"
)

const (
	DEV_CONFIG string = "../configs/appConfigDev.yaml"
	PROD_CONFIG string = "../configs/appConfigProd.yaml"
)
func main() {
	testMySql()
	//testCouchDB()
}

func getListUserUseCase(c container.Container) (usecase.ListUserUseCaseInterface, error){
	key := container.LIST_USER
	//value, found := registry.GetFromRegistry(sc.FactoryMap, key)
	value, err := c.GetInstance(key)
	if err!=nil  {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	if value !=nil {
		logger.Log.Debug("found list user use case: key=", key)
		return value.(usecase.ListUserUseCaseInterface), nil
	}
	return nil, nil
}

func getRegistrationUseCase(c container.Container) (usecase.RegistrationUseCaseInterface, error){
	key := container.REGISTRATION
	//value, found := registry.GetFromRegistry(sc.FactoryMap, key)
	value, err := c.GetInstance(key)
	if err!=nil  {
		//logger.Log.Errorf("%+v\n", err)
		return nil, errors.Wrap(err, "")
	}
	if value !=nil {
		logger.Log.Debug("found registration use case: key=", key)
		return value.(usecase.RegistrationUseCaseInterface), nil
	}
	return nil, nil
}

func buildContainer (filename string) (container.Container, error){
	factoryMap :=make(map[string]interface{})
	container := servicecontainer.ServiceContainer{factoryMap}

	err:= container.InitApp( filename)
	if err!=nil  {
		logger.Log.Errorf("%+v\n", err)
		return nil, err
	}
	return &container, nil
}
func testCouchDB() {
	filename := PROD_CONFIG
	container, err := buildContainer(filename)
	if err!=nil  {
		logger.Log.Errorf("%+v\n", err)
		return
	}
	testFindById(container)
}
func testMySql() {

	filename := DEV_CONFIG
	container, err := buildContainer(filename)
	if err!=nil  {
		logger.Log.Errorf("%+v\n", err)
		return
	}

	testListUser(container)
	//testListUser(container)
	testFindById(container)
	//testFindById(container)
	//testRegisterUser(container)
	//testModifyUser(container)
	//testUnregister(container)
	//
	//testModifyAndUnregister(container)
	//testModifyAndUnregisterWithTx(container)


}
func testUnregister(container container.Container) {

	ruci, err := getRegistrationUseCase(container)
	if err != nil {
		logger.Log.Fatal("registration interface build failed:%+v\n", err)
	}
	username := "Brian"
	err =ruci.UnregisterUser(username)
	if err != nil {
		logger.Log.Fatalf("testUnregister failed:%+v\n", err)
	}
	logger.Log.Infof("testUnregister successully")
}

func testRegisterUser(container container.Container) {
	ruci, err :=  getRegistrationUseCase(container)
	if err != nil {
		logger.Log.Fatal("registration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2018-12-09")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}

	user := model.User{Name: "Brian", Department:"Marketing", Created:created}

	resultUser, err :=ruci.RegisterUser(&user)
	if err != nil {
		logger.Log.Errorf("user registration failed:%+v\n", err)
	} else {
		logger.Log.Info("new user registered:", resultUser)
	}
}

func testModifyUser(container container.Container) {
	ruci, err :=  getRegistrationUseCase(container)
	if err != nil {
		logger.Log.Fatal("registration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2019-12-01")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}
	user := model.User{Id: 6, Name:"Aditi", Department:"HR", Created:created}
	err =ruci.ModifyUser(&user)
	if err !=nil {
		logger.Log.Infof("Modify user failed:%+v\n", err)
	} else {
		logger.Log.Info("user modified succeed:", user)
	}
}

func testListUser(container container.Container) {
	//rluf, err := container.RetrieveListUser()
	rluf, err := getListUserUseCase(container)
	if err != nil {
		logger.Log.Fatal("RetrieveListUser interface build failed:", err)
	}
	users, err := rluf.ListUser()
	if err != nil {
		logger.Log.Errorf("user list failed:%+v\n", err)
	}
	logger.Log.Info("user list:", users)
}


func testModifyAndUnregister(container container.Container) {
	ruci, err :=  getRegistrationUseCase(container)
	if err != nil {
		logger.Log.Fatal("RegisterRegistration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2018-12-09")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}
	user := model.User{Id: 6, Name:"Richard", Department:"Sales", Created:created}
	err = ruci.ModifyAndUnregister(&user)
	if err != nil {
		logger.Log.Errorf("ModifyAndUnregister failed:%+v\n", err)
	} else {
		logger.Log.Infof("ModifyAndUnregister succeed")
	}
}

func testModifyAndUnregisterWithTx(container container.Container) {
	ruci, err :=  getRegistrationUseCase(container)
	if err != nil {
		logger.Log.Fatal("RegisterRegistration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2018-12-09")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}
	user := model.User{Id: 8, Name:"Anshu", Department:"Sales", Created:created}
	err = ruci.ModifyAndUnregisterWithTx(&user)
	if err != nil {
		logger.Log.Errorf("ModifyAndUnregisterWithTx failed:%+v\n", err)
	} else {
		logger.Log.Infof("ModifyAndUnregisterWithTx succeed")
	}
}

func testFindById(container container.Container) {
	//It is uid in database. Make sure you have it in database, otherwise it won't find it.
	id :=11
	//rluf, err := container.RetrieveListUser()
	rluf, err := getListUserUseCase(container)
	if err != nil {
		logger.Log.Fatalf("RetrieveListUser interface build failed:%+v\n", err)
	}
	user, err := rluf.Find(id)
	if err != nil {
		logger.Log.Errorf("fin user failed failed:%+v\n", err)
	}
	logger.Log.Info("find user:", user)
}
