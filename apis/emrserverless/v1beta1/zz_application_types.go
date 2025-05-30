// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationInitParameters struct {

	// –  The CPU architecture of an application. Valid values are ARM64 or X86_64. Default value is X86_64.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// –  The configuration for an application to automatically start on job submission.
	AutoStartConfiguration []AutoStartConfigurationInitParameters `json:"autoStartConfiguration,omitempty" tf:"auto_start_configuration,omitempty"`

	// –  The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration []AutoStopConfigurationInitParameters `json:"autoStopConfiguration,omitempty" tf:"auto_stop_configuration,omitempty"`

	// –  The image configuration applied to all worker types.
	ImageConfiguration []ImageConfigurationInitParameters `json:"imageConfiguration,omitempty" tf:"image_configuration,omitempty"`

	// –  The capacity to initialize when the application is created.
	InitialCapacity []InitialCapacityInitParameters `json:"initialCapacity,omitempty" tf:"initial_capacity,omitempty"`

	// –  Enables the interactive use cases to use when running an application.
	InteractiveConfiguration []InteractiveConfigurationInitParameters `json:"interactiveConfiguration,omitempty" tf:"interactive_configuration,omitempty"`

	// –  The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity []MaximumCapacityInitParameters `json:"maximumCapacity,omitempty" tf:"maximum_capacity,omitempty"`

	// –  The name of the application.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  The network configuration for customer VPC connectivity.
	NetworkConfiguration []NetworkConfigurationInitParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// –  The EMR release version associated with the application.
	ReleaseLabel *string `json:"releaseLabel,omitempty" tf:"release_label,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  The type of application you want to start, such as spark or hive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ApplicationObservation struct {

	// –  The CPU architecture of an application. Valid values are ARM64 or X86_64. Default value is X86_64.
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// ARN of the cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// –  The configuration for an application to automatically start on job submission.
	AutoStartConfiguration []AutoStartConfigurationObservation `json:"autoStartConfiguration,omitempty" tf:"auto_start_configuration,omitempty"`

	// –  The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration []AutoStopConfigurationObservation `json:"autoStopConfiguration,omitempty" tf:"auto_stop_configuration,omitempty"`

	// The ID of the cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// –  The image configuration applied to all worker types.
	ImageConfiguration []ImageConfigurationObservation `json:"imageConfiguration,omitempty" tf:"image_configuration,omitempty"`

	// –  The capacity to initialize when the application is created.
	InitialCapacity []InitialCapacityObservation `json:"initialCapacity,omitempty" tf:"initial_capacity,omitempty"`

	// –  Enables the interactive use cases to use when running an application.
	InteractiveConfiguration []InteractiveConfigurationObservation `json:"interactiveConfiguration,omitempty" tf:"interactive_configuration,omitempty"`

	// –  The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity []MaximumCapacityObservation `json:"maximumCapacity,omitempty" tf:"maximum_capacity,omitempty"`

	// –  The name of the application.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  The network configuration for customer VPC connectivity.
	NetworkConfiguration []NetworkConfigurationObservation `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// –  The EMR release version associated with the application.
	ReleaseLabel *string `json:"releaseLabel,omitempty" tf:"release_label,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// –  The type of application you want to start, such as spark or hive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ApplicationParameters struct {

	// –  The CPU architecture of an application. Valid values are ARM64 or X86_64. Default value is X86_64.
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// –  The configuration for an application to automatically start on job submission.
	// +kubebuilder:validation:Optional
	AutoStartConfiguration []AutoStartConfigurationParameters `json:"autoStartConfiguration,omitempty" tf:"auto_start_configuration,omitempty"`

	// –  The configuration for an application to automatically stop after a certain amount of time being idle.
	// +kubebuilder:validation:Optional
	AutoStopConfiguration []AutoStopConfigurationParameters `json:"autoStopConfiguration,omitempty" tf:"auto_stop_configuration,omitempty"`

	// –  The image configuration applied to all worker types.
	// +kubebuilder:validation:Optional
	ImageConfiguration []ImageConfigurationParameters `json:"imageConfiguration,omitempty" tf:"image_configuration,omitempty"`

	// –  The capacity to initialize when the application is created.
	// +kubebuilder:validation:Optional
	InitialCapacity []InitialCapacityParameters `json:"initialCapacity,omitempty" tf:"initial_capacity,omitempty"`

	// –  Enables the interactive use cases to use when running an application.
	// +kubebuilder:validation:Optional
	InteractiveConfiguration []InteractiveConfigurationParameters `json:"interactiveConfiguration,omitempty" tf:"interactive_configuration,omitempty"`

	// –  The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	// +kubebuilder:validation:Optional
	MaximumCapacity []MaximumCapacityParameters `json:"maximumCapacity,omitempty" tf:"maximum_capacity,omitempty"`

	// –  The name of the application.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// –  The network configuration for customer VPC connectivity.
	// +kubebuilder:validation:Optional
	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// –  The EMR release version associated with the application.
	// +kubebuilder:validation:Optional
	ReleaseLabel *string `json:"releaseLabel,omitempty" tf:"release_label,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// –  The type of application you want to start, such as spark or hive.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AutoStartConfigurationInitParameters struct {

	// Enables the application to automatically start on job submission. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AutoStartConfigurationObservation struct {

	// Enables the application to automatically start on job submission. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AutoStartConfigurationParameters struct {

	// Enables the application to automatically start on job submission. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AutoStopConfigurationInitParameters struct {

	// Enables the application to automatically start on job submission. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The amount of idle time in minutes after which your application will automatically stop. Defaults to 15 minutes.
	IdleTimeoutMinutes *float64 `json:"idleTimeoutMinutes,omitempty" tf:"idle_timeout_minutes,omitempty"`
}

type AutoStopConfigurationObservation struct {

	// Enables the application to automatically start on job submission. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The amount of idle time in minutes after which your application will automatically stop. Defaults to 15 minutes.
	IdleTimeoutMinutes *float64 `json:"idleTimeoutMinutes,omitempty" tf:"idle_timeout_minutes,omitempty"`
}

type AutoStopConfigurationParameters struct {

	// Enables the application to automatically start on job submission. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The amount of idle time in minutes after which your application will automatically stop. Defaults to 15 minutes.
	// +kubebuilder:validation:Optional
	IdleTimeoutMinutes *float64 `json:"idleTimeoutMinutes,omitempty" tf:"idle_timeout_minutes,omitempty"`
}

type ImageConfigurationInitParameters struct {

	// The image URI.
	ImageURI *string `json:"imageUri,omitempty" tf:"image_uri,omitempty"`
}

type ImageConfigurationObservation struct {

	// The image URI.
	ImageURI *string `json:"imageUri,omitempty" tf:"image_uri,omitempty"`
}

type ImageConfigurationParameters struct {

	// The image URI.
	// +kubebuilder:validation:Optional
	ImageURI *string `json:"imageUri" tf:"image_uri,omitempty"`
}

type InitialCapacityConfigInitParameters struct {

	// The resource configuration of the initial capacity configuration.
	WorkerConfiguration []WorkerConfigurationInitParameters `json:"workerConfiguration,omitempty" tf:"worker_configuration,omitempty"`

	// The number of workers in the initial capacity configuration.
	WorkerCount *float64 `json:"workerCount,omitempty" tf:"worker_count,omitempty"`
}

type InitialCapacityConfigObservation struct {

	// The resource configuration of the initial capacity configuration.
	WorkerConfiguration []WorkerConfigurationObservation `json:"workerConfiguration,omitempty" tf:"worker_configuration,omitempty"`

	// The number of workers in the initial capacity configuration.
	WorkerCount *float64 `json:"workerCount,omitempty" tf:"worker_count,omitempty"`
}

type InitialCapacityConfigParameters struct {

	// The resource configuration of the initial capacity configuration.
	// +kubebuilder:validation:Optional
	WorkerConfiguration []WorkerConfigurationParameters `json:"workerConfiguration,omitempty" tf:"worker_configuration,omitempty"`

	// The number of workers in the initial capacity configuration.
	// +kubebuilder:validation:Optional
	WorkerCount *float64 `json:"workerCount" tf:"worker_count,omitempty"`
}

type InitialCapacityInitParameters struct {

	// The initial capacity configuration per worker.
	InitialCapacityConfig []InitialCapacityConfigInitParameters `json:"initialCapacityConfig,omitempty" tf:"initial_capacity_config,omitempty"`

	// The worker type for an analytics framework. For Spark applications, the key can either be set to Driver or Executor. For Hive applications, it can be set to HiveDriver or TezTask.
	InitialCapacityType *string `json:"initialCapacityType,omitempty" tf:"initial_capacity_type,omitempty"`
}

type InitialCapacityObservation struct {

	// The initial capacity configuration per worker.
	InitialCapacityConfig []InitialCapacityConfigObservation `json:"initialCapacityConfig,omitempty" tf:"initial_capacity_config,omitempty"`

	// The worker type for an analytics framework. For Spark applications, the key can either be set to Driver or Executor. For Hive applications, it can be set to HiveDriver or TezTask.
	InitialCapacityType *string `json:"initialCapacityType,omitempty" tf:"initial_capacity_type,omitempty"`
}

type InitialCapacityParameters struct {

	// The initial capacity configuration per worker.
	// +kubebuilder:validation:Optional
	InitialCapacityConfig []InitialCapacityConfigParameters `json:"initialCapacityConfig,omitempty" tf:"initial_capacity_config,omitempty"`

	// The worker type for an analytics framework. For Spark applications, the key can either be set to Driver or Executor. For Hive applications, it can be set to HiveDriver or TezTask.
	// +kubebuilder:validation:Optional
	InitialCapacityType *string `json:"initialCapacityType" tf:"initial_capacity_type,omitempty"`
}

type InteractiveConfigurationInitParameters struct {

	// Enables an Apache Livy endpoint that you can connect to and run interactive jobs.
	LivyEndpointEnabled *bool `json:"livyEndpointEnabled,omitempty" tf:"livy_endpoint_enabled,omitempty"`

	// Enables you to connect an application to Amazon EMR Studio to run interactive workloads in a notebook.
	StudioEnabled *bool `json:"studioEnabled,omitempty" tf:"studio_enabled,omitempty"`
}

type InteractiveConfigurationObservation struct {

	// Enables an Apache Livy endpoint that you can connect to and run interactive jobs.
	LivyEndpointEnabled *bool `json:"livyEndpointEnabled,omitempty" tf:"livy_endpoint_enabled,omitempty"`

	// Enables you to connect an application to Amazon EMR Studio to run interactive workloads in a notebook.
	StudioEnabled *bool `json:"studioEnabled,omitempty" tf:"studio_enabled,omitempty"`
}

type InteractiveConfigurationParameters struct {

	// Enables an Apache Livy endpoint that you can connect to and run interactive jobs.
	// +kubebuilder:validation:Optional
	LivyEndpointEnabled *bool `json:"livyEndpointEnabled,omitempty" tf:"livy_endpoint_enabled,omitempty"`

	// Enables you to connect an application to Amazon EMR Studio to run interactive workloads in a notebook.
	// +kubebuilder:validation:Optional
	StudioEnabled *bool `json:"studioEnabled,omitempty" tf:"studio_enabled,omitempty"`
}

type MaximumCapacityInitParameters struct {

	// The maximum allowed CPU for an application.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// The maximum allowed disk for an application.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// The maximum allowed resources for an application.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type MaximumCapacityObservation struct {

	// The maximum allowed CPU for an application.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// The maximum allowed disk for an application.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// The maximum allowed resources for an application.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type MaximumCapacityParameters struct {

	// The maximum allowed CPU for an application.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu" tf:"cpu,omitempty"`

	// The maximum allowed disk for an application.
	// +kubebuilder:validation:Optional
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// The maximum allowed resources for an application.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory" tf:"memory,omitempty"`
}

type NetworkConfigurationInitParameters struct {

	// The array of security group Ids for customer VPC connectivity.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The array of subnet Ids for customer VPC connectivity.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type NetworkConfigurationObservation struct {

	// The array of security group Ids for customer VPC connectivity.
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The array of subnet Ids for customer VPC connectivity.
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type NetworkConfigurationParameters struct {

	// The array of security group Ids for customer VPC connectivity.
	// +kubebuilder:validation:Optional
	// +listType=set
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The array of subnet Ids for customer VPC connectivity.
	// +kubebuilder:validation:Optional
	// +listType=set
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`
}

type WorkerConfigurationInitParameters struct {

	// The maximum allowed CPU for an application.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// The maximum allowed disk for an application.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// The maximum allowed resources for an application.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type WorkerConfigurationObservation struct {

	// The maximum allowed CPU for an application.
	CPU *string `json:"cpu,omitempty" tf:"cpu,omitempty"`

	// The maximum allowed disk for an application.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// The maximum allowed resources for an application.
	Memory *string `json:"memory,omitempty" tf:"memory,omitempty"`
}

type WorkerConfigurationParameters struct {

	// The maximum allowed CPU for an application.
	// +kubebuilder:validation:Optional
	CPU *string `json:"cpu" tf:"cpu,omitempty"`

	// The maximum allowed disk for an application.
	// +kubebuilder:validation:Optional
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// The maximum allowed resources for an application.
	// +kubebuilder:validation:Optional
	Memory *string `json:"memory" tf:"memory,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Application is the Schema for the Applications API. Manages an EMR Serverless Application
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.releaseLabel) || (has(self.initProvider) && has(self.initProvider.releaseLabel))",message="spec.forProvider.releaseLabel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
