package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doCmd represents the do command
var childCmd = &cobra.Command{
	Use:   "child",
	Short: "this is a child command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	PreRunE: preRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		return childRun(rootViper)
	},
}

func init() {
	rootCmd.AddCommand(childCmd)

	childCmd.Flags().String("child.str", "child string", "string flag for child")
	childCmd.Flags().BoolP("child.bool", "b", false, "bool flag for child")
}

func childRun(v *viper.Viper) error {
	var c cfgChild
	if err := v.Unmarshal(&c); err != nil {
		return err
	}
	logger.Sugar().Info("settings on child: %+v", c)
	return nil
}
