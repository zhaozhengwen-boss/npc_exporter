package collector

import "github.com/prometheus/client_golang/prometheus"

var (
	deviceCountDesc = prometheus.NewDesc(
		"npu_device_count",
		"NPU 设备总数",
		nil,
		nil,
	)

	aicoreUtilDesc = prometheus.NewDesc(
		"npu_aicore_utilization_percent",
		"AI Core 计算单元使用率（百分比）",
		[]string{"device_id", "chip_id"},
		nil,
	)

	memoryUsedDesc = prometheus.NewDesc(
		"npu_memory_used_mb",
		"已使用的显存（MB）",
		[]string{"device_id", "chip_id"},
		nil,
	)

	memoryTotalDesc = prometheus.NewDesc(
		"npu_memory_total_mb",
		"总可用显存（MB）",
		[]string{"device_id", "chip_id"},
		nil,
	)

	temperatureDesc = prometheus.NewDesc(
		"npu_temperature_celsius",
		"设备温度（摄氏度）",
		[]string{"device_id", "chip_id"},
		nil,
	)
)
