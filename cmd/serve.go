package cmd

import (
	"github.com/enjoypi/3rd/player"
	"github.com/enjoypi/god"
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

	serveCmd.Flags().StringArray("etcd.endpoints", []string{"127.0.0.1:2379"}, "string flag for child")
	serveCmd.Flags().Int64("etcdttl", 10, "ttl to etcd")
	serveCmd.Flags().String("nodepath", "nodes", "path of nodes info")
	serveCmd.Flags().String("node.type", "default", "service type")
	serveCmd.Flags().Uint16("node.id", 0, "service type")
	serveCmd.Flags().String("listenaddress", "", "listen address")
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

	srv, err := node.NewService(1, &player.Manager{})
	if err != nil {
		return err
	}
	return srv.Serve(v.GetString("listenaddress"))
}
