package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// 自訂數值型態的測量數據。
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "CPU 目前的溫度。",
	})
	// 計數型態的測量數據，並帶有自訂標籤。
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "硬碟發生錯誤的次數。",
		},
		[]string{"device"},
	)
)

func init() {
	// 測量數據必須註冊才會暴露給外界知道：
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func SetCpuTemp() {
	for {
		cpuTemp.Set(30 + rand.Float64()*30)
		time.Sleep(time.Second * 15)
	}
}

func SetHDFailures() {
	for {
		time.Sleep(time.Second * 60)
		if rand.Float32() < 0.5 {
			hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
		}
	}
}

func main() {
	// 配置測量數據的數值。
	go SetCpuTemp()
	go SetHDFailures()

	// 我們會用 Prometheus 所提供的預設處理函式在 "/metrics" 路徑監控著。
	// 這會暴露我們的數據內容，所以 Prometheus 就能夠獲取這些數據。
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
