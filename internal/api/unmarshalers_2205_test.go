//go:build 2205

package api

import (
	"encoding/json"
	"testing"

	"github.com/Myrgued/prometheus-slurm-exporter/internal/util"
)

func TestUnmarshalDiagResponse(t *testing.T) {
	var r DiagResp
	fb := util.ReadTestDataBytes("slurm_v0_0_38_diag.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal diag response: %v\n", err)
	}
}

func TestUnmarshalJobsResponse(t *testing.T) {
	var r JobsResp
	fb := util.ReadTestDataBytes("slurm_v0_0_38_get_jobs.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal jobs response: %v\n", err)
	}
}

func TestUnmarshalNodesResponse(t *testing.T) {
	var r NodesResp
	fb := util.ReadTestDataBytes("slurm_v0_0_38_get_nodes.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal nodes response: %v\n", err)
	}
}

func TestUnmarshalPartitionsResponse(t *testing.T) {
	var r PartitionsResp
	fb := util.ReadTestDataBytes("slurm_v0_0_38_get_partitions.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal partition response: %v\n", err)
	}
}
