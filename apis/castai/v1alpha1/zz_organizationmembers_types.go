/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OrganizationMembersObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of email addresses corresponding to users who should be given member access to the organization.
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// CAST AI organization ID.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// A list of email addresses corresponding to users who should be given owner access to the organization.
	Owners []*string `json:"owners,omitempty" tf:"owners,omitempty"`

	// A list of email addresses corresponding to users who should be given viewer access to the organization.
	Viewers []*string `json:"viewers,omitempty" tf:"viewers,omitempty"`
}

type OrganizationMembersParameters struct {

	// A list of email addresses corresponding to users who should be given member access to the organization.
	// +kubebuilder:validation:Optional
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// CAST AI organization ID.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// A list of email addresses corresponding to users who should be given owner access to the organization.
	// +kubebuilder:validation:Optional
	Owners []*string `json:"owners,omitempty" tf:"owners,omitempty"`

	// A list of email addresses corresponding to users who should be given viewer access to the organization.
	// +kubebuilder:validation:Optional
	Viewers []*string `json:"viewers,omitempty" tf:"viewers,omitempty"`
}

// OrganizationMembersSpec defines the desired state of OrganizationMembers
type OrganizationMembersSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationMembersParameters `json:"forProvider"`
}

// OrganizationMembersStatus defines the observed state of OrganizationMembers.
type OrganizationMembersStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationMembersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationMembers is the Schema for the OrganizationMemberss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type OrganizationMembers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.organizationId)",message="organizationId is a required parameter"
	Spec   OrganizationMembersSpec   `json:"spec"`
	Status OrganizationMembersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationMembersList contains a list of OrganizationMemberss
type OrganizationMembersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationMembers `json:"items"`
}

// Repository type metadata.
var (
	OrganizationMembers_Kind             = "OrganizationMembers"
	OrganizationMembers_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationMembers_Kind}.String()
	OrganizationMembers_KindAPIVersion   = OrganizationMembers_Kind + "." + CRDGroupVersion.String()
	OrganizationMembers_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationMembers_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationMembers{}, &OrganizationMembersList{})
}