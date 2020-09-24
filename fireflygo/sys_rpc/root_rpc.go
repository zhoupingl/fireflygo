package sys_rpc

import (
	"github.com/sc4599/fireflygo/cluster"

	"github.com/sc4599/fireflygo/clusterserver"
	"github.com/sc4599/fireflygo/logger"
	"github.com/sc4599/fireflygo/utils"
)

type RootRpc struct {
}

/*
子节点连上来的通知
*/
func (this *RootRpc) TakeProxy(request *cluster.RpcRequest) {
	name := request.Rpcdata.Args[0].(string)
	logger.Info("child node " + name + " connected to " + utils.GlobalObject.Name)
	//加到childs并且绑定链接connetion对象
	clusterserver.GlobalClusterServer.AddChild(name, request.Fconn)
}
