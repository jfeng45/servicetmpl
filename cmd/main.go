package main

import (
	"github.com/jfeng45/servicetmpl/appcontainer"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/jfeng45/servicetmpl/tools"
	"github.com/jfeng45/servicetmpl/tools/logger"
	"time"
)

func main() {
	testMySql()
	//testCouchDB()
}
func testCouchDB() {
	factoryMap :=make(map[string]interface{})

	filename := "../configs/appConfigProd.yaml"
	err:= appcontainer.InitApp(factoryMap, filename)
	if err!=nil  {
		logger.Log.Errorf("%+v\n", err)
	}
	testFindById(factoryMap)
}
func testMySql() {
	factoryMap :=make(map[string]interface{})

	filename := "../configs/appConfigDev.yaml"
	err:= appcontainer.InitApp(factoryMap, filename)
	if err!=nil  {
		logger.Log.Errorf("%+v\n", err)
	}

	//testListUser(factoryMap)
	//testFindById(factoryMap)
	//testRegisterUser(factoryMap)
	//testModifyUser(factoryMap)
	//testUnregister(factoryMap)

	//testModifyAndUnregister(factoryMap)
	testModifyAndUnregisterWithTx(factoryMap)
	//testTransaction(factoryMap)

}
func testUnregister(factoryMap map[string]interface{}) {

	ruci, err := appcontainer.RetrieveRegistration(factoryMap)
	if err != nil {
		logger.Log.Fatal("registration interface build failed:%+v\n", err)
	}
	username := "Richard"
	//username := "Peter"
	err =ruci.UnregisterUser(username)
	if err != nil {
		logger.Log.Fatalf("testUnregister failed:%+v\n", err)
	}
	logger.Log.Infof("testUnregister successully")
}

func testRegisterUser(factoryMap map[string]interface{}) {
	ruci, err := appcontainer.RetrieveRegistration(factoryMap)
	if err != nil {
		logger.Log.Fatal("registration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2018-12-09")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}
	//user := usermodel.User{Name: "Richard", Department:"finance", Created:created}
	//user := usermodel.User{Name: "Tony", Department:"IT", Created:created}
	user := model.User{Name: "Anshu", Department:"IT", Created:created}
	//created = time.Time{}
	//user := usermodel.User{Name: "", Department:"IT", Created:created}

	resultUser, err :=ruci.RegisterUser(&user)
	if err != nil {
		logger.Log.Errorf("user registration failed:%+v\n", err)
	} else {
		logger.Log.Info("new user registered:", resultUser)
	}
}

func testModifyUser(factoryMap map[string]interface{}) {
	ruci, err := appcontainer.RetrieveRegistration(factoryMap)
	if err != nil {
		logger.Log.Fatal("registration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2019-12-01")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}
	user := model.User{Id: 3, Name:"Brian", Department:"HR", Created:created}
	//user := usermodel.User{Name:"Brian", Department:"HR", Created:created}
	err =ruci.ModifyUser(&user)
	if err !=nil {
		logger.Log.Infof("Modify user failed:%+v\n", err)
	} else {
		logger.Log.Info("user modified succeed:", user)
	}
}

func testListUser(factoryMap map[string]interface{}) {
	rluf, err := appcontainer.RetrieveListUser(factoryMap)
	if err != nil {
		logger.Log.Fatal("RetrieveListUser interface build failed:", err)
	}
	users, err := rluf.ListUser()
	if err != nil {
		logger.Log.Errorf("user list failed:%+v\n", err)
	}
	logger.Log.Info("user list:", users)
}

//func testTransaction(factoryMap map[string]interface{}) {
//	ruci, err := appcontainer.RetrieveRegistration(factoryMap)
//	if err != nil {
//		logger.Log.Fatal("RegisterRegistration interface build failed:%+v\n", err)
//	}
//	created, err := time.Parse(timea.FORMAT_ISO8601_DATE,"2018-12-09")
//	if err != nil {
//		logger.Log.Errorf("date format err:%+v\n", err)
//	}
//	user :=model.User{Name:"Anshu", Department:"Sales", Created:created}
//	ruci.Transaction(&user)
//	if err != nil {
//		logger.Log.Errorf("transaction failed:%+v\n", err)
//	} else {
//		logger.Log.Infof("transaction succeed")
//	}
//}

func testModifyAndUnregister(factoryMap map[string]interface{}) {
	ruci, err := appcontainer.RetrieveRegistration(factoryMap)
	if err != nil {
		logger.Log.Fatal("RegisterRegistration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2018-12-09")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}
	//user := usermodel.User{Name:"Anshu", Department:"Sales", Created:created}
	user := model.User{Id: 4, Name:"Peter", Department:"Sales", Created:created}
	//user :=model.User{Id:2, Name:"Aditi", Department:"Marketing", Created:created}
	err = ruci.ModifyAndUnregister(&user)
	if err != nil {
		logger.Log.Errorf("ModifyAndUnregister failed:%+v\n", err)
	} else {
		logger.Log.Infof("ModifyAndUnregister succeed")
	}
}

func testModifyAndUnregisterWithTx(factoryMap map[string]interface{}) {
	ruci, err := appcontainer.RetrieveRegistration(factoryMap)
	if err != nil {
		logger.Log.Fatal("RegisterRegistration interface build failed:%+v\n", err)
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE,"2018-12-09")
	if err != nil {
		logger.Log.Errorf("date format err:%+v\n", err)
	}
	user := model.User{Id: 3, Name:"Anshu", Department:"Sales", Created:created}
	err = ruci.ModifyAndUnregisterWithTx(&user)
	if err != nil {
		logger.Log.Errorf("ModifyAndUnregisterWithTx failed:%+v\n", err)
	} else {
		logger.Log.Infof("ModifyAndUnregisterWithTx succeed")
	}
}

func testFindById(factoryMap map[string]interface{}) {
	//It is uid in database. Make sure you have it in database, otherwise it won't find it.
	id :=12
	rluf, err := appcontainer.RetrieveListUser(factoryMap)
	if err != nil {
		logger.Log.Fatalf("RetrieveListUser interface build failed:%+v\n", err)
	}
	user, err := rluf.Find(id)
	if err != nil {
		logger.Log.Errorf("fin user failed failed:%+v\n", err)
	}
	logger.Log.Info("find user:", user)
}
