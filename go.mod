module npu_exporter

go 1.24.0

toolchain go1.24.4

require (
	github.com/prometheus/client_golang v1.23.2
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.5 // indirect
	github.com/prometheus/procfs v0.19.2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	golang.org/x/sys v0.40.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

// 所有 indirect 依赖可由 go mod tidy 自动生成，无需手动维护
// 只需保留 replace 指令

replace github.com/docker/docker => github.com/moby/moby v28.5.2+incompatible
