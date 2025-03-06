package api

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Myrgued/prometheus-slurm-exporter/internal/types"
)

// Instead of using the Openapi-generated data in our slurm package, we
// take the unmarshaled response data and load everything we care about
// into version-agnostic models so our implementation doesn't have to depend
// on version. We do still store the version tag in the data struct so
// if there's any cases where we need a conditional in the code, we have
// that information stored.

// Each of these models were copied and modified from the openapi
// models for simplicity. We can add to these models as we want to extract more
// information

type DiagData struct {
	ApiVersion             string
	ServerThreadCount      int32
	AgentQueueSize         int32
	DbdAgentQueueSize      int32
	ScheduleCycleLast      int32
	ScheduleCycleMean      int64
	ScheduleCyclePerMinute int64
	BfDepthMean            int64
	BfCycleLast            int32
	BfCycleMean            int64
	BfBackfilledJobs       int32
	BfLastBackfilledJobs   int32
	BfBackfilledHetJobs    int32
}

func NewDiagData() *DiagData {
	return &DiagData{
		ApiVersion: apiVersion,
	}
}

func (d *DiagData) SetServerThreadCount(v *int32) error {
	if v == nil {
		return fmt.Errorf("server thread count not found in diag data")
	}
	d.ServerThreadCount = *v
	return nil
}

func (d *DiagData) SetAgentQueueSize(v *int32) error {
	if v == nil {
		return fmt.Errorf("agent queue size not found in diag data")
	}
	d.AgentQueueSize = *v
	return nil
}

func (d *DiagData) SetDbdAgentQueueSize(v *int32) error {
	if v == nil {
		return fmt.Errorf("dbd agent queue size not found in diag data")
	}
	d.DbdAgentQueueSize = *v
	return nil
}

func (d *DiagData) SetScheduleCycleLast(v *int32) error {
	if v == nil {
		return fmt.Errorf("schedule cycle last not found in diag data")
	}
	d.ScheduleCycleLast = *v
	return nil
}

func (d *DiagData) SetScheduleCycleMean(v *int64) error {
	if v == nil {
		return fmt.Errorf("schedule cycle mean not found in diag data")
	}
	d.ScheduleCycleMean = *v
	return nil
}

func (d *DiagData) SetScheduleCyclePerMinute(v *int64) error {
	if v == nil {
		return fmt.Errorf("schedule cycle per minute not found in diag data")
	}
	d.ScheduleCyclePerMinute = *v
	return nil
}

func (d *DiagData) SetBfDepthMean(v *int64) error {
	if v == nil {
		return fmt.Errorf("backfill depth mean not found in diag data")
	}
	d.BfDepthMean = *v
	return nil
}

func (d *DiagData) SetBfCycleLast(v *int32) error {
	if v == nil {
		return fmt.Errorf("backfill cycle last not found in diag data")
	}
	d.BfCycleLast = *v
	return nil
}

func (d *DiagData) SetBfCycleMean(v *int64) error {
	if v == nil {
		return fmt.Errorf("backfill cycle mean not found in diag data")
	}
	d.BfCycleMean = *v
	return nil
}

func (d *DiagData) SetBfBackfilledJobs(v *int32) error {
	if v == nil {
		return fmt.Errorf("backfilled jobs not found in diag data")
	}
	d.BfBackfilledJobs = *v
	return nil
}

// TODO: This is probably not correct, should revisit this number
func (d *DiagData) SetBfLastBackfilledJobs(v *int32) error {
	if v == nil {
		return fmt.Errorf("last backfilled jobs not found in diag data")
	}
	d.BfLastBackfilledJobs = *v
	return nil
}

func (d *DiagData) SetBfBackfilledHetJobs(v *int32) error {
	if v == nil {
		return fmt.Errorf("backfilled heterogeneous jobs not found in diag data")
	}
	d.BfBackfilledHetJobs = *v
	return nil
}

func (d *DiagData) FromResponse(r DiagResp) error {
	var err error
	if err = d.SetServerThreadCount(r.Statistics.ServerThreadCount); err != nil {
		return err
	}
	if err = d.SetAgentQueueSize(r.Statistics.AgentQueueSize); err != nil {
		return err
	}
	if err = d.SetDbdAgentQueueSize(r.Statistics.DbdAgentQueueSize); err != nil {
		return err
	}
	if err = d.SetScheduleCycleLast(r.Statistics.ScheduleCycleLast); err != nil {
		return err
	}
	if err = d.SetScheduleCycleMean(r.Statistics.ScheduleCycleMean); err != nil {
		return err
	}
	if err = d.SetScheduleCyclePerMinute(r.Statistics.ScheduleCyclePerMinute); err != nil {
		return err
	}
	if err = d.SetBfDepthMean(r.Statistics.BfDepthMean); err != nil {
		return err
	}
	if err = d.SetBfCycleLast(r.Statistics.BfCycleLast); err != nil {
		return err
	}
	if err = d.SetBfCycleMean(r.Statistics.BfCycleMean); err != nil {
		return err
	}
	if err = d.SetBfLastBackfilledJobs(r.Statistics.BfLastBackfilledJobs); err != nil {
		return err
	}
	if err = d.SetBfBackfilledJobs(r.Statistics.BfBackfilledJobs); err != nil {
		return err
	}
	if err = d.SetBfBackfilledHetJobs(r.Statistics.BfBackfilledHetJobs); err != nil {
		return err
	}
	return nil
}

type NodesData struct {
	ApiVersion string
	Nodes      []NodeData
}

type NodeData struct {
	Name          string
	Hostname      string
	States        []types.NodeState
	Tres          string
	TresUsed      string
	Partitions    []string
	AllocMemory   int64
	RealMemory    int64
	AllocCpus     int32
	AllocIdleCpus int32
	OtherCpus     int32
	Cpus          int32
	GPUTotal      int32
	GPUAllocated  int32
}

func NewNodesData() *NodesData {
	return &NodesData{
		ApiVersion: apiVersion,
	}
}

func (n *NodeData) SetName(name *string) error {
	if name == nil {
		return fmt.Errorf("node name not found in node information")
	}
	n.Name = *name
	return nil
}

func (n *NodeData) SetHostname(name *string) error {
	if name == nil {
		return fmt.Errorf("node hostname not found in node information")
	}
	n.Hostname = *name
	return nil
}

func (n *NodeData) SetPartitions(partitions []string) error {
	n.Partitions = partitions
	return nil
}

func (n *NodeData) SetTres(tres *string) {
	if tres != nil {
		n.Tres = *tres
	}
}

func (n *NodeData) SetTresUsed(tresUsed *string) {
	if tresUsed != nil {
		n.TresUsed = *tresUsed
	}
}

func (n *NodeData) SetAllocMemory(allocMemory *int64) error {
	if allocMemory == nil {
		n.AllocMemory = 0
		return nil
	}
	n.AllocMemory = *allocMemory
	return nil
}

func (n *NodeData) SetTotalMemory(totalMemory *int64) error {
	if totalMemory == nil {
		n.RealMemory = 0
		return nil
	}
	n.RealMemory = *totalMemory
	return nil
}

func (n *NodeData) SetAllocCPUs(allocCPUs *int32) error {
	if allocCPUs == nil {
		n.AllocCpus = 0
		return nil
	}
	n.AllocCpus = *allocCPUs
	return nil
}

func (n *NodeData) SetIdleCPUs(idleCPUs *int32) error {
	if idleCPUs == nil {
		n.AllocIdleCpus = 0
		return nil
	}
	n.AllocIdleCpus = *idleCPUs
	return nil
}

// This isn't in the api so it's always 0
func (n *NodeData) SetOtherCPUs() error {
	n.OtherCpus = 0
	return nil
}

func (n *NodeData) SetTotalCPUs(totalCPUs *int32) error {
	if totalCPUs == nil {
		n.Cpus = 0
		return nil
	}
	n.Cpus = *totalCPUs
	return nil
}

func (n *NodeData) SetNodeGPUTotal(tresString *string) error {
	parts := strings.Split(*tresString, ",")
	for _, p := range parts {
		if strings.Contains(p, "gres/gpu=") {
			gp := strings.Split(p, "=")
			if len(gp) != 2 {
				return fmt.Errorf("found gpu in tres but failed to parse: %s", p)
			}
			ns := gp[1]
			ng, err := strconv.Atoi(ns)
			if err != nil {
				return fmt.Errorf("failed to parse number of gpus from tres: %s", p)
			}
			n.GPUTotal = int32(ng)
			return nil
		}
	}
	n.GPUTotal = 0
	return nil
}

func (n *NodeData) SetNodeGPUAllocated(tresString *string) error {
	parts := strings.Split(*tresString, ",")
	for _, p := range parts {
		if strings.Contains(p, "gres/gpu=") {
			gp := strings.Split(p, "=")
			if len(gp) != 2 {
				return fmt.Errorf("found gpu in tres but failed to parse: %s", p)
			}
			ns := gp[1]
			ng, err := strconv.Atoi(ns)
			if err != nil {
				return fmt.Errorf("failed to parse number of gpus from tres: %s", p)
			}
			n.GPUAllocated = int32(ng)
			return nil
		}
	}
	n.GPUAllocated = 0
	return nil
}

// Modified for Slurm 22.05
func (n *NodeData) SetNodeStates(states *string) error {
    if states == nil {
        // No state found in the node response
        return fmt.Errorf("node state not found in node")
    }

    strStates := strings.Split(*states, ",")
    var nodeStates []types.NodeState

    patterns := map[*regexp.Regexp]types.NodeState{
        regexp.MustCompile(`^alloc`):          types.NodeStateAlloc,
        regexp.MustCompile(`^comp`):           types.NodeStateComp,
        regexp.MustCompile(`^down`):           types.NodeStateDown,
        regexp.MustCompile(`^drain`):          types.NodeStateDrain,
        regexp.MustCompile(`^fail`):           types.NodeStateFail,
        regexp.MustCompile(`^err`):            types.NodeStateErr,
        regexp.MustCompile(`^idle`):           types.NodeStateIdle,
        regexp.MustCompile(`^maint`):          types.NodeStateMaint,
        regexp.MustCompile(`^mix`):            types.NodeStateMix,
        regexp.MustCompile(`^planned`):        types.NodeStatePlanned,
        regexp.MustCompile(`^res`):            types.NodeStateResv,
        regexp.MustCompile(`^not_responding`): types.NodeStateNotResponding,
        regexp.MustCompile(`^invalid`):        types.NodeStateInvalid,
        regexp.MustCompile(`^invalid_reg`):    types.NodeStateInvalidReg,
        regexp.MustCompile(`^dynamic_norm`):   types.NodeStateDynamicNorm,
    }

    for _, s := range strStates {
        state := strings.ToLower(strings.TrimSpace(s))
        matched := false

        for re, stateUnit := range patterns {
            if re.MatchString(state) {
                nodeStates = append(nodeStates, stateUnit)
                matched = true
                break
            }
        }

        if !matched {
            return fmt.Errorf("failed to match cpu state against known states: %v", state)
        }
    }

    n.States = nodeStates
    return nil
}

func (n *NodeData) GetNodeStatesString(delim string) (string, error) {
	strStates := make([]string, len(n.States))
	for i, s := range n.States {
		strStates[i] = string(s)
	}
	return strings.Join(strStates, delim), nil
}

func (d *NodesData) FromResponse(r NodesResp) error {
	var err error
	for _, n := range r.Nodes {
		nd := NodeData{}
		if err = nd.SetName(n.Name); err != nil {
			return err
		}
		if err = nd.SetHostname(n.Hostname); err != nil {
			return err
		}
		if err = nd.SetNodeStates(n.State); err != nil {
			return err
		}
		if err = nd.SetPartitions(n.Partitions); err != nil {
			return err
		}
		nd.SetTres(n.Tres)
		nd.SetTresUsed(n.TresUsed)
		if err = nd.SetTotalCPUs(n.Cpus); err != nil {
			return err
		}
		if err = nd.SetAllocCPUs(n.AllocCpus); err != nil {
			return err
		}
		if err = nd.SetIdleCPUs(n.AllocIdleCpus); err != nil {
			return err
		}
		if err = nd.SetOtherCPUs(); err != nil {
			return err
		}

		if err = nd.SetTotalMemory(n.RealMemory); err != nil {
			return err
		}
		if err = nd.SetAllocMemory(n.AllocMemory); err != nil {
			return err
		}

		if err = nd.SetNodeGPUAllocated(n.TresUsed); err != nil {
			return err
		}
		if err = nd.SetNodeGPUTotal(n.Tres); err != nil {
			return err
		}

		d.Nodes = append(d.Nodes, nd)
	}

	return nil
}

type JobsData struct {
	ApiVersion string
	Jobs       []JobData
}

type JobData struct {
	Account    string
	UserName   string
	JobState   types.JobState
	Cpus       int32
	Partition  string
	Dependency string
}

func NewJobsData() *JobsData {
	return &JobsData{
		ApiVersion: apiVersion,
	}
}

func (j *JobData) SetJobAccount(name *string) error {
	if name == nil {
		return fmt.Errorf("failed to find account name in job")
	}
	j.Account = *name
	return nil
}

func (j *JobData) SetJobUserName(name *string) error {
	if name == nil {
		return fmt.Errorf("failed to find username in job")
	}
	j.UserName = *name
	return nil
}

func (j *JobData) SetJobPartitionName(name *string) error {
	if name == nil {
		return fmt.Errorf("failed to find partition name in job")
	}
	j.Partition = *name
	return nil
}

func (j *JobData) SetJobDependency(dependency *string) error {
	if dependency == nil {
		j.Dependency = ""
	}
	j.Dependency = *dependency
	return nil
}


// Modified for Slurm 22.05
func (j *JobData) SetJobState(states *string) error {
    if states == nil {
        return fmt.Errorf("job state not found in job")
    }

    state := strings.ToLower(strings.TrimSpace(*states))

    patterns := map[*regexp.Regexp]types.JobState{
        regexp.MustCompile(`^completed`):      types.JobStateCompleted,
        regexp.MustCompile(`^pending`):        types.JobStatePending,
        regexp.MustCompile(`^failed`):         types.JobStateFailed,
        regexp.MustCompile(`^running`):        types.JobStateRunning,
        regexp.MustCompile(`^suspended`):      types.JobStateSuspended,
        regexp.MustCompile(`^out_of_memory`):  types.JobStateOutOfMemory,
        regexp.MustCompile(`^timeout`):        types.JobStateTimeout,
        regexp.MustCompile(`^cancelled`):      types.JobStateCancelled,
        regexp.MustCompile(`^completing`):     types.JobStateCompleting,
        regexp.MustCompile(`^configuring`):    types.JobStateConfiguring,
        regexp.MustCompile(`^node_fail`):      types.JobStateNodeFail,
        regexp.MustCompile(`^preempted`):      types.JobStatePreempted,
    }

    for re, stateUnit := range patterns {
        if re.MatchString(state) {
            j.JobState = stateUnit
            return nil
        }
    }

    return fmt.Errorf("failed to match job state against known states: %v", state)
}


func (j *JobData) SetJobCPUs(jobcpus *int32) error {
	if jobcpus == nil {
		j.Cpus = 0
		return nil
	}
	j.Cpus = *jobcpus
	return nil
}

func (d *JobsData) FromResponse(r JobsResp) error {
	var err error
	for _, j := range r.Jobs {
		jd := JobData{}
		if err = jd.SetJobAccount(j.Account); err != nil {
			return err
		}
		if err = jd.SetJobUserName(j.UserName); err != nil {
			return err
		}
		if err = jd.SetJobPartitionName(j.Partition); err != nil {
			return err
		}
		if err = jd.SetJobState(j.JobState); err != nil {
			return err
		}
		if err = jd.SetJobDependency(j.Dependency); err != nil {
			return err
		}
		if err = jd.SetJobCPUs(j.JobResources.Cpus); err != nil {
			return err
		}
		d.Jobs = append(d.Jobs, jd)
	}

	return nil
}

type PartitionsData struct {
	ApiVersion string
	Partitions []PartitionData
}

type PartitionData struct {
	Name      string
	Cpus      int32
	OtherCpus int32
	Nodes     string
}

func NewPartitionsData() *PartitionsData {
	return &PartitionsData{
		ApiVersion: apiVersion,
	}
}

func (p *PartitionData) SetName(name *string) error {
	if name == nil {
		return fmt.Errorf("failed to find name in partition")
	}
	p.Name = *name
	return nil
}

func (p *PartitionData) SetTotalCPUs(totalCPUs *int32) error {
	if totalCPUs == nil {
		return fmt.Errorf("failed to find total cpus in partition")
	}
	p.Cpus = *totalCPUs
	return nil
}

// this isnt in the api so its always 0
func (p *PartitionData) SetOtherCPUs() error {
	p.OtherCpus = 0
	return nil
}

func (p *PartitionData) SetNodeList(configuredNodes *string) error {
	if configuredNodes == nil {
		return fmt.Errorf("failed to find node list for partition")
	}
	p.Nodes = *configuredNodes
	return nil
}

func (d *PartitionsData) FromResponse(r PartitionsResp) error {
	var err error
	for _, p := range r.Partitions {
		pd := PartitionData{}
		if err = pd.SetName(p.Name); err != nil {
			return err
		}
		if err = pd.SetTotalCPUs(p.Cpus); err != nil {
			return err
		}
		if err = pd.SetOtherCPUs(); err != nil {
			return err
		}
		if err = pd.SetNodeList(p.Nodes); err != nil {
			return err
		}
		d.Partitions = append(d.Partitions, pd)
	}

	return nil
}

type SharesData struct {
	ApiVersion string
	Shares     []ShareData
}

type ShareData struct {
	Name           string
	EffectiveUsage float64
}

func NewSharesData() *SharesData {
	return &SharesData{
		ApiVersion: apiVersion,
	}
}

func (s *ShareData) SetName(name *string) error {
	if name == nil {
		return fmt.Errorf("failed to find name in fair share")
	}
	s.Name = *name
	return nil
}

func (s *ShareData) SetEffectiveUsage(effectiveUsage *float64) error {
	if effectiveUsage == nil {
		s.EffectiveUsage = float64(0)
	}
	s.EffectiveUsage = *effectiveUsage
	return nil
}

func (d *SharesData) FromResponse(r SharesResp) error {
	var err error
	for _, s := range r.Shares.Shares {
		sd := ShareData{}
		if err = sd.SetName(s.Name); err != nil {
			return err
		}
		if err = sd.SetEffectiveUsage(s.EffectiveUsage); err != nil {
			return err
		}

		d.Shares = append(d.Shares, sd)
	}

	return nil
}

// This is used for unmarshaling errors on 500 status codes
type APIErrorData struct {
	Errors []struct {
		Description string `json:"description"`
		ErrorNumber int    `json:"error_number"`
		Error       string `json:"error"`
		Source      string `json:"source"`
	} `json:"errors"`
}

func (aed APIErrorData) ToString() string {
	var errStrings []string
	for _, e := range aed.Errors {
		errStrings = append(errStrings, fmt.Sprintf("{description=%s, error_number=%d, error=%s, source=%s}", e.Description, e.ErrorNumber, e.Error, e.Source))
	}
	return fmt.Sprintf("slurm api errors [%s]", strings.Join(errStrings, ","))
}
