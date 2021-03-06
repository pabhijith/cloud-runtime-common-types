package common_types

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AloSpec) DeepCopyInto(out *AloSpec) {
	*out = *in
	if in.Tlos != nil {
		in, out := &in.Tlos, &out.Tlos
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Elos != nil {
		in, out := &in.Elos, &out.Elos
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Slos != nil {
		in, out := &in.Slos, &out.Slos
		*out = make([]AloSloRef, len(*in))
		copy(*out, *in)
	}
	in.Remediations.DeepCopyInto(&out.Remediations)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AloSpec.
func (in *AloSpec) DeepCopy() *AloSpec {
	if in == nil {
		return nil
	}
	out := new(AloSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AloRemediations) DeepCopyInto(out *AloRemediations) {
	*out = *in
	out.CloudBurst = in.CloudBurst
	if in.HttpWebhooks != nil {
		in, out := &in.HttpWebhooks, &out.HttpWebhooks
		*out = make([]AloRemediationHttpWebhook, len(*in))
		copy(*out, *in)
	}
	if in.Tas != nil {
		in, out := &in.Tas, &out.Tas
		*out = make([]AloRemediationTAS, len(*in))
		copy(*out, *in)
	}
	if in.VeloEdge != nil {
		in, out := &in.VeloEdge, &out.VeloEdge
		*out = make([]AloRemediationVeloEdge, len(*in))
		copy(*out, *in)
	}
	out.AwsAsg = in.AwsAsg
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AloRemediations.
func (in *AloRemediations) DeepCopy() *AloRemediations {
	if in == nil {
		return nil
	}
	out := new(AloRemediations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AloStatus) DeepCopyInto(out *AloStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AloCondition, len(*in))
		copy(*out, *in)
	}
	if in.RemediationStatus != nil {
		in, out := &in.RemediationStatus, &out.RemediationStatus
		*out = make([]RemediationStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AloStatus.
func (in *AloStatus) DeepCopy() *AloStatus {
	if in == nil {
		return nil
	}
	out := new(AloStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloSpec) DeepCopyInto(out *CloSpec) {
	*out = *in
	if in.Clis != nil {
		in, out := &in.Clis, &out.Clis
		*out = make([]CloIndicatorMetric, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloSpec.
func (in *CloSpec) DeepCopy() *CloSpec {
	if in == nil {
		return nil
	}
	out := new(CloSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloStatus) DeepCopyInto(out *CloStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CloCondition, len(*in))
		copy(*out, *in)
	}
	if in.BurstClusterSuportingGnses != nil {
		in, out := &in.BurstClusterSuportingGnses, &out.BurstClusterSuportingGnses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloStatus.
func (in *CloStatus) DeepCopy() *CloStatus {
	if in == nil {
		return nil
	}
	out := new(CloStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLoSpec) DeepCopyInto(out *ExternalLoSpec) {
	*out = *in
	if in.Elis != nil {
		in, out := &in.Elis, &out.Elis
		*out = make([]EloIndicatorMetric, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLoSpec.
func (in *ExternalLoSpec) DeepCopy() *ExternalLoSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalLoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalLoStatus) DeepCopyInto(out *ExternalLoStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]EloCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLoStatus.
func (in *ExternalLoStatus) DeepCopy() *ExternalLoStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalLoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TloSpec) DeepCopyInto(out *TloSpec) {
	*out = *in
	if in.Tlis != nil {
		in, out := &in.Tlis, &out.Tlis
		*out = make([]TloIndicatorMetric, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TloSpec.
func (in *TloSpec) DeepCopy() *TloSpec {
	if in == nil {
		return nil
	}
	out := new(TloSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TloStatus) DeepCopyInto(out *TloStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TloCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TloStatus.
func (in *TloStatus) DeepCopy() *TloStatus {
	if in == nil {
		return nil
	}
	out := new(TloStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AloRemediationHttpWebhook) DeepCopyInto(out *AloRemediationHttpWebhook) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AloRemediationHttpWebhook.
func (in *AloRemediationHttpWebhook) DeepCopy() *AloRemediationHttpWebhook {
	if in == nil {
		return nil
	}
	out := new(AloRemediationHttpWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EloCondition) DeepCopyInto(out *EloCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EloCondition.
func (in *EloCondition) DeepCopy() *EloCondition {
	if in == nil {
		return nil
	}
	out := new(EloCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EloIndicatorMetric) DeepCopyInto(out *EloIndicatorMetric) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EloIndicatorMetric.
func (in *EloIndicatorMetric) DeepCopy() *EloIndicatorMetric {
	if in == nil {
		return nil
	}
	out := new(EloIndicatorMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TloCondition) DeepCopyInto(out *TloCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TloCondition.
func (in *TloCondition) DeepCopy() *TloCondition {
	if in == nil {
		return nil
	}
	out := new(TloCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TloIndicatorMetric) DeepCopyInto(out *TloIndicatorMetric) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TloIndicatorMetric.
func (in *TloIndicatorMetric) DeepCopy() *TloIndicatorMetric {
	if in == nil {
		return nil
	}
	out := new(TloIndicatorMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloCondition) DeepCopyInto(out *CloCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloCondition.
func (in *CloCondition) DeepCopy() *CloCondition {
	if in == nil {
		return nil
	}
	out := new(CloCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloIndicatorMetric) DeepCopyInto(out *CloIndicatorMetric) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloIndicatorMetric.
func (in *CloIndicatorMetric) DeepCopy() *CloIndicatorMetric {
	if in == nil {
		return nil
	}
	out := new(CloIndicatorMetric)
	in.DeepCopyInto(out)
	return out
}
