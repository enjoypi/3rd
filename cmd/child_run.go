package cmd

import (
	"github.com/spf13/viper"
)

// config for sub
type cfgChild struct {
	Config config
	Child  child
}

type child struct {
	Bool bool
	Str  string
}

type config struct {
	File string
}

func childRun(v *viper.Viper) error {
	var c cfgChild
	if err := v.Unmarshal(&c); err != nil {
		return err
	}
	logger.Sugar().Info("settings on child: %+v", c)
	return nil
}
