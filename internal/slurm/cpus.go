/* Copyright 2017-2022 Lucas Crownover, Victor Penso, Matteo Dessalvi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package slurm

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lcrownover/prometheus-slurm-exporter/internal/cache"
	"github.com/prometheus/client_golang/prometheus"
)

type cpusMetrics struct {
	alloc float64
	idle  float64
	other float64
	total float64
}

func NewCPUsMetrics() *cpusMetrics {
	return &cpusMetrics{}
}

func (cm *cpusMetrics) AddAlloc(n float64) {
	cm.alloc += n
}

// ParseCPUMetrics pulls out total cluster cpu states of alloc,idle,other,total
func ParseCPUsMetrics(ctx context.Context) (*cpusMetrics, error) {
	cm := NewCPUsMetrics()

	jc := cache.GetResponseCache(ctx).JobsCache()
	jobs, err := jc.Jobs()
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs for cpu metrics: %v", err)
	}
	for _, j := range *jobs {
		state, err := GetJobState(j)
		if err != nil {
			slog.Error("failed to get job state", "error", err)
			continue
		}
		cpus, err := GetJobCPUs(j)
		if err != nil {
			slog.Error("failed to get job cpus", "error", err)
			continue
		}
		if *state == JobStateRunning {
			cm.AddAlloc(*cpus)
		}
	}
	return cm, nil
}

/*
 * Implement the Prometheus Collector interface and feed the
 * Slurm scheduler metrics into it.
 * https://godoc.org/github.com/prometheus/client_golang/prometheus#Collector
 */
func NewCPUsCollector(ctx context.Context) *CPUsCollector {
	return &CPUsCollector{
		ctx:   ctx,
		alloc: prometheus.NewDesc("slurm_cpus_alloc", "Allocated CPUs", nil, nil),
		idle:  prometheus.NewDesc("slurm_cpus_idle", "Idle CPUs", nil, nil),
		other: prometheus.NewDesc("slurm_cpus_other", "Mix CPUs", nil, nil),
		total: prometheus.NewDesc("slurm_cpus_total", "Total CPUs", nil, nil),
	}
}

type CPUsCollector struct {
	ctx   context.Context
	alloc *prometheus.Desc
	idle  *prometheus.Desc
	other *prometheus.Desc
	total *prometheus.Desc
}

// Send all metric descriptions
func (cc *CPUsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- cc.alloc
	ch <- cc.idle
	ch <- cc.other
	ch <- cc.total
}

func (cc *CPUsCollector) Collect(ch chan<- prometheus.Metric) {
	cm, err := ParseCPUsMetrics(cc.ctx)
	if err != nil {
		slog.Error("failed to collect cpus metrics", "error", err)
		return
	}
	ch <- prometheus.MustNewConstMetric(cc.alloc, prometheus.GaugeValue, cm.alloc)
	ch <- prometheus.MustNewConstMetric(cc.idle, prometheus.GaugeValue, cm.idle)
	ch <- prometheus.MustNewConstMetric(cc.other, prometheus.GaugeValue, cm.other)
	ch <- prometheus.MustNewConstMetric(cc.total, prometheus.GaugeValue, cm.total)
}

//
//
// all the old stuff below
//
//

// Execute the sinfo command and return its output
func CPUsDataOld() []byte {
	cmd := exec.Command("sinfo", "-h", "-o %C")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	out, _ := io.ReadAll(stdout)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return out
}

func CPUsGetMetricsOld() *cpusMetrics {
	return ParseCPUsMetricsOld(CPUsDataOld())
}

func ParseCPUsMetricsOld(input []byte) *cpusMetrics {
	var cm cpusMetrics
	if strings.Contains(string(input), "/") {
		splitted := strings.Split(strings.TrimSpace(string(input)), "/")
		cm.alloc, _ = strconv.ParseFloat(splitted[0], 64)
		cm.idle, _ = strconv.ParseFloat(splitted[1], 64)
		cm.other, _ = strconv.ParseFloat(splitted[2], 64)
		cm.total, _ = strconv.ParseFloat(splitted[3], 64)
	}
	return &cm
}

func NewCPUsCollectorOld() *CPUsCollectorOld {
	return &CPUsCollectorOld{
		alloc: prometheus.NewDesc("slurm_old_cpus_alloc", "Allocated CPUs", nil, nil),
		idle:  prometheus.NewDesc("slurm_old_cpus_idle", "Idle CPUs", nil, nil),
		other: prometheus.NewDesc("slurm_old_cpus_other", "Mix CPUs", nil, nil),
		total: prometheus.NewDesc("slurm_old_cpus_total", "Total CPUs", nil, nil),
	}
}

type CPUsCollectorOld struct {
	alloc *prometheus.Desc
	idle  *prometheus.Desc
	other *prometheus.Desc
	total *prometheus.Desc
}

// Send all metric descriptions
func (cc *CPUsCollectorOld) Describe(ch chan<- *prometheus.Desc) {
	ch <- cc.alloc
	ch <- cc.idle
	ch <- cc.other
	ch <- cc.total
}
func (cc *CPUsCollectorOld) Collect(ch chan<- prometheus.Metric) {
	cm := CPUsGetMetricsOld()
	ch <- prometheus.MustNewConstMetric(cc.alloc, prometheus.GaugeValue, cm.alloc)
	ch <- prometheus.MustNewConstMetric(cc.idle, prometheus.GaugeValue, cm.idle)
	ch <- prometheus.MustNewConstMetric(cc.other, prometheus.GaugeValue, cm.other)
	ch <- prometheus.MustNewConstMetric(cc.total, prometheus.GaugeValue, cm.total)
}
