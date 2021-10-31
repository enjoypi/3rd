package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	configFile           string
	configRemoteEndpoint string
	configRemotePath     string
	configType           string
	rootViper            = viper.New()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "3rd",
	Short: "the template of cobra",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application`,

	PreRunE: preRunE,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&configFile, "config.file", "c", "", "config file")

	rootCmd.PersistentFlags().StringVar(&configRemoteEndpoint,
		"config.remote.endpoint",
		"", //"127.0.0.1:2379",
		"the endpoint of remote config")
	rootCmd.PersistentFlags().StringVar(&configRemotePath,
		"config.remote.path",
		"3rd/config",
		"the path of remote config")

	rootCmd.PersistentFlags().StringVar(&configType, "config.type", "yaml", "the type of config format")
	rootCmd.PersistentFlags().BoolP("verbose", "V", false, "verbose")

	rootCmd.PersistentFlags().String("log.level", "info", "level of logger")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().Bool("version", false, "show version")
}

func preRunE(cmd *cobra.Command, args []string) (err error) {
	// Viper uses the following precedence order. Each item takes precedence over the item below it:
	//
	// explicit call to Set
	// flag
	// env
	// config
	// key/value store
	// default
	//
	// Viper configuration keys are case insensitive.
	v := rootViper
	v.SetConfigType(configType)

	// remote config, key/value store
	initRemoteConfig(v)

	// local config
	if configFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(configFile)

		// If a config file is found, read it in.
		if err := v.ReadInConfig(); err != nil {
			return err
		}
	}

	// env
	v.AutomaticEnv() // read in environment variables that match

	// flag
	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	// log level in flags maybe wrong, reset
	//if lvl := zap.LevelFlag(v.GetString("log.level"), zap.InfoLevel, ""); lvl != nil {
	//	logger = logger.WithOptions(zap.NewAtomicLevel().SetLevel(zap.DebugLevel))
	//	sugar.Warn("current log level: ", lvl.String())
	//}
	if out, err := yaml.Marshal(v.AllSettings()); err == nil {
		fmt.Print(string(out))
	} else {
		fmt.Print("all settings:", v.AllSettings())
	}

	return nil
}

func initRemoteConfig(v *viper.Viper) {
	//if configRemoteEndpoint == "" {
	//	return
	//}
	//
	//sugar := logger.Sugar()
	//defer sugar.Sync()
	//
	//sugar.Info("reading from ", zap.String("etcd", configRemoteEndpoint))
	//cli, err := clientv3.New(clientv3.Config{
	//	Endpoints:   []string{configRemoteEndpoint},
	//	DialTimeout: 5 * time.Second,
	//})
	//if err != nil {
	//	sugar.Warn(err)
	//	return
	//}
	//defer cli.Close()
	//
	//resp, err := cli.Get(context.Background(), configRemotePath)
	//if err != nil {
	//	sugar.Warn(err)
	//	return
	//}
	//
	//for _, kv := range resp.Kvs {
	//	if err := v.MergeConfig(bytes.NewBuffer(kv.Value)); err == nil {
	//		sugar.Debug("remote settings: ", v.AllSettings())
	//	} else {
	//		sugar.Warn(err)
	//	}
	//}
}

func showConfig(v *viper.Viper) {

}
