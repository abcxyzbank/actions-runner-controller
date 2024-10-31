package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var githubScaleSetControllerSubsystem = "gha_controller"

var labels = []string{
	"name",
	"namespace",
	"repository",
	"organization",
	"enterprise",
	"listener_type",
}

type CommonLabels struct {
	Name         string
	Namespace    string
	Repository   string
	Organization string
	Enterprise   string
	ListenerType string
}

func (l *CommonLabels) labels() prometheus.Labels {
	return prometheus.Labels{
		"name":         l.Name,
		"namespace":    l.Namespace,
		"repository":   l.Repository,
		"organization": l.Organization,
		"enterprise":   l.Enterprise,
		"listener_type": l.ListenerType,
	}
}

var (
	pendingEphemeralRunners = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "pending_ephemeral_runners",
			Help:      "Number of ephemeral runners in a pending state.",
		},
		labels,
	)
	runningEphemeralRunners = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "running_ephemeral_runners",
			Help:      "Number of ephemeral runners in a running state.",
		},
		labels,
	)
	failedEphemeralRunners = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "failed_ephemeral_runners",
			Help:      "Number of ephemeral runners in a failed state.",
		},
		labels,
	)
	runningListeners = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "running_listeners",
			Help:      "Number of listeners in a running state.",
		},
		labels,
	)
	runningJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "running_jobs",
			Help:      "Number of running jobs.",
		},
		labels,
	)
	queuedJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "queued_jobs",
			Help:      "Number of queued jobs.",
		},
		labels,
	)
	idleRunners = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "idle_runners",
			Help:      "Number of idle runners.",
		},
		labels,
	)
	failedJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "failed_jobs",
			Help:      "Number of failed jobs.",
		},
		labels,
	)
	averageJobDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "average_job_duration",
			Help:      "Average job duration.",
		},
		labels,
	)
	resourceUtilizationCPU = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "resource_utilization_cpu",
			Help:      "Resource utilization (CPU).",
		},
		labels,
	)
	resourceUtilizationMemory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: githubScaleSetControllerSubsystem,
			Name:      "resource_utilization_memory",
			Help:      "Resource utilization (Memory).",
		},
		labels,
	)
)

func RegisterMetrics() {
	metrics.Registry.MustRegister(
		pendingEphemeralRunners,
		runningEphemeralRunners,
		failedEphemeralRunners,
		runningListeners,
		runningJobs,
		queuedJobs,
		idleRunners,
		failedJobs,
		averageJobDuration,
		resourceUtilizationCPU,
		resourceUtilizationMemory,
	)
}

func SetEphemeralRunnerCountsByStatus(commonLabels CommonLabels, pending, running, failed int) {
	pendingEphemeralRunners.With(commonLabels.labels()).Set(float64(pending))
	runningEphemeralRunners.With(commonLabels.labels()).Set(float64(running))
	failedEphemeralRunners.With(commonLabels.labels()).Set(float64(failed))
}

func AddRunningListener(commonLabels CommonLabels) {
	runningListeners.With(commonLabels.labels()).Set(1)
}

func SubRunningListener(commonLabels CommonLabels) {
	runningListeners.With(commonLabels.labels()).Set(0)
}

func SetRunningJobs(commonLabels CommonLabels, count int) {
	runningJobs.With(commonLabels.labels()).Set(float64(count))
}

func SetQueuedJobs(commonLabels CommonLabels, count int) {
	queuedJobs.With(commonLabels.labels()).Set(float64(count))
}

func SetIdleRunners(commonLabels CommonLabels, count int) {
	idleRunners.With(commonLabels.labels()).Set(float64(count))
}

func SetFailedJobs(commonLabels CommonLabels, count int) {
	failedJobs.With(commonLabels.labels()).Set(float64(count))
}

func SetAverageJobDuration(commonLabels CommonLabels, duration float64) {
	averageJobDuration.With(commonLabels.labels()).Set(duration)
}

func SetResourceUtilizationCPU(commonLabels CommonLabels, utilization float64) {
	resourceUtilizationCPU.With(commonLabels.labels()).Set(utilization)
}

func SetResourceUtilizationMemory(commonLabels CommonLabels, utilization float64) {
	resourceUtilizationMemory.With(commonLabels.labels()).Set(utilization)
}
