//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FalcoCtl) DeepCopyInto(out *FalcoCtl) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FalcoCtl.
func (in *FalcoCtl) DeepCopy() *FalcoCtl {
	if in == nil {
		return nil
	}
	out := new(FalcoCtl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FalcoServiceConfig) DeepCopyInto(out *FalcoServiceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.FalcoVersion != nil {
		in, out := &in.FalcoVersion, &out.FalcoVersion
		*out = new(string)
		**out = **in
	}
	if in.AutoUpdate != nil {
		in, out := &in.AutoUpdate, &out.AutoUpdate
		*out = new(bool)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(string)
		**out = **in
	}
	if in.FalcoCtl != nil {
		in, out := &in.FalcoCtl, &out.FalcoCtl
		*out = new(FalcoCtl)
		**out = **in
	}
	if in.Gardener != nil {
		in, out := &in.Gardener, &out.Gardener
		*out = new(Gardener)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomWebhook != nil {
		in, out := &in.CustomWebhook, &out.CustomWebhook
		*out = new(Webhook)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FalcoServiceConfig.
func (in *FalcoServiceConfig) DeepCopy() *FalcoServiceConfig {
	if in == nil {
		return nil
	}
	out := new(FalcoServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FalcoServiceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gardener) DeepCopyInto(out *Gardener) {
	*out = *in
	if in.UseFalcoRules != nil {
		in, out := &in.UseFalcoRules, &out.UseFalcoRules
		*out = new(bool)
		**out = **in
	}
	if in.UseFalcoIncubatingRules != nil {
		in, out := &in.UseFalcoIncubatingRules, &out.UseFalcoIncubatingRules
		*out = new(bool)
		**out = **in
	}
	if in.UseFalcoSandboxRules != nil {
		in, out := &in.UseFalcoSandboxRules, &out.UseFalcoSandboxRules
		*out = new(bool)
		**out = **in
	}
	if in.RuleRefs != nil {
		in, out := &in.RuleRefs, &out.RuleRefs
		*out = make([]Rule, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gardener.
func (in *Gardener) DeepCopy() *Gardener {
	if in == nil {
		return nil
	}
	out := new(Gardener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.CustomHeaders != nil {
		in, out := &in.CustomHeaders, &out.CustomHeaders
		*out = new(string)
		**out = **in
	}
	if in.Checkcerts != nil {
		in, out := &in.Checkcerts, &out.Checkcerts
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhook.
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}
