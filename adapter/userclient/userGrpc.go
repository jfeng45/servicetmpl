// Package userclient is client library if you need to call the user Micro-service as a client.
// It provides client library and the data transformation service.
package userclient

import (
	"github.com/golang/protobuf/ptypes"
	uspb "github.com/jfeng45/servicetmpl/adapter/userclient/generatedclient"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/pkg/errors"
)

// GrpcToUser converts from grpc User type to domain Model user type
func GrpcToUser(user *uspb.User) (*model.User, error) {
	if user == nil {
		//logger.Log.Debug("user is nil")
		return nil, nil
	}
	resultUser := model.User{}

	resultUser.Id = int(user.Id)
	resultUser.Name = user.Name
	resultUser.Department = user.Department
	created, err := ptypes.Timestamp(user.Created)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	resultUser.Created = created
	return &resultUser, nil
}

// UserToGrpc converts from domain Model User type to grpc user type
func UserToGrpc(user *model.User) (*uspb.User, error) {
	if user == nil {
		//logger.Log.Debug("user is nil")
		return nil, nil
	}
	resultUser := uspb.User{}
	resultUser.Id = int32(user.Id)
	resultUser.Name = user.Name
	resultUser.Department = user.Department
	created, err := ptypes.TimestampProto(user.Created)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	resultUser.Created = created
	return &resultUser, nil
}

// UserListToGrpc converts from array of domain Model User type to array of grpc user type
func UserListToGrpc(ul []model.User) ([]*uspb.User, error) {
	var gul []*uspb.User
	for _, user := range ul {
		gu, err := UserToGrpc(&user)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}
		gul = append(gul, gu)
	}
	return gul, nil
}
