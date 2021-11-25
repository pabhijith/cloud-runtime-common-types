package common_types

// AloSpec defines the desired state of Alo
type AloSpec struct {
	// +kubebuilder:validation:Optional
	Description string `json:"description"`
	// +kubebuilder:validation:Required
	Gns string `json:"gns"`
	// +kubebuilder:validation:Required
	CloudruntimeIdentifier string `json:"cloudruntimeIdentifier"`
	// +kubebuilder:validation:Optional
	Tlos []string `json:"tlos"`
	// +kubebuilder:validation:Optional
	Clo string `json:"clo"`
	// +kubebuilder:validation:Optional
	Elos []string `json:"elos"`
	// +kubebuilder:validation:Optional
	Slos []AloSloRef `json:"slos"`
	// +kubebuilder:validation:Optional
	Remediations AloRemediations `json:"remediations"`
}

// AloStatus defines the observed state of Alo
type AloStatus struct {
	ClusterNamespace   string         `json:"clusterNamespace"`
	Conditions         []AloCondition `json:"conditions"`
	Status             bool           `json:"status"`
	Reason             string         `json:"reason"`
	Message            string         `json:"message"`
	LastProbeTime      string         `json:"lastProbeTime"`
	LastTransitionTime string         `json:"lastTransitionTime"`
	TotalTlos          int32          `json:"totalTlos"`
	TotalClos          int32          `json:"totalClos"`
	TotalElos          int32          `json:"totalElos"`
	TotalSlos          int32          `json:"totalSlos"`
	// +kubebuilder:validation:Optional
	RemediationStatus []RemediationStatus `json:"remediationStatus"`
}

// ALO to SLO references
type AloSloRef struct {
	Name    string `json:"name"`
	Service string `json:"service"`
}

// ALO Spec Remediations
type AloRemediations struct {
	// +kubebuilder:validation:Optional
	CloudBurst AloRemediationCloudBurst `json:"cloudBurst"`

	// +kubebuilder:validation:Optional
	HttpWebhooks []AloRemediationHttpWebhook `json:"httpWebhooks"`

	// +kubebuilder:validation:Optional
	Tas []AloRemediationTAS `json:"tas"`

	// +kubebuilder:validation:Optional
	VeloEdge []AloRemediationVeloEdge `json:"veloEdge"`

	// +kubebuilder:validation:Optional
	AwsAsg AloRemediationAwsAsg `json:"awsAsg"`
}

// ALO SPec Remediation Cloud Burst
type AloRemediationCloudBurst struct {
	BurstMode          BurstMode          `json:"burstMode"`
	ProvisioningPolicy ProvisioningPolicy `json:"provisioningPolicy"`
	ReplicationOption  ReplicationOption  `json:"replicationOption"`
	// +kubebuilder:validation:Optional
	BurstCapacitySpec AloRemediationBurstCapacitySpec `json:"burstCapacitySpec"`
	MaxBurst          int32                           `json:"maxBurst"`
	// Quota                          AloRemediationCloudBurstQuota `json:"quota"`
	AppDeploymentManifestsRepoPath string `json:"appDeploymentManifestsRepoPath"`
	AppDeploymentManifestsDirPath  string `json:"appDeploymentManifestsDirPath"`
}

// BurstMode Type Values
type BurstMode string

const (
	BurstModeReserved BurstMode = "Reserved"
	BurstModeShared   BurstMode = "Shared"
)

// ProvisioningPolicy Type Values
type ProvisioningPolicy string

const (
	ProvisioningPolicyOD ProvisioningPolicy = "On-Demand"
	ProvisioningPolicyPP ProvisioningPolicy = "Pre-Provisioned"
)

// ReplicationOption Type Values
type ReplicationOption string

const (
	ReplicationOptionCurrent ReplicationOption = "Current"
	ReplicationOptionAll     ReplicationOption = "All"
)

// AloRemediationBurstCapacitySpec ALO Spec Remediation Cloud Burst Quota
type AloRemediationBurstCapacitySpec struct {
	WorkerNodes    int32  `json:"workerNodes"`
	WorkerNodeType string `json:"workerNodeType"`
	Region         string `json:"region"`
}

// ALO Condition struct for storing current condition of the ALO
type AloCondition struct {
	Type    AloConditionType   `json:"type"`
	Status  bool               `json:"status"`
	Reason  AloConditionReason `json:"reason"`
	Message string             `json:"message"`
}

// ALO Condition Type Values
type AloConditionType string

const (
	TloReady          AloConditionType = "TloReady"
	CloReady          AloConditionType = "CloReady"
	SloReady          AloConditionType = "SloReady"
	BurstClusterReady AloConditionType = "BurstClusterReady"
	HTTPWebhookReady  AloConditionType = "HTTPWebhookReady"
)

// ALO Condition reason
type AloConditionReason string

const (
	AloObjectMissing AloConditionReason = "ObjectMissing"
	TloStatusBreach  AloConditionReason = "TloStatusBreach"
	TloMetricsError  AloConditionReason = "TloMetricsError"
	CloStatusBreach  AloConditionReason = "CloStatusBreach"
	CloMetricsError  AloConditionReason = "CloMetricsError"
	EloStatusBreach  AloConditionReason = "EloStatusBreach"
	EloMetricsError  AloConditionReason = "EloMetricsError"
	SloStatusBreach  AloConditionReason = "SloStatusBreach"
	SloMetricsError  AloConditionReason = "SloMetricsError"
	CrClusterError   AloConditionReason = "CRClusterError"
)

type RemediationStatus struct {
	MetricName      string `json:"metricName"`
	RemediationType string `json:"remediationType"`
	CurrentVal      int32  `json:"currentVal"`
	TriggerTime     string `json:"triggerTime"`
}

// ALO Remediation Webhooks
type AloRemediationHttpWebhook struct {
	Url string `json:"url"`
	Elo string `json:"elo"`
}

// AloRemediationTAS :
type AloRemediationTAS struct {
	Elo         string `json:"elo"`
	ServiceName string `json:"serviceName"`
	// +kubebuilder:validation:Optional
	MetricName string `json:"metricName"`
	// +kubebuilder:validation:Optional
	MinInstances int `json:"minInstances"`
	// +kubebuilder:validation:Optional
	MaxInstances int `json:"maxInstances"`
}

// AloRemediationVeloEdge :
type AloRemediationVeloEdge struct {
	Elo         string `json:"elo"`
	ServiceName string `json:"serviceName"`
}

// AloRemediationAwsAsg :
type AloRemediationAwsAsg struct {
	Elo             string `json:"elo"`
	MetricName      string `json:"metricName"`
	AwsAccessKeyID  string `json:"awsAccessKeyId"`
	AwsAccessSecret string `json:"awsAccessSecret"`
	// +kubebuilder:validation:Optional
	AwsSessionToken string `json:"awsSessionToken"`
}
