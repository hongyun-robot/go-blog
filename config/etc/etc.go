package etc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ArticleRpc  zrpc.RpcClientConf
	ClassifyRpc zrpc.RpcClientConf
	//BaseRpc        zrpc.RpcClientConf
}