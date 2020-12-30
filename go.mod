module github.com/enjoypi/3rd

go 1.15

replace github.com/enjoypi/god => ../god

replace github.com/enjoypi/gostatechart => ../gostatechart

require (
	github.com/enjoypi/god v0.0.0-00010101000000-000000000000
	github.com/enjoypi/gostatechart v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.4.3
	github.com/nats-io/nats.go v1.10.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.16.0
	gopkg.in/yaml.v2 v2.4.0
)
