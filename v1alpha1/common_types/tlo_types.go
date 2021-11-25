package common_types

// TloSpec defines the desired state of Tlo
type TloSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	Tlis []TloIndicatorMetric `json:"tlis"`
	// +kubebuilder:validation:Required
	TransactionType string `json:"transactionType"`
	// +kubebuilder:validation:Required
	ServiceName string `json:"serviceName"`
	// +kubebuilder:validation:Required
	OperationName string `json:"operationName"`
	// +kubebuilder:validation:Required
	AloName string `json:"aloName"`
	// +kubebuilder:validation:Optional
	TloLatencyAvgTimeSlotInSec int64 `json:"tloLatencyAvgTimeSlotInSec"`
}

// TloStatus defines the observed state of Tlo
type TloStatus struct {
	Conditions            []TloCondition `json:"conditions"`
	Status                bool           `json:"status"`
	Reason                string         `json:"reason"`
	Message               string         `json:"message"`
	LastProbeTime         string         `json:"lastProbeTime"`
	LastTransitionTime    string         `json:"lastTransitionTime"`
	TotalTlis             int32          `json:"totalTlis"`
	TotalTransactionTypes int32          `json:"totalTransactionTypes"`
	ScalingFlag           int32          `json:"scalingFlag"`
}

// IndicatorMetric spec to define the indicator metrics for a LO
type TloIndicatorMetric struct {
	MetricName TloIndicatorMetricName `json:"name"`
	UpperLimit int                    `json:"upper"`
	LowerLimit int                    `json:"lower"`
}

// TLO Indicator Metric Name Values
type TloIndicatorMetricName string

const (
	TLOP50LATENCY TloIndicatorMetricName = "P50LATENCY"
	TLOP90LATENCY TloIndicatorMetricName = "P90LATENCY"
	TLOP99LATENCY TloIndicatorMetricName = "P99LATENCY"
	TLOERRORRATE  TloIndicatorMetricName = "ERRORRATE"
)

// TLO Condition struct for storing current condition of the TLO
type TloCondition struct {
	Type TloConditionType `json:"type"`
	// This ELO Name will be filled only in case of Type=EliReady and not in case of Type=TliReady
	EloName   string             `json:"eloName"`
	EloType   string             `json:"eloType"`
	Status    bool               `json:"status"`
	Reason    TloConditionReason `json:"reason"`
	Message   string             `json:"message"`
	Indicator int32              `json:"indicator"`
}

// TLO Condition Type Values
type TloConditionType string

const (
	// Complete Transaction Indicators Status
	TliReady TloConditionType = "TliReady"
	// ELO in the context of TLO Indicators Status
	EliReady TloConditionType = "EliReady"
	EloReady TloConditionType = "EloReady"
)

// TLO Condition reason
type TloConditionReason string

const (
	TliLatencyBreached TloConditionReason = "LatencyBreached"
	TliMetricsError    TloConditionReason = "MetricsError"
)
