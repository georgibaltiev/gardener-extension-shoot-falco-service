//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	service "github.com/gardener/gardener-extension-falco/pkg/apis/service"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*FalcoServiceConfig)(nil), (*service.FalcoServiceConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_FalcoServiceConfig_To_service_FalcoServiceConfig(a.(*FalcoServiceConfig), b.(*service.FalcoServiceConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*service.FalcoServiceConfig)(nil), (*FalcoServiceConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_service_FalcoServiceConfig_To_v1alpha1_FalcoServiceConfig(a.(*service.FalcoServiceConfig), b.(*FalcoServiceConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_FalcoServiceConfig_To_service_FalcoServiceConfig(in *FalcoServiceConfig, out *service.FalcoServiceConfig, s conversion.Scope) error {
	out.FalcoVersion = (*string)(unsafe.Pointer(in.FalcoVersion))
	out.UseFalcoIncubatingRules = in.UseFalcoIncubatingRules
	out.UseFalcoSandboxRules = in.UseFalcoSandboxRules
	out.RuleRefs = *(*[]string)(unsafe.Pointer(&in.RuleRefs))
	return nil
}

// Convert_v1alpha1_FalcoServiceConfig_To_service_FalcoServiceConfig is an autogenerated conversion function.
func Convert_v1alpha1_FalcoServiceConfig_To_service_FalcoServiceConfig(in *FalcoServiceConfig, out *service.FalcoServiceConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_FalcoServiceConfig_To_service_FalcoServiceConfig(in, out, s)
}

func autoConvert_service_FalcoServiceConfig_To_v1alpha1_FalcoServiceConfig(in *service.FalcoServiceConfig, out *FalcoServiceConfig, s conversion.Scope) error {
	out.FalcoVersion = (*string)(unsafe.Pointer(in.FalcoVersion))
	out.UseFalcoIncubatingRules = in.UseFalcoIncubatingRules
	out.UseFalcoSandboxRules = in.UseFalcoSandboxRules
	out.RuleRefs = *(*[]string)(unsafe.Pointer(&in.RuleRefs))
	return nil
}

// Convert_service_FalcoServiceConfig_To_v1alpha1_FalcoServiceConfig is an autogenerated conversion function.
func Convert_service_FalcoServiceConfig_To_v1alpha1_FalcoServiceConfig(in *service.FalcoServiceConfig, out *FalcoServiceConfig, s conversion.Scope) error {
	return autoConvert_service_FalcoServiceConfig_To_v1alpha1_FalcoServiceConfig(in, out, s)
}
