package config

import (
	"context"
	stdlog "log"

	"github.com/microgate-io/microgate-lib-go/v1/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func GetConfig(microgateConn *grpc.ClientConn) *Configuration {
	client := NewConfigServiceClient(microgateConn)
	ctx := context.Background()
	resp, err := client.GetConfig(ctx, new(emptypb.Empty))
	if err != nil {
		stdlog.Fatalln("unable to fetch configuration", err)

	}
	log.Infow(ctx, "fetched configuration", "config", resp.Config)
	return resp.Config
}

func (c *Configuration) FindBool(name string, absent ...bool) bool {
	for key, each := range c.Entries {
		if key == name {
			return each.GetBoolValue()
		}
	}
	if len(absent) > 0 {
		return absent[0]
	}
	return false
}

func (c *Configuration) FindInt(name string, absent ...int) int {
	for key, each := range c.Entries {
		if key == name {
			return int(each.GetIntValue())
		}
	}
	if len(absent) > 0 {
		return absent[0]
	}
	return 0
}

func (c *Configuration) FindString(name string, absent ...string) string {
	for key, each := range c.Entries {
		if key == name {
			return each.GetStringValue()
		}
	}
	if len(absent) > 0 {
		return absent[0]
	}
	return ""
}
