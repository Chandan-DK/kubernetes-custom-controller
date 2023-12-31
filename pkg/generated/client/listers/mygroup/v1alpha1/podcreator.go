/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Chandan-DK/kubernetes-custom-controller/api/mygroup/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodCreatorLister helps list PodCreators.
// All objects returned here must be treated as read-only.
type PodCreatorLister interface {
	// List lists all PodCreators in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PodCreator, err error)
	// PodCreators returns an object that can list and get PodCreators.
	PodCreators(namespace string) PodCreatorNamespaceLister
	PodCreatorListerExpansion
}

// podCreatorLister implements the PodCreatorLister interface.
type podCreatorLister struct {
	indexer cache.Indexer
}

// NewPodCreatorLister returns a new PodCreatorLister.
func NewPodCreatorLister(indexer cache.Indexer) PodCreatorLister {
	return &podCreatorLister{indexer: indexer}
}

// List lists all PodCreators in the indexer.
func (s *podCreatorLister) List(selector labels.Selector) (ret []*v1alpha1.PodCreator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodCreator))
	})
	return ret, err
}

// PodCreators returns an object that can list and get PodCreators.
func (s *podCreatorLister) PodCreators(namespace string) PodCreatorNamespaceLister {
	return podCreatorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodCreatorNamespaceLister helps list and get PodCreators.
// All objects returned here must be treated as read-only.
type PodCreatorNamespaceLister interface {
	// List lists all PodCreators in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PodCreator, err error)
	// Get retrieves the PodCreator from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PodCreator, error)
	PodCreatorNamespaceListerExpansion
}

// podCreatorNamespaceLister implements the PodCreatorNamespaceLister
// interface.
type podCreatorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodCreators in the indexer for a given namespace.
func (s podCreatorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PodCreator, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodCreator))
	})
	return ret, err
}

// Get retrieves the PodCreator from the indexer for a given namespace and name.
func (s podCreatorNamespaceLister) Get(name string) (*v1alpha1.PodCreator, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podcreator"), name)
	}
	return obj.(*v1alpha1.PodCreator), nil
}
