/*
Copyright 2017 The Kubernetes Authors All rights reserved.

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

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
type Event struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Required. Time when this Event was first observed.
	EventTime metav1.MicroTime `json:"eventTime" protobuf:"bytes,2,opt,name=eventTime"`
	// Optional. Data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `json:"series,omitempty" protobuf:"bytes,3,opt,name=series"`
	// Required. Kubernetes component that generated the Event.
	Origin EventSource `json:"eventSource" protobuf:"bytes,4,opt,name=eventSource"`
	// Required. What Origin did/failed to do.
	Action EventAction `json:"action" protobuf:"bytes,5,opt,name=action"`
	// Optional. On what Origin acted upon.
	Object *v1.ObjectReference `json:"object,omitempty" protobuf:"bytes,6,opt,name=object"`
	// Optional secondary Object for more complex actions.
	SecondaryObject *v1.ObjectReference `json:"secondaryObject,omitempty" protobuf:"bytes,7,opt,name=secondaryObject"`
	// Required. Severity of the Event.
	Severity EventSeverity `json:"severity" protobuf:"bytes,8,opt,name=severity"`

	// Optional. A human-readable description of the status of this operation.
	// TODO: decide on maximum length.
	Message string `json:"message,omitempty" protobuf:"bytes,9,opt,name=message"`
}

type EventSeries struct {
	// Unique identifier of the Event series.
	UID string `json:"uid" protobuf:"bytes,1,opt,name=uid"`
	// Number of Events in this series until last heartbeat.
	Count int32 `json:"count" protobuf:"varint,2,opt,name=count"`
	// Time when last Event from the series was seen before last heartbeat.
	LastObservedTime metav1.MicroTime `json:"lastObservedTime" protobuf:"bytes,3,opt,name=lastObservedTime"`
	// Last time when seried data was updated.
	LastHeartbeat metav1.MicroTime `json:"lastHeartbeat" protobuf:"bytes,4,opt,name=lastHeartbeat"`
	// Information whether this series is finished.
	FinishMarker bool `json:"finishMarker" protobuf:"varint,5,opt,name=finishMarker"`
}

// Information about component that emits Event
type EventSource struct {
	// Component from which the event is generated.
	Component string `json:"component" protobuf:"bytes,1,opt,name=component"`
	// String that identifies the component emitting the Event
	ID string `json:"id" protobuf:"bytes,2,opt,name=id"`
}

// Information about action that was taken
type EventAction struct {
	Action string `json:"action" protobuf:"bytes,1,opt,name=action"`
	Reason string `json:"reason" protobuf:"bytes,2,opt,name=reason"`
}

type EventSeverity string

const (
	// e.g. Pod scheduled
	EventSeverityInfo EventSeverity = "Info"

	// e.g. Pod unable to schedule, cluster autoscaler adding Node
	EventSeverityWarning EventSeverity = "Warning"

	// e.g. Node unreachable - something is wrong, but there's a
	// possibility it will recover by itself
	EventSeverityError EventSeverity = "Error"

	// e.g. ClusterAutoscaler out of quota - something needs user
	// action
	EventSeverityFatal EventSeverity = "Fatal"
)
