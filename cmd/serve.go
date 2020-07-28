package cmd

import (
	"github.com/enjoypi/3rd/agent"
	"github.com/enjoypi/god"
	"github.com/enjoypi/god/pb"
	"github.com/enjoypi/god/services/mesh"
	"github.com/enjoypi/god/services/net"
	"github.com/enjoypi/god/transports/message_bus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// doCmd represents the do command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "to serve",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	PreRunE: preRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		return serveRun(rootViper, logger)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String("mesh.advertiseaddress", "", "advertise address for grpc")
	serveCmd.Flags().String("mesh.defaulttimeout", "10s", "ttl to etcd")
	serveCmd.Flags().String("mesh.net.net.listenaddress", ":1119", "listen address for grpc")
	serveCmd.Flags().Int64("mesh.grantttl", 10, "ttl to etcd")
	serveCmd.Flags().String("mesh.path", "/nodes", "root path of mesh")
	serveCmd.Flags().Int("mesh.retrytimes", 10, "times to retry when dialing etcd")

	serveCmd.Flags().StringSlice("etcd.endpoints", []string{"127.0.0.1:2379"}, "endpoints of etcd")
	serveCmd.Flags().Bool("etcd.permitwithoutstream", true, "ensures that the keepalive logic is running even without any active streams")

	serveCmd.Flags().String("nats.url", "nats://127.0.0.1:4222", "nats url")
	//serveCmd.Flags().String("nats.readtimeout", "10s", "default ")

	serveCmd.Flags().String("net.listenaddress", "", "listen address")

	serveCmd.Flags().String("node.type", "default", "service type")
	serveCmd.Flags().Uint16("node.id", 0, "service type")
}

func serveRun(v *viper.Viper, logger *zap.Logger) error {
	var cfg god.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}
	logger.Sugar().Infof("god.Config:\n%+v", cfg)

	node, err := god.NewNode(&cfg, logger)
	if err != nil {
		return err
	}

	err = newMesh(v, logger, node)
	if err != nil {
		return err
	}
	//
	//if err := newMeshServer(v, logger, node); err != nil {
	//	return err
	//}

	if err := newMessageBus(v, logger, node); err != nil {
		return err
	}

	if err := newNet(v, logger, node); err != nil {
		return err
	}

	return node.Serve()
}

func newMesh(v *viper.Viper, logger *zap.Logger, node *god.Node) error {
	var meshCfg mesh.Config
	if err := v.Unmarshal(&meshCfg); err != nil {
		return err
	}
	logger.Sugar().Infof("mesh config:\n%+v", meshCfg)

	return node.AddService(pb.ServiceType_Mesh, mesh.NewService(meshCfg, logger, pb.TransportType_MessageBus))
}

//func newMeshServer(v *viper.Viper, logger *zap.Logger, node *god.Node) error {
//	var cfg mesh.Config
//	if err := v.Unmarshal(&cfg); err != nil {
//		return err
//	}
//	logger.Sugar().Infof("mesh server config:\n%+v", cfg.Mesh.Net)
//
//	svc := net.NewService(
//		cfg.Mesh.Net,
//		logger,
//		node,
//		(*mesh.Manager)(nil),
//		(*mesh.Session)(nil),
//	)
//
//	return node.AddService(pb.ServiceType_Receiver, svc)
//}
//
func newMessageBus(v *viper.Viper, logger *zap.Logger, node *god.Node) error {
	var cfg message_bus.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}
	logger.Sugar().Infof("NATS config:\n%+v", cfg)

	trans := message_bus.NewTransport(
		cfg,
		logger,
		node.ID,
	)

	return node.AddTransport(pb.TransportType_MessageBus, trans)
}

func newNet(v *viper.Viper, logger *zap.Logger, node *god.Node) error {
	var cfg net.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}
	logger.Sugar().Infof("agent config:\n%+v", cfg)

	svc := net.NewService(
		cfg,
		logger,
		node,
		(*agent.Manager)(nil),
		(*agent.StateActive)(nil),
	)

	return node.AddService(pb.ServiceType_Agent, svc)
}
