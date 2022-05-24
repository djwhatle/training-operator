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
// limitations under the License

package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// Define const label values
const (
	NotDistributed   = "not_distributed"
	Distributed      = "distributed"
	NotGangScheduled = "not_gang_scheduled"
	GangScheduled    = "gang_scheduled"
)

// Define all the prometheus counters for all jobs
var (
	jobsCreatedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "training_operator_jobs_created_total",
			Help: "Counts number of jobs created",
		},
		[]string{"job_namespace", "framework", "distributed", "gang_scheduling"},
	)
	jobsDeletedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "training_operator_jobs_deleted_total",
			Help: "Counts number of jobs deleted",
		},
		[]string{"job_namespace", "framework", "distributed", "gang_scheduling"},
	)
	jobsSuccessfulCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "training_operator_jobs_successful_total",
			Help: "Counts number of jobs successful",
		},
		[]string{"job_namespace", "framework", "distributed", "gang_scheduling"},
	)
	jobsFailedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "training_operator_jobs_failed_total",
			Help: "Counts number of jobs failed",
		},
		[]string{"job_namespace", "framework", "distributed", "gang_scheduling"},
	)
	jobsRestartedCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "training_operator_jobs_restarted_total",
			Help: "Counts number of jobs restarted",
		},
		[]string{"job_namespace", "framework", "distributed", "gang_scheduling"},
	)
)

func init() {
	// Register custom metrics with the global prometheus registry
	metrics.Registry.MustRegister(jobsCreatedCount,
		jobsDeletedCount,
		jobsSuccessfulCount,
		jobsFailedCount,
		jobsRestartedCount)
}

func CreatedJobsCounterInc(job_namespace, framework, distributed, gang_scheduling string) {
	jobsCreatedCount.WithLabelValues(job_namespace, framework, distributed, gang_scheduling).Inc()
}

func DeletedJobsCounterInc(job_namespace, framework, distributed, gang_scheduling string) {
	jobsDeletedCount.WithLabelValues(job_namespace, framework, distributed, gang_scheduling).Inc()
}

func SuccessfulJobsCounterInc(job_namespace, framework, distributed, gang_scheduling string) {
	jobsSuccessfulCount.WithLabelValues(job_namespace, framework, distributed, gang_scheduling).Inc()
}

func FailedJobsCounterInc(job_namespace, framework, distributed, gang_scheduling string) {
	jobsFailedCount.WithLabelValues(job_namespace, framework, distributed, gang_scheduling).Inc()
}

func RestartedJobsCounterInc(job_namespace, framework, distributed, gang_scheduling string) {
	jobsRestartedCount.WithLabelValues(job_namespace, framework, distributed, gang_scheduling).Inc()
}

func MapBoolToDistributed(distributed bool) string {
	if distributed {
		return Distributed
	} else {
		return NotDistributed
	}
}

func MapBoolToGangScheduling(gangScheduled bool) string {
	if gangScheduled {
		return GangScheduled
	} else {
		return NotGangScheduled
	}
}
