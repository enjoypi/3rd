package cmd

import (
	"github.com/enjoypi/god"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var agentCommand = &cobra.Command{
	Use:   "agent",
	Short: "agent server for users",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	PreRunE: preRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := god.Initialize(rootViper); err != nil {
			return err
		}
		god.Wait()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(agentCommand)

	flags := agentCommand.Flags()

	flags.String("mesh.advertiseaddress", "", "advertise address for grpc")
	flags.String("mesh.defaulttimeout", "10s", "ttl to etcd")
	flags.Int64("mesh.grantttl", 10, "ttl to etcd")
	flags.String("mesh.path", "/nodes", "root path of mesh")
	flags.Int("mesh.retrytimes", 10, "times to retry when dialing etcd")

	flags.StringSlice("etcd.endpoints", []string{"127.0.0.1:2379"}, "endpoints of etcd")
	flags.Bool("etcd.permitwithoutstream", true, "ensures that the keepalive logic is running even without any active streams")

	flags.String("nats.url", "nats://127.0.0.1:4222", "nats url")
	//serveCmd.Flags().String("nats.readtimeout", "10s", "default ")

	flags.String("socket.listenaddress", "127.0.0.1:1119", "listen address")

	flags.StringArray("node.apps", nil, "service type")
	flags.String("node.type", "default", "service type")
	flags.Uint16("node.id", 0, "service type")
}
