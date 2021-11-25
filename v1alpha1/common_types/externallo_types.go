package common_types

// ExternalLoSpec defines the desired state of ExternalLo
type ExternalLoSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	Elis []EloIndicatorMetric `json:"elis"`
	// +kubebuilder:validation:Required
	SpanName string `json:"spanName"`
	// +kubebuilder:validation:Optional
	Type string `json:"type"`
	// +kubebuilder:validation:Optional
	EloLatencyAvgTimeSlotInSec int64 `json:"eloLatencyAvgTimeSlotInSec"`
}

// ExternalLoStatus defines the observed state of ExternalLo
type ExternalLoStatus struct {
	Conditions         []EloCondition `json:"conditions"`
	Status             bool           `json:"status"`
	Reason             string         `json:"reason"`
	Message            string         `json:"message"`
	LastProbeTime      string         `json:"lastProbeTime"`
	LastTransitionTime string         `json:"lastTransitionTime"`
	TotalElis          int32          `json:"totalElis"`
}

// IndicatorMetric spec to define the indicator metrics for a LO
type EloIndicatorMetric struct {
	MetricName EloIndicatorMetricName `json:"name"`
	UpperLimit int                    `json:"upper"`
	LowerLimit int                    `json:"lower"`
}

// ELO Indicator Metric Name Values
type EloIndicatorMetricName string

const (
	ELOP50LATENCY EloIndicatorMetricName = "P50LATENCY"
	ELOP90LATENCY EloIndicatorMetricName = "P90LATENCY"
	ELOP99LATENCY EloIndicatorMetricName = "P99LATENCY"
	ELOERRORRATE  EloIndicatorMetricName = "ERRORRATE"
)

// ELO Condition struct for storing current condition of the TLO
type EloCondition struct {
	Type    EloConditionType   `json:"type"`
	Status  bool               `json:"status"`
	Reason  EloConditionReason `json:"reason"`
	Message string             `json:"message"`
}

// ELO Condition Type Values
type EloConditionType string

const (
// EliReady EloConditionType = "EliReady"
)

// ELO Condition reason
type EloConditionReason string

const (
	EliLatencyBreached EloConditionReason = "LatencyBreached"
	EliMetricsError    EloConditionReason = "MetricsError"
)
