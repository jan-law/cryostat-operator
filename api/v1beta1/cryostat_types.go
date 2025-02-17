// Copyright The Cryostat Authors
//
// The Universal Permissive License (UPL), Version 1.0
//
// Subject to the condition set forth below, permission is hereby granted to any
// person obtaining a copy of this software, associated documentation and/or data
// (collectively the "Software"), free of charge and under any and all copyright
// rights in the Software, and any and all patent rights owned or freely
// licensable by each licensor hereunder covering either (i) the unmodified
// Software as contributed to or provided by such licensor, or (ii) the Larger
// Works (as defined below), to deal in both
//
// (a) the Software, and
// (b) any piece of software and/or hardware listed in the lrgrwrks.txt file if
// one is included with the Software (each a "Larger Work" to which the Software
// is contributed by such licensors),
//
// without restriction, including without limitation the rights to copy, create
// derivative works of, display, perform, and distribute the Software and make,
// use, sell, offer for sale, import, export, have made, and have sold the
// Software and the Larger Work(s), and to sublicense the foregoing rights on
// either these or other terms.
//
// This license is subject to the following condition:
// The above copyright notice and either this complete permission notice or at
// a minimum a reference to the UPL must be included in all copies or
// substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CryostatSpec defines the desired state of Cryostat
type CryostatSpec struct {
	// Deploy a pared-down Cryostat instance with no Grafana Dashboard or JFR Data Source.
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Minimal Deployment",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Minimal bool `json:"minimal"`
	// List of TLS certificates to trust when connecting to targets
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Trusted TLS Certificates"
	TrustedCertSecrets []CertificateSecret `json:"trustedCertSecrets,omitempty"`
	// List of Flight Recorder Event Templates to preconfigure in Cryostat
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Event Templates"
	EventTemplates []TemplateConfigMap `json:"eventTemplates,omitempty"`
	// Use cert-manager to secure in-cluster communication between Cryostat components.
	// Requires cert-manager to be installed.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Enable cert-manager Integration",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	EnableCertManager *bool `json:"enableCertManager"`
	// Options to customize the storage for Flight Recordings and Templates
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	StorageOptions *StorageConfiguration `json:"storageOptions,omitempty"`
	// Options to customize the services created for the Cryostat application and Grafana dashboard.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ServiceOptions *ServiceConfigList `json:"serviceOptions,omitempty"`
	// Options to control how the operator exposes the application outside of the cluster
	// using an Ingress or Route.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	NetworkOptions *NetworkConfigurationList `json:"networkOptions,omitempty"`
	// Options to configure Cryostat Automated Report Analysis
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ReportOptions *ReportConfiguration `json:"reportOptions,omitempty"`
	// The maximum number of WebSocket client connections allowed (minimum 1, default unlimited)
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Max WebSocket Connections",xDescriptors={"urn:alm:descriptor:com.tectonic.ui:number"}
	// +kubebuilder:validation:Minimum=1
	MaxWsConnections int32 `json:"maxWsConnections,omitempty"`
	// Options to customize the JMX target connections cache for the Cryostat application
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="JMX Connections Cache Options"
	JmxCacheOptions *JmxCacheOptions `json:"jmxCacheOptions,omitempty"`
	// Resource requirements for the Cryostat deployment
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Resources ResourceConfigList `json:"resources,omitempty"`
}

type ResourceConfigList struct {
	// Resource requirements for the Cryostat application. If specifying a memory limit, at least 768MiB is recommended.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	CoreResources corev1.ResourceRequirements `json:"coreResources,omitempty"`
	// Resource requirements for the JFR Data Source container
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	DataSourceResources corev1.ResourceRequirements `json:"dataSourceResources,omitempty"`
	// Resource requirements for the Grafana container
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	GrafanaResources corev1.ResourceRequirements `json:"grafanaResources,omitempty"`
}

// CryostatStatus defines the observed state of Cryostat
type CryostatStatus struct {
	// Conditions of the components managed by the Cryostat Operator
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=status,displayName="Cryostat Conditions",xDescriptors={"urn:alm:descriptor:io.kubernetes.conditions"}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Address of the deployed Cryostat web application
	// +operator-sdk:csv:customresourcedefinitions:type=status,xDescriptors={"urn:alm:descriptor:org.w3:link"}
	ApplicationURL string `json:"applicationUrl"`
}

// CryostatConditionType refers to a Condition type that may be used in status.conditions
type CryostatConditionType string

const (
	// Whether the main Cryostat deployment is available
	ConditionTypeMainDeploymentAvailable CryostatConditionType = "MainDeploymentAvailable"
	// Whether the main Cryostat deployment is progressing
	ConditionTypeMainDeploymentProgressing CryostatConditionType = "MainDeploymentProgressing"
	// If pods within the main Cryostat deployment failed to be created or destroyed
	ConditionTypeMainDeploymentReplicaFailure CryostatConditionType = "MainDeploymentReplicaFailure"
	// If enabled, whether the reports deployment is available
	ConditionTypeReportsDeploymentAvailable CryostatConditionType = "ReportsDeploymentAvailable"
	// If enabled, whether the reports deployment is progressing
	ConditionTypeReportsDeploymentProgressing CryostatConditionType = "ReportsDeploymentProgressing"
	// If enabled, whether pods in the reports deployment failed to be created or destroyed
	ConditionTypeReportsDeploymentReplicaFailure CryostatConditionType = "ReportsDeploymentReplicaFailure"
	// If enabled, whether TLS setup is complete for the Cryostat components
	ConditionTypeTLSSetupComplete CryostatConditionType = "TLSSetupComplete"
)

// StorageConfiguration provides customization to the storage created by
// the operator to hold Flight Recordings and Recording Templates. If no
// configurations are specified, a PVC will be created by default.
type StorageConfiguration struct {
	// Configuration for the Persistent Volume Claim to be created
	// by the operator.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	PVC *PersistentVolumeClaimConfig `json:"pvc,omitempty"`

	// Configuration for an EmptyDir to be created
	// by the operator instead of a PVC.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EmptyDir *EmptyDirConfig `json:"emptyDir,omitempty"`
}

// ReportConfiguration is used to determine how many replicas of cryostat-reports
// the operator should create and what the resource limits of those containers
// should be. If no replicas are created then Cryostat is configured to use basic
// subprocess report generation. If at least one replica is created then Cryostat
// is configured to use remote report generation, pointed at a load balancer service
// in front of the cryostat-reports replicas.
type ReportConfiguration struct {
	// The number of report sidecar replica containers to deploy.
	// Each replica can service one report generation request at a time.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:podCount"}
	Replicas int32 `json:"replicas,omitempty"`
	// The resources allocated to each sidecar replica.
	// A replica with more resources can handle larger input recordings and will process them faster.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// When zero report sidecar replicas are requested, SubProcessMaxHeapSize configures
	// the maximum heap size of the basic subprocess report generator in MiB.
	// The default heap size is `200` (MiB).
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:number"}
	SubProcessMaxHeapSize int32 `json:"subProcessMaxHeapSize,omitempty"`
}

// ServiceConfig provides customization for a service created
// by the operator.
type ServiceConfig struct {
	// Type of service to create. Defaults to "ClusterIP".
	// +optional
	ServiceType *corev1.ServiceType `json:"serviceType,omitempty"`
	// Annotations to add to the service during its creation.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// Labels to add to the service during its creation.
	// The labels with keys "app" and "component" are reserved
	// for use by the operator.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// CoreServiceConfig provides customization for the service handling
// traffic for the Cryostat application
type CoreServiceConfig struct {
	// HTTP port number for the Cryostat application service.
	// Defaults to 8181.
	// +optional
	HTTPPort *int32 `json:"httpPort,omitempty"`
	// Remote JMX port number for the Cryostat application service.
	// Defaults to 9091.
	// +optional
	JMXPort       *int32 `json:"jmxPort,omitempty"`
	ServiceConfig `json:",inline"`
}

// GrafanaServiceConfig provides customization for the service handling
// traffic for the Grafana dashboard
type GrafanaServiceConfig struct {
	// HTTP port number for the Grafana dashboard service.
	// Defaults to 3000.
	// +optional
	HTTPPort      *int32 `json:"httpPort,omitempty"`
	ServiceConfig `json:",inline"`
}

// ReportsServiceConfig provides customization for the service handling
// traffic for the cryostat-reports sidecars
type ReportsServiceConfig struct {
	// HTTP port number for the cryostat-reports service.
	// Defaults to 10000.
	// +optional
	HTTPPort      *int32 `json:"httpPort,omitempty"`
	ServiceConfig `json:",inline"`
}

// ServiceConfigList holds the service configuration for each
// service created by the operator.
type ServiceConfigList struct {
	// Specification for the service responsible for the Cryostat application.
	// +optional
	CoreConfig *CoreServiceConfig `json:"coreConfig,omitempty"`
	// Specification for the service responsible for the Cryostat Grafana dashboard.
	// +optional
	GrafanaConfig *GrafanaServiceConfig `json:"grafanaConfig,omitempty"`
	// Specification for the service responsible for the cryostat-reports sidecars.
	// +optional
	ReportsConfig *ReportsServiceConfig `json:"reportsConfig,omitempty"`
}

// NetworkConfiguration provides customization for the corresponding ingress,
// which allows a service to be exposed when running in a Kubernetes environment
type NetworkConfiguration struct {
	// Configuration for an ingress object.
	// Currently subpaths are not supported, so unique hosts must be specified
	// (if a single external IP is being used) to differentiate between ingresses/services
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	IngressSpec *netv1.IngressSpec `json:"ingressSpec,omitempty"`
	// Annotations to add to the ingress during its creation.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Annotations map[string]string `json:"annotations,omitempty"`
	// Labels to add to the ingress during its creation.
	// The label with key "app" is reserved for use by the operator.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Labels map[string]string `json:"labels,omitempty"`
}

// NetworkConfigurationList holds all three NetworkConfiguration objects that specify
// the Ingress configurations for the services created by the operator for
// the main Cryostat deployment
type NetworkConfigurationList struct {
	// Specifications for ingress that exposes the cryostat service
	// (which serves the cryostat web-client)
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	CoreConfig *NetworkConfiguration `json:"coreConfig,omitempty"`
	// Specifications for ingress that exposes the cryostat-command service
	// (which serves the websocket command channel)
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:hidden"}
	//
	// Deprecated: CommandConfig is no longer used.
	CommandConfig *NetworkConfiguration `json:"commandConfig,omitempty"`
	// Specifications for ingress that exposes the cryostat-grafana service
	// (which serves the grafana dashboard)
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	GrafanaConfig *NetworkConfiguration `json:"grafanaConfig,omitempty"`
}

// PersistentVolumeClaimConfig holds all customization options to
// configure a Persistent Volume Claim to be created and managed
// by the operator.
type PersistentVolumeClaimConfig struct {
	// Annotations to add to the Persistent Volume Claim during its creation.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// Labels to add to the Persistent Volume Claim during its creation.
	// The label with key "app" is reserved for use by the operator.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// Spec for a Persistent Volume Claim, whose options will override the
	// defaults used by the operator. Unless overriden, the PVC will be
	// created with the default Storage Class and 500MiB of storage.
	// Once the operator has created the PVC, changes to this field have
	// no effect.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Spec *corev1.PersistentVolumeClaimSpec `json:"spec,omitempty"`
}

// EmptyDirConfig holds all customization options to
// configure an EmptyDir to be created and managed
// by the operator.
type EmptyDirConfig struct {
	// When enabled, Cryostat will use EmptyDir volumes instead of a Persistent Volume Claim. Any PVC configurations will be ignored.
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	Enabled bool `json:"enabled,omitempty"`
	// Unless specified, the emptyDir volume will be mounted on
	// the same storage medium backing the node. Setting this field to
	// "Memory" will mount the emptyDir on a tmpfs (RAM-backed filesystem).
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldDependency:storageOptions.emptyDir.enabled:true"}
	Medium corev1.StorageMedium `json:"medium,omitempty"`
	// The maximum memory limit for the emptyDir. Default is unbounded.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:fieldDependency:storageOptions.emptyDir.enabled:true"}
	// +kubebuilder:validation:Pattern=^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
	SizeLimit string `json:"sizeLimit,omitempty"`
}

// JmxCacheConfig provides customization for the JMX target connections
// cache for the Cryostat application
type JmxCacheOptions struct {
	// The maximum number of JMX connections to cache. Use `-1` for an unlimited cache size (TTL expiration only). Defaults to `-1`.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:number"}
	// +kubebuilder:validation:Minimum=-1
	TargetCacheSize int32 `json:"targetCacheSize,omitempty"`
	// The time to live (in seconds) for cached JMX connections. Defaults to `10`.
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:number"}
	// +kubebuilder:validation:Minimum=1
	TargetCacheTTL int32 `json:"targetCacheTTL,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=cryostats,scope=Namespaced

// Cryostat contains configuration options for controlling the Deployment of
// the Cryostat application and its related components. A Cryostat instance
// must be created to instruct the operator to deploy the Cryostat application.
//+operator-sdk:csv:customresourcedefinitions:resources={{Deployment,v1},{Ingress,v1},{PersistentVolumeClaim,v1},{Secret,v1},{Service,v1},{Route,v1},{ConsoleLink,v1}}
// +kubebuilder:printcolumn:name="Application URL",type=string,JSONPath=`.status.applicationUrl`
type Cryostat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CryostatSpec   `json:"spec,omitempty"`
	Status CryostatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CryostatList contains a list of Cryostat
type CryostatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cryostat `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cryostat{}, &CryostatList{})
}

// DefaultCertificateKey will be used when looking up the certificate within a secret,
// if a key is not manually specified
const DefaultCertificateKey = corev1.TLSCertKey

type CertificateSecret struct {
	// Name of secret in the local namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:io.kubernetes:Secret"}
	SecretName string `json:"secretName"`
	// Key within secret containing the certificate
	// +optional
	CertificateKey *string `json:"certificateKey,omitempty"`
}

// A ConfigMap containing a .jfc template file
type TemplateConfigMap struct {
	// Name of config map in the local namespace
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:io.kubernetes:ConfigMap"}
	ConfigMapName string `json:"configMapName"`
	// Filename within config map containing the template file
	Filename string `json:"filename"`
}
