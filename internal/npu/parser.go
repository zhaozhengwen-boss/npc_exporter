package npu

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Stats 表示一个 NPU 设备的监控数据
type Stats struct {
	DeviceID    int
	ChipID      int
	Utilization float64
	MemoryUsed  int
	MemoryTotal int
	Temperature int
}

// CollectAll 执行 npu-smi info 并返回所有设备数据
func CollectAll() ([]Stats, error) {
	cmd := exec.Command("npu-smi", "info")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run npu-smi info: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	var stats []Stats

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if !strings.HasPrefix(line, "|") || len(line) < 50 {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		col1 := strings.TrimSpace(parts[1]) // Chip Device
		col2 := strings.TrimSpace(parts[2]) // Bus-Id
		col3 := strings.TrimSpace(parts[3]) // AICore + Mem

		if !strings.Contains(col2, ":") || !strings.Contains(col2, ".") {
			continue
		}

		fields1 := strings.Fields(col1)
		if len(fields1) < 2 {
			continue
		}
		chipID, _ := strconv.Atoi(fields1[0])
		deviceID, _ := strconv.Atoi(fields1[1])

		fields3 := strings.Fields(col3)
		if len(fields3) < 4 || fields3[2] != "/" {
			continue
		}

		util, _ := strconv.ParseFloat(fields3[0], 64)
		memUsed, _ := strconv.Atoi(fields3[1])
		memTotal, _ := strconv.Atoi(fields3[3])

		temp := 0
		if i > 0 {
			prevLine := strings.TrimSpace(lines[i-1])
			if strings.HasPrefix(prevLine, "|") {
				prevParts := strings.Split(prevLine, "|")
				if len(prevParts) >= 4 {
					prevCol3 := strings.TrimSpace(prevParts[3])
					prevFields3 := strings.Fields(prevCol3)
					if len(prevFields3) > 1 {
						temp, _ = strconv.Atoi(prevFields3[1])
					}
				}
			}
		}

		stats = append(stats, Stats{
			DeviceID:    deviceID,
			ChipID:      chipID,
			Utilization: util,
			MemoryUsed:  memUsed,
			MemoryTotal: memTotal,
			Temperature: temp,
		})
	}

	return stats, nil
}
