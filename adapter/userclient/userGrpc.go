package userclient

import (
	"github.com/golang/protobuf/ptypes"
	uspb "github.com/jfeng45/servicetmpl/adapter/userclient/generatedclient"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/pkg/errors"
)

// convert from grpc User type to domain Model user type
func GrpcToUser(user *uspb.User) (*model.User, error){
	resultUser := model.User{}

	resultUser.Id = int(user.Id)
	resultUser.Name = user.Name
	resultUser.Department = user.Department
	created, err := ptypes.Timestamp(user.Created)
	if err != nil {
		return nil, errors.Wrap(err,"")
	}
	resultUser.Created = created
	panic("test")
	return &resultUser, nil
}

// convert from domain Model User type to grpc user type
func UserToGrpc(user *model.User) (*uspb.User, error) {
	resultUser := uspb.User{}
	resultUser.Id = int32(user.Id)
	resultUser.Name = user.Name
	resultUser.Department = user.Department
	created, err := ptypes.TimestampProto (user.Created)
	if err != nil {
		return nil, errors.Wrap(err,"")
	}
	resultUser.Created = created
	return &resultUser, nil
}

// convert from array of domain Model User type to array of grpc user type
func UserListToGrpc(ul []model.User) ([]*uspb.User, error) {
	gul :=[]*uspb.User{}
	for _, user :=range ul {
		gu, err :=UserToGrpc(&user)
		if err != nil {
			return nil, errors.Wrap(err,"")
		}
		gul =append(gul, gu)
	}
	return gul, nil
}