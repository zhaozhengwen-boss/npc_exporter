// cmd/npu-exporter/main.go
package main

import (
	"net/http"

	"npu_exporter/internal/collector"
	"npu_exporter/internal/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger.Init()
	log := logger.Logger

	// 创建NPU收集器
	prometheus.MustRegister(collector.NewNPUCollector())

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>NPU Exporter</title></head>
			<body>
				<h1>NPU Exporter</h1>
				<p><a href="/metrics">Metrics</a></p>
			</body>
		</html>`))
	})

	port := ":9102"
	log.Info("Starting exporter", "port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Error("Server failed", "error", err)
	}
}
