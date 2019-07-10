package userclient

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	uspb "github.com/jfeng45/servicetmpl/adapter/userclient/generatedclient"
	"github.com/jfeng45/servicetmpl/model"
)

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

	return &resultUser, nil
}
