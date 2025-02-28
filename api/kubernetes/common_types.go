/*
Copyright © 2023 XigXog

This Source Code Form is subject to the terms of the Mozilla Public License,
v2.0. If a copy of the MPL was not distributed with this file, You can obtain
one at https://mozilla.org/MPL/2.0/.
*/

// +kubebuilder:object:generate=true
package kubernetes

import (
	corev1 "k8s.io/api/core/v1"
)

type PodSpec struct {
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication
	// controllers and services. [More
	// info](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels).
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations is an unstructured key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata.
	// They are not queryable and should be preserved when modifying objects.
	// [More
	// info](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations).
	Annotations map[string]string `json:"annotations,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a
	// node. Selector which must match a node's labels for the pod to be
	// scheduled on that node. [More
	// info](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/).
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// NodeName is a request to schedule this pod onto a specific node. If it is
	// non-empty, the scheduler simply schedules this pod onto that node,
	// assuming that it fits resource requirements.
	NodeName string `json:"nodeName,omitempty"`
	// If specified, the pod's scheduling constraints
	Affinity *corev1.Affinity `json:"affinity,omitempty"`
	// If specified, the pod's tolerations.
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
}

type ContainerSpec struct {
	// Compute Resources required by this container. Cannot be updated. [More
	// info](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/).
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
	// Periodic probe of container liveness. Container will be restarted if the
	// probe fails. Cannot be updated. [More
	// info](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes).
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty"`
	// Periodic probe of container service readiness. Container will be removed
	// from service endpoints if the probe fails. Cannot be updated. [More
	// info](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes).
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty"`
	// StartupProbe indicates that the Pod has successfully initialized. If
	// specified, no other probes are executed until this completes
	// successfully. If this probe fails, the Pod will be restarted, just as if
	// the livenessProbe failed. This can be used to provide different probe
	// parameters at the beginning of a Pod's lifecycle, when it might take a
	// long time to load data or warm a cache, than during steady-state
	// operation. This cannot be updated. [More
	// info](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes).
	StartupProbe *corev1.Probe `json:"startupProbe,omitempty"`
}

type LoggerSpec struct {
	// +kubebuilder:validation:Enum=debug;info;warn;error
	Level string `json:"level,omitempty"`
	// +kubebuilder:validation:Enum=json;console
	Format string `json:"format,omitempty"`
}
