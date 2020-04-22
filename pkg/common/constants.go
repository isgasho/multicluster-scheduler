/*
 * Copyright 2020 The Multicluster-Scheduler Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

var (
	ProxySchedulerName     = "admiralty-proxy"
	CandidateSchedulerName = "admiralty-candidate"

	LabelAndTaintKeyVirtualKubeletProvider = "virtual-kubelet.io/provider"
	VirtualKubeletProviderName             = "admiralty"

	KeyPrefix = "multicluster.admiralty.io/"

	// annotations on source pod (by user) and proxy pods (copied by mutating admission webhook)

	AnnotationKeyElect = KeyPrefix + "elect"

	// annotations on proxy pods (by mutating admission webhook)

	KeyPrefixSourcePod = KeyPrefix + "sourcepod-"

	AnnotationKeySourcePodManifest = KeyPrefixSourcePod + "manifest"

	// annotations on delegate pod chaperons (by scheduler plugins)

	AnnotationKeyIsReserved      = KeyPrefix + "is-reserved"
	AnnotationKeyIsUnschedulable = KeyPrefix + "is-unschedulable"
	AnnotationKeyIsAllowed       = KeyPrefix + "is-allowed"
	AnnotationKeyIsBound         = KeyPrefix + "is-bound"
	AnnotationKeyBindingFailed   = KeyPrefix + "binding-failed"

	// annotations on delegate services (by global service controller)

	AnnotationKeyIsDelegate = KeyPrefix + "is-delegate"

	// labels on proxy pods and services (used by post-delete hook to clean up finalizers)

	LabelKeyHasFinalizer = KeyPrefix + "has-finalizer"

	LabelKeyParentName = KeyPrefix + "parent-name"
)
