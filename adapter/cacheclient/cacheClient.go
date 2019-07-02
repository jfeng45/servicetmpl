// Package cacheclient is the wrapper around the thrid party gRPC Cache Microservice.
// It encapsulates the logic to call outside service, to make it transparent to the business logic layer.

package cacheclient

import (
	"context"
	pb "github.com/jfeng45/servicetmpl/adapter/cacheclient/generatedclient"
	"github.com/jfeng45/servicetmpl/tools"
	"google.golang.org/grpc"
)

// CacheDataGrpc represents the gRPC connection handler
type CacheDataGrpc struct {
	Conn grpc.ClientConn
}

// getCacheClient creates a gRPC client
func getCacheClient(conn grpc.ClientConn) pb.CacheServiceClient {
	return pb.NewCacheServiceClient(&conn)
}

// Get handles call to Get function on Cache service
func (cdg CacheDataGrpc) Get(key string) ([]byte, error) {
	cacheClient := getCacheClient(cdg.Conn)
	resp, err := cacheClient.Get(context.Background(), &pb.GetReq{Key: key})
	if err != nil {
		return nil, err
	} else {
		return resp.Value, err
	}
}

// Store handles call to Store function on Cache service
func (cdg CacheDataGrpc) Store(key string, value []byte) error {
	cacheClient := getCacheClient(cdg.Conn)
	ctx := context.Background()
	_, err:= cacheClient.Store(ctx, &pb.StoreReq{Key:key,Value:value})

	if err != nil {
		return err
	} else {
		tools.Log.Debug("store called")
	}
	return nil
}
