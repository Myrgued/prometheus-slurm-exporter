//go:build 2205

package api

import (
	"github.com/Myrgued/prometheus-slurm-exporter/internal/types"
)

var versionedEndpoints = []endpoint{
	{types.ApiJobsEndpointKey, "jobs", "/slurm/v0.0.38/jobs"},
	{types.ApiNodesEndpointKey, "nodes", "/slurm/v0.0.38/nodes"},
	{types.ApiPartitionsEndpointKey, "partitions", "/slurm/v0.0.38/partitions"},
	{types.ApiDiagEndpointKey, "diag", "/slurm/v0.0.38/diag"},
	{types.ApiSharesEndpointKey, "shares", "/slurm/v0.0.40/shares"},
}
