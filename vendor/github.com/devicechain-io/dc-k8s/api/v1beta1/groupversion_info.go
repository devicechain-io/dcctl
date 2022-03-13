/*
Copyright © 2022 SiteWhere LLC - All Rights Reserved
Unauthorized copying of this file, via any medium is strictly prohibited.
Proprietary and confidential.
*/

// Package v1beta1 contains API Schema definitions for the core v1beta1 API group
//+kubebuilder:object:generate=true
//+groupName=core.devicechain.io
package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "core.devicechain.io", Version: "v1beta1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
