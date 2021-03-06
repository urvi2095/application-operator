// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v2alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2.Application":       schema_pkg_apis_application_v2alpha2_Application(ref),
		"github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2.ApplicationSpec":   schema_pkg_apis_application_v2alpha2_ApplicationSpec(ref),
		"github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2.ApplicationStatus": schema_pkg_apis_application_v2alpha2_ApplicationStatus(ref),
	}
}

func schema_pkg_apis_application_v2alpha2_Application(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Application is the Schema for the applications API",
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
							Ref: ref("github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2.ApplicationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2.ApplicationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2.ApplicationSpec", "github.com/urvi2095/application-operator/pkg/apis/application/v2alpha2.ApplicationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_application_v2alpha2_ApplicationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApplicationSpec defines the desired state of Application",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"applicationVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"replicas", "applicationVersion"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_application_v2alpha2_ApplicationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApplicationStatus defines the observed state of Application",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"applicationVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"replicas", "applicationVersion"},
			},
		},
		Dependencies: []string{},
	}
}
