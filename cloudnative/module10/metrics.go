package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

//CreateLatencyHistogram工厂方法返回一个*prometheus.HistogramVec实例，bucket的设置在函数体内
func CreateLatencyHistogram(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:      "latency_seconds",
			Namespace: namespace,
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 14),
		}, []string{"step"},
	)
}

//定义一个对象LatencyMetrics，内部包含一个直方图对象和时间值
type LatencyMetrics struct {
	histo *prometheus.HistogramVec
	start time.Time
}

//实现一个方法ObserveTotal，封装直方图的接口方法Observe(float64)，将时间差作为参数传入，并且添加标签内容到直方图对象
func (t *LatencyMetrics) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

//创建一个LatencyMetrics的工厂方法，参数是一个直方图对象
func NewLatencyMetrics(histo *prometheus.HistogramVec) *LatencyMetrics {
	return &LatencyMetrics{
		histo: histo,
		start: time.Now(),
	}
}

//定义一个常量，描述指标所在namespace
const MetricsNamespace = "httpserver"

//创建一个直方图对象实例
var latencyHistogram = CreateLatencyHistogram(MetricsNamespace, "Time to spend")

//创建一个方法Register，封装prometheus的注册方法，并将上述的直方图对象实例传入
func Register() {
	err := prometheus.Register(latencyHistogram)
	if err != nil {
		fmt.Println(err)
	}
}

//创建一个方法NewLatency，封装以注册中所用直方图对象实例作为参数生成的LatencyMetrics实例
func NewLatency() *LatencyMetrics {
	return NewLatencyMetrics(latencyHistogram)
}
