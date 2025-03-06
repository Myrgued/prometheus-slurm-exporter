//go:build 2311

package api

import (
	"encoding/json"
	"testing"

	"github.com/Myrgued/prometheus-slurm-exporter/internal/util"
)

func TestUnmarshalDiagResponse(t *testing.T) {
	var r DiagResp
	fb := util.ReadTestDataBytes("V0040OpenapiDiagResp.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal diag response: %v\n", err)
	}
}

func TestUnmarshalJobsResponse(t *testing.T) {
	var r JobsResp
	fb := util.ReadTestDataBytes("V0040OpenapiJobInfoResp.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal jobs response: %v\n", err)
	}
}

func TestUnmarshalNodesResponse(t *testing.T) {
	var r NodesResp
	fb := util.ReadTestDataBytes("V0040OpenapiNodesResp.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal nodes response: %v\n", err)
	}
}

func TestUnmarshalPartitionsResponse(t *testing.T) {
	var r PartitionsResp
	fb := util.ReadTestDataBytes("V0040OpenapiPartitionResp.json")
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal partition response: %v\n", err)
	}
}

func TestUnmarshalSharesResponse(t *testing.T) {
	var r SharesResp
	fb := util.ReadTestDataBytes("V0040OpenapiSharesResp.json")
	fb = util.CleanseInfinity(fb)
	err := json.Unmarshal(fb, &r)
	if err != nil {
		t.Fatalf("failed to unmarshal shares response: %v\n", err)
	}
}
