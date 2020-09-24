package fireflygo

import (
	"fmt"
	"github.com/sc4599/fireflygo/cluster"
	"github.com/sc4599/fireflygo/clusterserver"
	_ "github.com/sc4599/fireflygo/fnet"
	"github.com/sc4599/fireflygo/fserver"
	"github.com/sc4599/fireflygo/iface"
	"github.com/sc4599/fireflygo/logger"
	"github.com/sc4599/fireflygo/sys_rpc"
	"github.com/sc4599/fireflygo/telnetcmd"
	_ "github.com/sc4599/fireflygo/timer"
	"github.com/sc4599/fireflygo/utils"
)

func NewfireflygoTcpServer() iface.Iserver {
	//do something
	//debugport 是否开放
	if utils.GlobalObject.DebugPort > 0 {
		if utils.GlobalObject.Host != "" {
			fserver.NewTcpServer("telnet_server", "tcp4", utils.GlobalObject.Host,
				utils.GlobalObject.DebugPort, 100, cluster.NewTelnetProtocol()).Start()
		} else {
			fserver.NewTcpServer("telnet_server", "tcp4", "127.0.0.1",
				utils.GlobalObject.DebugPort, 100, cluster.NewTelnetProtocol()).Start()
		}
		logger.Debug(fmt.Sprintf("telnet tool start: %s:%d.", utils.GlobalObject.Host, utils.GlobalObject.DebugPort))

	}

	//add command
	if utils.GlobalObject.CmdInterpreter != nil {
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewPprofCpuCommand())
	}

	s := fserver.NewServer()
	return s
}

func NewfireflygoMaster(cfg string) *clusterserver.Master {
	s := clusterserver.NewMaster(cfg)
	//add rpc
	s.AddRpcRouter(&sys_rpc.MasterRpc{})
	//add command
	if utils.GlobalObject.CmdInterpreter != nil {
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewPprofCpuCommand())
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewCloseServerCommand())
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewReloadCfgCommand())
	}
	return s
}

func NewfireflygoCluterServer(nodename, cfg string) *clusterserver.ClusterServer {
	s := clusterserver.NewClusterServer(nodename, cfg)
	//add rpc
	s.AddRpcRouter(&sys_rpc.ChildRpc{})
	s.AddRpcRouter(&sys_rpc.RootRpc{})
	//add cmd
	if utils.GlobalObject.CmdInterpreter != nil {
		utils.GlobalObject.CmdInterpreter.AddCommand(telnetcmd.NewPprofCpuCommand())
	}
	return s
}
