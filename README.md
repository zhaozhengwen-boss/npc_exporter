# NPU Exporter

NPU Exporter 是一个用于采集华为 Ascend NPU (Neural Processing Unit) 设备监控指标的 Prometheus Exporter。它通过解析 `npu-smi info` 命令的输出，将 NPU 的关键性能指标（如利用率、显存、温度等）转换为 Prometheus 格式的 Metrics，方便用户在 Grafana 等工具中进行可视化监控。

## 功能特性

*   **轻量级**: 使用 Go 语言编写，单一二进制文件部署。
*   **简单易用**: 开箱即用，默认监听 9102 端口。
*   **多维度指标**: 提供设备级和芯片级的详细监控数据。
*   **兼容性**: 标准 Prometheus Metrics 格式。

## 前置条件

*   **操作系统**: Linux (支持华为 NPU 驱动的系统)。
*   **驱动工具**: 运行环境必须安装华为 NPU 驱动，且 `npu-smi` 命令可直接执行。
    *   验证方法：在终端运行 `npu-smi info`，应能看到 NPU 设备信息表。

## 构建与安装

### 从源码构建

确保你已经安装了 Go (1.24+)。

1.  克隆仓库：
    ```bash
    git clone https://github.com/zhaozhengwen-boss/npc_exporter.git
    cd npu_exporter
    ```

2.  下载依赖并构建：
    ```bash
    go mod tidy
    go build -o npu_exporter main.go
    ```

3.  构建完成后，当前目录下会生成 `npu_exporter` 可执行文件。

## 运行

直接运行编译好的二进制文件：

```bash
./npu_exporter
```

程序默认监听 **9102** 端口。

启动成功后，访问 `http://localhost:9102/metrics` 即可看到采集到的指标数据。

## 指标详解 (Metrics)

Exporter 提供以下核心指标，所有指标（除总数外）都包含 `device_id` 和 `chip_id` 标签，以便区分不同的 NPU 设备。

| 指标名称 (Metric Name) | 类型 (Type) | 描述 (Description) |
| :--- | :--- | :--- |
| `npu_device_count` | Gauge | NPU 设备总数 |
| `npu_aicore_utilization_percent` | Gauge | AI Core 计算单元使用率 (%) |
| `npu_memory_used_mb` | Gauge | 已使用的显存大小 (MB) |
| `npu_memory_total_mb` | Gauge | 总可用显存大小 (MB) |
| `npu_temperature_celsius` | Gauge | 设备温度 (°C) |

## Prometheus 配置示例

在你的 `prometheus.yml` 配置文件中添加以下 job：

```yaml
scrape_configs:
  - job_name: 'npu_exporter'
    static_configs:
      - targets: ['localhost:9102']
```

## Grafana 面板

你可以使用上述指标在 Grafana 中创建仪表盘。例如：
*   **AI Core 利用率**: `npu_aicore_utilization_percent`
*   **显存使用情况**: `npu_memory_used_mb / npu_memory_total_mb * 100`

## 目录结构

*   internal/`: 内部逻辑包。
    *   `collector/`: Prometheus 采集器实现。
    *   `npu/`: `npu-smi` 命令执行与输出解析逻辑。
    *   `logger/`: 日志工具。

## 常见问题

**Q: 运行报错 "failed to run npu-smi info"？**
A: 请确保当前用户有权限执行 `npu-smi` 命令，且该命令在系统的 PATH 路径下。通常需要 root 权限或将用户加入特定组。

## 许可证

MIT License
