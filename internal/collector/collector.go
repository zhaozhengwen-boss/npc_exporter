// internal/collector/collector.go
package collector

import (
	"npu_exporter/internal/logger"
	"npu_exporter/internal/npu"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type NPUCollector struct {
}

// 简化NewNPUCollector函数
func NewNPUCollector() *NPUCollector {
	return &NPUCollector{}
}

func (c *NPUCollector) Describe(ch chan<- *prometheus.Desc) {
	// NPU指标
	ch <- deviceCountDesc
	ch <- aicoreUtilDesc
	ch <- memoryUsedDesc
	ch <- memoryTotalDesc
	ch <- temperatureDesc
}

func (c *NPUCollector) Collect(ch chan<- prometheus.Metric) {
	// NPU指标收集逻辑
	stats, err := npu.CollectAll()
	if err != nil {
		logger.Logger.Error("Failed to collect NPU stats", "error", err)
	} else {
		ch <- prometheus.MustNewConstMetric(deviceCountDesc, prometheus.GaugeValue, float64(len(stats)))
		for _, s := range stats {
			labels := []string{strconv.Itoa(s.DeviceID), strconv.Itoa(s.ChipID)}
			ch <- prometheus.MustNewConstMetric(aicoreUtilDesc, prometheus.GaugeValue, s.Utilization, labels...)
			ch <- prometheus.MustNewConstMetric(memoryUsedDesc, prometheus.GaugeValue, float64(s.MemoryUsed), labels...)
			ch <- prometheus.MustNewConstMetric(memoryTotalDesc, prometheus.GaugeValue, float64(s.MemoryTotal), labels...)
			ch <- prometheus.MustNewConstMetric(temperatureDesc, prometheus.GaugeValue, float64(s.Temperature), labels...)
		}
	}
}
