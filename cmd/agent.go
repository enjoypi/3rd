package cmd

import (
	//"github.com/enjoypi/3rd/agent"
	"github.com/enjoypi/god"
	"github.com/enjoypi/god/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// doCmd represents the do command
var agentCommand = &cobra.Command{
	Use:   "agent",
	Short: "agent server for users",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	PreRunE: preRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		return serveRun(rootViper, logger)
	},
}

func init() {
	rootCmd.AddCommand(agentCommand)
	agentCommand.Flags().String("mesh.advertiseaddress", "", "advertise address for grpc")
	agentCommand.Flags().String("mesh.defaulttimeout", "10s", "ttl to etcd")
	agentCommand.Flags().String("mesh.net.net.listenaddress", ":1119", "listen address for grpc")
	agentCommand.Flags().Int64("mesh.grantttl", 10, "ttl to etcd")
	agentCommand.Flags().String("mesh.path", "/nodes", "root path of mesh")
	agentCommand.Flags().Int("mesh.retrytimes", 10, "times to retry when dialing etcd")

	agentCommand.Flags().StringSlice("etcd.endpoints", []string{"127.0.0.1:2379"}, "endpoints of etcd")
	agentCommand.Flags().Bool("etcd.permitwithoutstream", true, "ensures that the keepalive logic is running even without any active streams")

	agentCommand.Flags().String("nats.url", "nats://127.0.0.1:4222", "nats url")
	//serveCmd.Flags().String("nats.readtimeout", "10s", "default ")

	agentCommand.Flags().String("net.listenaddress", "", "listen address")

	agentCommand.Flags().String("node.type", "default", "service type")
	agentCommand.Flags().Uint16("node.id", 0, "service type")
}

func serveRun(v *viper.Viper, logger *zap.Logger) error {
	var cfg god.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}
	logger.Info("god.Config", zap.Any("config", cfg))

	//god.PanicOnError(core.StartService(agent.NewService()))
	return core.Serve()
}
