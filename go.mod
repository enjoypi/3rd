module github.com/enjoypi/3rd

go 1.14

//replace github.com/coreos/etcd => go.etcd.io/etcd v3.3.22+incompatible

//replace go.etcd.io/etcd => github.com/coreos/etcd v3.3.22+incompatible

replace github.com/enjoypi/god => ../god

require (
	github.com/enjoypi/god v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	go.uber.org/zap v1.15.0
	gopkg.in/yaml.v2 v2.3.0
)
