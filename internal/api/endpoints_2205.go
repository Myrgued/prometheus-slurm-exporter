//go:build 2205

package api

import (
	"github.com/Myrgued/prometheus-slurm-exporter/internal/types"
)

var versionedEndpoints = []endpoint{
	{types.ApiJobsEndpointKey, "jobs", "/slurm/v0.0.36/jobs"},
	{types.ApiNodesEndpointKey, "nodes", "/slurm/v0.0.36/nodes"},
	{types.ApiPartitionsEndpointKey, "partitions", "/slurm/v0.0.36/partitions"},
	{types.ApiDiagEndpointKey, "diag", "/slurm/v0.0.36/diag"},
}
