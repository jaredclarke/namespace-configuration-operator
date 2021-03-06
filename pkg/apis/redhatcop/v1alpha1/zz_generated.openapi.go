// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/redhat-cop/namespace-configuration-operator/pkg/apis/redhatcop/v1alpha1.NamespaceConfig":       schema_pkg_apis_redhatcop_v1alpha1_NamespaceConfig(ref),
		"github.com/redhat-cop/namespace-configuration-operator/pkg/apis/redhatcop/v1alpha1.NamespaceConfigSpec":   schema_pkg_apis_redhatcop_v1alpha1_NamespaceConfigSpec(ref),
		"github.com/redhat-cop/namespace-configuration-operator/pkg/apis/redhatcop/v1alpha1.NamespaceConfigStatus": schema_pkg_apis_redhatcop_v1alpha1_NamespaceConfigStatus(ref),
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_NamespaceConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NamespaceConfig is the Schema for the namespaceconfigs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/namespace-configuration-operator/pkg/apis/redhatcop/v1alpha1.NamespaceConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/namespace-configuration-operator/pkg/apis/redhatcop/v1alpha1.NamespaceConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/namespace-configuration-operator/pkg/apis/redhatcop/v1alpha1.NamespaceConfigSpec", "github.com/redhat-cop/namespace-configuration-operator/pkg/apis/redhatcop/v1alpha1.NamespaceConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_NamespaceConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NamespaceConfigSpec defines the desired state of NamespaceConfig",
				Properties: map[string]spec.Schema{
					"selector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector", "k8s.io/apimachinery/pkg/runtime.RawExtension"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_NamespaceConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NamespaceConfigStatus defines the observed state of NamespaceConfig",
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"lastUpdate": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
