module github.com/uptimus/common

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.27.0

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/spf13/viper v1.7.1
	gitlab.wvservices.com/waves/gateways/gw-commons v0.0.521
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)
