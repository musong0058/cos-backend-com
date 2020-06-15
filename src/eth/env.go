package eth

import (
	"cos-backend-com/src/common/proto"

	"cos-backend-com/src/common/providers/session"
)

var (
	Env = &Config{}
)

type Config struct {
	proto.CommonEnvConfig

	AdminAccessKey    string `conf:"admin_access_key"`
	AdminAccessSecret string `conf:"admin_access_secret"`

	Session session.SessionConfig `conf:"session"`

	// 内部服务
	Service proto.ServiceEndpoint `conf:"service"`

	EthClient struct {
		EndPoint  string `conf:"end_point"`
		InfuraKey string `conf:"infura_key"`
	} `conf:"eth"`
}
