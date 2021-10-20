package config

import (
	"context"

	"github.com/microgate-io/microgate-lib-go/v1/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetConfig(microgateConn *grpc.ClientConn) (*Configuration, error) {
	client := NewConfigServiceClient(microgateConn)
	ctx := context.Background()
	resp, err := client.GetConfig(ctx, new(emptypb.Empty))
	if err != nil {
		return nil, log.ErrorWithLog(ctx, err, "failed to get configuration")
	}
	log.Infow(ctx, "fetched configuration", "config", resp.Config)
	return resp.Config, nil
}

func (c *Configuration) FindBool(name string) bool {
	for key, each := range c.Entries {
		if key == name {
			return each.GetBoolValue()
		}
	}
	return false
}
