package cmd

import (
	"github.com/enjoypi/3rd/agent"
	"github.com/enjoypi/3rd/conf"
	"github.com/enjoypi/god"
	"github.com/enjoypi/god/services/mesh"
	"github.com/enjoypi/god/services/net"
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

	serveCmd.Flags().Int64("mesh.grantttl", 10, "ttl to etcd")
	serveCmd.Flags().String("mesh.defaulttimeout", "0", "ttl to etcd")
	serveCmd.Flags().String("mesh.path", "nodes", "listen address")
	serveCmd.Flags().StringSlice("etcd.endpoints", []string{"127.0.0.1:2379"}, "string flag for child")
	serveCmd.Flags().Bool("etcd.PermitWithoutStream", true, "ensures that the keepalive logic is running even without any active streams")

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
	logger.Sugar().Infof("mesh.Config:\n%+v", meshCfg)

	return node.AddService(conf.MeshService, mesh.NewService(meshCfg, logger))
}

func newNet(v *viper.Viper, logger *zap.Logger, node *god.Node) error {
	var cfg net.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}
	logger.Sugar().Infof("net.Config:\n%+v", cfg)

	svc := net.NewService(
		cfg,
		logger,
		(*agent.Manager)(nil),
		(*agent.StateActive)(nil),
	)

	return node.AddService(conf.AgentService, svc)
}
