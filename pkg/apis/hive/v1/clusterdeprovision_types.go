package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterDeprovisionSpec defines the desired state of ClusterDeprovision
type ClusterDeprovisionSpec struct {
	// InfraID is the identifier generated during installation for a cluster. It is used for tagging/naming resources in cloud providers.
	InfraID string `json:"infraID"`

	// ClusterID is a globally unique identifier for the cluster to deprovision. It will be used if specified.
	ClusterID string `json:"clusterID,omitempty"`

	// Platform contains platform-specific configuration for a ClusterDeprovision
	Platform ClusterDeprovisionPlatform `json:"platform,omitempty"`
}

// ClusterDeprovisionStatus defines the observed state of ClusterDeprovision
type ClusterDeprovisionStatus struct {
	// Completed is true when the uninstall has completed successfully
	Completed bool `json:"completed,omitempty"`
}

// ClusterDeprovisionPlatform contains platform-specific configuration for the
// deprovision
type ClusterDeprovisionPlatform struct {
	// AWS contains AWS-specific deprovision settings
	AWS *AWSClusterDeprovision `json:"aws,omitempty"`
	// Azure contains Azure-specific deprovision settings
	Azure *AzureClusterDeprovision `json:"azure,omitempty"`
	// GCP contains GCP-specific deprovision settings
	GCP *GCPClusterDeprovision `json:"gcp,omitempty"`
	// OpenStack contains OpenStack-specific deprovision settings
	OpenStack *OpenStackClusterDeprovision `json:"openstack,omitempty"`
	// VSphere contains VMWare vSphere-specific deprovision settings
	VSphere *VSphereClusterDeprovision `json:"vsphere,omitempty"`
}

// AWSClusterDeprovision contains AWS-specific configuration for a ClusterDeprovision
type AWSClusterDeprovision struct {
	// Region is the AWS region for this deprovisioning
	Region string `json:"region"`

	// CredentialsSecretRef is the AWS account credentials to use for deprovisioning the cluster
	CredentialsSecretRef *corev1.LocalObjectReference `json:"credentialsSecretRef,omitempty"`
}

// AzureClusterDeprovision contains Azure-specific configuration for a ClusterDeprovision
type AzureClusterDeprovision struct {
	// CredentialsSecretRef is the Azure account credentials to use for deprovisioning the cluster
	CredentialsSecretRef *corev1.LocalObjectReference `json:"credentialsSecretRef,omitempty"`
}

// GCPClusterDeprovision contains GCP-specific configuration for a ClusterDeprovision
type GCPClusterDeprovision struct {
	// Region is the GCP region for this deprovision
	Region string `json:"region"`
	// CredentialsSecretRef is the GCP account credentials to use for deprovisioning the cluster
	CredentialsSecretRef *corev1.LocalObjectReference `json:"credentialsSecretRef,omitempty"`
}

// OpenStackClusterDeprovision contains OpenStack-specific configuration for a ClusterDeprovision
type OpenStackClusterDeprovision struct {
	// Cloud is the secion in the clouds.yaml secret below to use for auth/connectivity.
	Cloud string `json:"cloud"`
	// CredentialsSecretRef is the OpenStack account credentials to use for deprovisioning the cluster
	CredentialsSecretRef *corev1.LocalObjectReference `json:"credentialsSecretRef,omitempty"`
}

// VSphereClusterDeprovision contains VMware vSphere-specific configuration for a ClusterDeprovision
type VSphereClusterDeprovision struct {
	// CredentialsSecretRef is the vSphere account credentials to use for deprovisioning the cluster
	CredentialsSecretRef corev1.LocalObjectReference `json:"credentialsSecretRef"`
	// CertificatesSecretRef refers to a secret that contains the vSphere CA certificates
	// necessary for communicating with the VCenter.
	CertificatesSecretRef corev1.LocalObjectReference `json:"certificatesSecretRef"`
	// VCenter is the vSphere vCenter hostname.
	VCenter string `json:"vCenter"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterDeprovision is the Schema for the clusterdeprovisions API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="InfraID",type="string",JSONPath=".spec.infraID"
// +kubebuilder:printcolumn:name="ClusterID",type="string",JSONPath=".spec.clusterID"
// +kubebuilder:printcolumn:name="Completed",type="boolean",JSONPath=".status.completed"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:path=clusterdeprovisions,shortName=cdr,scope=Namespaced
type ClusterDeprovision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterDeprovisionSpec   `json:"spec,omitempty"`
	Status ClusterDeprovisionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterDeprovisionList contains a list of ClusterDeprovision
type ClusterDeprovisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterDeprovision `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterDeprovision{}, &ClusterDeprovisionList{})
}
