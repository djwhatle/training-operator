//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2021 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kubeflow/common/pkg/apis/common/v1.JobCondition":             schema_pkg_apis_common_v1_JobCondition(ref),
		"github.com/kubeflow/common/pkg/apis/common/v1.JobStatus":                schema_pkg_apis_common_v1_JobStatus(ref),
		"github.com/kubeflow/common/pkg/apis/common/v1.ReplicaSpec":              schema_pkg_apis_common_v1_ReplicaSpec(ref),
		"github.com/kubeflow/common/pkg/apis/common/v1.ReplicaStatus":            schema_pkg_apis_common_v1_ReplicaStatus(ref),
		"github.com/kubeflow/common/pkg/apis/common/v1.RunPolicy":                schema_pkg_apis_common_v1_RunPolicy(ref),
		"github.com/kubeflow/common/pkg/apis/common/v1.SchedulingPolicy":         schema_pkg_apis_common_v1_SchedulingPolicy(ref),
		"github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1.TFJob":     schema_pkg_apis_tensorflow_v1_TFJob(ref),
		"github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1.TFJobList": schema_pkg_apis_tensorflow_v1_TFJobList(ref),
		"github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1.TFJobSpec": schema_pkg_apis_tensorflow_v1_TFJobSpec(ref),
	}
}

func schema_pkg_apis_common_v1_JobCondition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobCondition describes the state of the job at a certain point.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of job condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "The reason for the condition's last transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Description: "The last time this condition was updated.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_common_v1_JobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobStatus represents the current observed state of the training Job.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is an array of current observed job conditions.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kubeflow/common/pkg/apis/common/v1.JobCondition"),
									},
								},
							},
						},
					},
					"replicaStatuses": {
						SchemaProps: spec.SchemaProps{
							Description: "ReplicaStatuses is map of ReplicaType and ReplicaStatus, specifies the status of each replica.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kubeflow/common/pkg/apis/common/v1.ReplicaStatus"),
									},
								},
							},
						},
					},
					"startTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents time when the job was acknowledged by the job controller. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"completionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"lastReconcileTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents last time when the job was reconciled. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"conditions", "replicaStatuses"},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/common/pkg/apis/common/v1.JobCondition", "github.com/kubeflow/common/pkg/apis/common/v1.ReplicaStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_common_v1_ReplicaSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReplicaSpec is a description of the replica",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"template": {
						SchemaProps: spec.SchemaProps{
							Description: "Template is the object that describes the pod that will be created for this replica. RestartPolicy in PodTemplateSpec will be overide by RestartPolicy in ReplicaSpec",
							Ref:         ref("k8s.io/api/core/v1.PodTemplateSpec"),
						},
					},
					"restartPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Restart policy for all replicas within the job. One of Always, OnFailure, Never and ExitCode. Default to Never.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.PodTemplateSpec"},
	}
}

func schema_pkg_apis_common_v1_ReplicaStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ReplicaStatus represents the current observed state of the replica.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"active": {
						SchemaProps: spec.SchemaProps{
							Description: "The number of actively running pods.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"succeeded": {
						SchemaProps: spec.SchemaProps{
							Description: "The number of pods which reached phase Succeeded.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"failed": {
						SchemaProps: spec.SchemaProps{
							Description: "The number of pods which reached phase Failed.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"labelSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_common_v1_RunPolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RunPolicy encapsulates various runtime policies of the distributed training job, for example how to clean up resources and how long the job can stay active.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cleanPodPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "CleanPodPolicy defines the policy to kill pods after the job completes. Default to Running.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ttlSecondsAfterFinished": {
						SchemaProps: spec.SchemaProps{
							Description: "TTLSecondsAfterFinished is the TTL to clean up jobs. It may take extra ReconcilePeriod seconds for the cleanup, since reconcile gets called periodically. Default to infinite.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"activeDeadlineSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"backoffLimit": {
						SchemaProps: spec.SchemaProps{
							Description: "Optional number of retries before marking this job failed.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"schedulingPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "SchedulingPolicy defines the policy related to scheduling, e.g. gang-scheduling",
							Ref:         ref("github.com/kubeflow/common/pkg/apis/common/v1.SchedulingPolicy"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/common/pkg/apis/common/v1.SchedulingPolicy"},
	}
}

func schema_pkg_apis_common_v1_SchedulingPolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SchedulingPolicy encapsulates various scheduling policies of the distributed training job, for example `minAvailable` for gang-scheduling.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"minAvailable": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"queue": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"minResources": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
									},
								},
							},
						},
					},
					"priorityClass": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/api/resource.Quantity"},
	}
}

func schema_pkg_apis_tensorflow_v1_TFJob(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TFJob represents a TFJob resource.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
							Description: "Specification of the desired state of the TFJob.",
							Ref:         ref("github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1.TFJobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recently observed status of the TFJob. Populated by the system. Read-only.",
							Ref:         ref("github.com/kubeflow/common/pkg/apis/common/v1.JobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/common/pkg/apis/common/v1.JobStatus", "github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1.TFJobSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_tensorflow_v1_TFJobList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TFJobList is a list of TFJobs.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard list metadata.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Description: "List of TFJobs.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1.TFJob"),
									},
								},
							},
						},
					},
				},
				Required: []string{"items"},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/training-operator/pkg/apis/tensorflow/v1.TFJob", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_pkg_apis_tensorflow_v1_TFJobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TFJobSpec is a desired state description of the TFJob.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"runPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "RunPolicy encapsulates various runtime policies of the distributed training job, for example how to clean up resources and how long the job can stay active.",
							Ref:         ref("github.com/kubeflow/common/pkg/apis/common/v1.RunPolicy"),
						},
					},
					"successPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "SuccessPolicy defines the policy to mark the TFJob as succeeded. Default to \"\", using the default rules.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tfReplicaSpecs": {
						SchemaProps: spec.SchemaProps{
							Description: "A map of TFReplicaType (type) to ReplicaSpec (value). Specifies the TF cluster configuration. For example,\n  {\n    \"PS\": ReplicaSpec,\n    \"Worker\": ReplicaSpec,\n  }",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kubeflow/common/pkg/apis/common/v1.ReplicaSpec"),
									},
								},
							},
						},
					},
					"enableDynamicWorker": {
						SchemaProps: spec.SchemaProps{
							Description: "A switch to enable dynamic worker",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"runPolicy", "tfReplicaSpecs"},
			},
		},
		Dependencies: []string{
			"github.com/kubeflow/common/pkg/apis/common/v1.ReplicaSpec", "github.com/kubeflow/common/pkg/apis/common/v1.RunPolicy"},
	}
}
