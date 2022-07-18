/*
Copyright © 2021 Ci4Rail GmbH <engineering@ci4rail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StreamSpec defines the configuration of a Stream
type StreamSpec struct {
	// Name of the stream
	Name string `json:"name"`

	// Subject defines the subjects of the stream
	Subjects []string `json:"subjects"`

	// Public defines if the stream shall be exported
	// +kubebuilder:default:=false
	Public bool `json:"public,omitempty"`

	// Global defines if the stream is local only or global
	// +kubebuilder:default:=true
	Global bool `json:"global,omitempty"`

	// Streams are stored on the server, this can be one of many backends and all are usable in clustering mode.
	// Allowed values are: file, memory
	// +kubebuilder:default:=file
	Storage string `json:"storage,omitempty"`

	// Messages are retained either based on limits like size and age (Limits), as long as there are Consumers (Interest) or until any worker processed them (Work Queue)
	// Allowed values are: limits, interest, workqueue
	// +kubebuilder:default:=limits
	Retention string `json:"retention,omitempty"`

	// MaxMsgsPerSubject defines the amount of messages to keep in the store for this Stream per unique subject, when exceeded oldest messages are removed, -1 for unlimited.
	// +kubebuilder:default:=-1
	MaxMsgsPerSubject int64 `json:"maxMsgsPerSubject,omitempty"`

	// MaxMsgs defines the amount of messages to keep in the store for this Stream, when exceeded oldest messages are removed, -1 for unlimited.
	// +kubebuilder:default:=-1
	MaxMsgs int64 `json:"maxMsgs,omitempty"`

	// MaxBytes defines the combined size of all messages in a Stream, when exceeded oldest messages are removed, -1 for unlimited.
	// +kubebuilder:default:=-1
	MaxBytes int64 `json:"maxBytes,omitempty"`

	// MaxAge defines the oldest messages that can be stored in the Stream, any messages older than this period will be removed, -1 for unlimited. Supports units (s)econds, (m)inutes, (h)ours, (d)ays, (M)onths, (y)ears.
	// +kubebuilder:default:="1y"
	MaxAge string `json:"maxAge,omitempty"`

	// MaxMsgSize defines the maximum size any single message may be to be accepted by the Stream.
	// +kubebuilder:default:=-1
	MaxMsgSize int32 `json:"maxMsgSize,omitempty"`

	// Discard defines if once the Stream reach it's limits of size or messages the 'new' policy will prevent further messages from being added while 'old' will delete old messages.
	// Allowed values are: new, old
	// +kubebuilder:default:="old"
	Discard string `json:"discard,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// App is the name of the distrubuted application the network is created for.
	App string `json:"app"`

	// Namespace is the namespace the credentials shall be stored in..
	Namespace string `json:"namespace,omitempty"`

	// Streams is a list of streams in the network.
	Streams []StreamSpec `json:"streams"`
}

// ParticipatingSpec defines the current state of participating nodes and pods
type ParticipatingSpec struct {
	// Nodes is a list of kubernetes nodes that currently are hosting participating components.
	// key is node, value is current state of the participating node ("pending", "created", "terminating")
	Nodes map[string]string `json:"nodes"`

	// Pods is a map of node names to a list of pod names indicating the pods running on the node.
	// key is node, value is list of pod names
	Pods map[string][]string `json:"pods"`

	// PodsCreated is a map of node names to a list of pod names indicating that the pods are being created.
	// key is node, value is list of pod names
	PodsCreating map[string][]string `json:"podsCreating"`

	// PodsTerminating is a map of node names to a list of pod names indicating that the pods are terminating.
	// key is node, value is list of pod names
	PodsTerminating map[string][]string `json:"podsTerminating"`

	// Components is a list of participating components in the network with their corresponding types ("edge" or "cloud").
	Components map[string]string `json:"components"`
}

// NetworkInfoSpec defines the observed state of Network
type NetworkInfoSpec struct {
	Participating ParticipatingSpec `json:"participating"`
	// UsedAccount is the account that is used for the network.
	UsedAccount string `json:"usedAccount,omitempty"`
}

//+kubebuilder:object:root=true
//+genclient

// Network is the Schema for the networks API
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NetworkSpec     `json:"spec,omitempty"`
	Info NetworkInfoSpec `json:"info,omitempty"`
}

//+kubebuilder:object:root=true

// NetworkList contains a list of Network
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
