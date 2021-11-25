package common_types

// CloSpec defines the desired state of Clo
type CloSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	Clis []CloIndicatorMetric `json:"clis"`
	// +kubebuilder:validation:Required
	Gns string `json:"gns"`
	// +kubebuilder:validation:Required
	AloName string `json:"aloName"`
	// +kubebuilder:validation:Optional
	MetricsAvgCalculationTime int64 `json:"metricsAvgCalculationTime"`
}

// CloStatus defines the observed state of ClO
type CloStatus struct {
	Conditions         []CloCondition `json:"conditions"`
	Status             bool           `json:"status"`
	Reason             string         `json:"reason"`
	Message            string         `json:"message"`
	LastProbeTime      string         `json:"lastProbeTime"`
	LastTransitionTime string         `json:"lastTransitionTime"`
	TotalClis          int32          `json:"totalClis"`
	ScalingFlag        int32          `json:"scalingFlag"`
	// Capacity Data
	TotalNodes     int32 `json:"totalNodes"`
	CpuCapacity    int32 `json:"cpuCapacity"`
	MemoryCapacity int32 `json:"emoryCapacity"`
	// Burst Cluster Supporting GNSes
	BurstClusterSuportingGnses []string `json:"BurstClusterSuportingGnses"`
}

// IndicatorMetric spec to define the indicator metrics for a LO
type CloIndicatorMetric struct {
	MetricName CloIndicatorMetricName `json:"name"`
	UpperLimit int                    `json:"upper"`
	LowerLimit int                    `json:"lower"`
}

// CLO Indicator Metric Name Values
type CloIndicatorMetricName string

const (
	CPUUTILIZATION    CloIndicatorMetricName = "CPUUTILIZATION"
	MEMORYUTILIZATION CloIndicatorMetricName = "MEMORYUTILIZATION"
)

// CLO Condition struct for storing current condition of the CLO
type CloCondition struct {
	Type    CloConditionType   `json:"type"`
	Status  bool               `json:"status"`
	Reason  CloConditionReason `json:"reason"`
	Message string             `json:"message"`
}

// CLO Condition Type Values
type CloConditionType string

const (
	CliReady CloConditionType = "CliReady"
)

// CLO Condition reason
type CloConditionReason string

const (
	CliMetricsBreached CloConditionReason = "CliMetricsBreached"
	CliMetricsError    CloConditionReason = "CliMetricsError"
)
