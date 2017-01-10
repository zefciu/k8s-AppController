/*
Copyright 2016 The Kubernetes Authors.

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
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/runtime"
	versionedwatch "k8s.io/client-go/pkg/watch/versioned"
)

// GroupName is the group name use in this package
const GroupName = "apps"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1alpha1"}
var OriginalGroupVersion = unversioned.GroupVersion{Group: "alphaapps", Version: "v1alpha1"}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes, addDefaultingFuncs, addConversionFuncs)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&PetSet{},
		&PetSetList{},
		&v1.ListOptions{},
		&v1.DeleteOptions{},
	)
	scheme.AddKnownTypes(OriginalGroupVersion,
		&PetSet{},
		&PetSetList{},
		&v1.ListOptions{},
		&v1.DeleteOptions{})
	versionedwatch.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func (obj *PetSet) GetObjectKind() unversioned.ObjectKind     { return &obj.TypeMeta }
func (obj *PetSetList) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
