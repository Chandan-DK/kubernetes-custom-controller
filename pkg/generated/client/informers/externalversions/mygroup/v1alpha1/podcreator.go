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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	mygroupv1alpha1 "github.com/Chandan-DK/kubernetes-custom-controller/api/mygroup/v1alpha1"
	versioned "github.com/Chandan-DK/kubernetes-custom-controller/pkg/generated/client/clientset/versioned"
	internalinterfaces "github.com/Chandan-DK/kubernetes-custom-controller/pkg/generated/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Chandan-DK/kubernetes-custom-controller/pkg/generated/client/listers/mygroup/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodCreatorInformer provides access to a shared informer and lister for
// PodCreators.
type PodCreatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodCreatorLister
}

type podCreatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodCreatorInformer constructs a new informer for PodCreator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodCreatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodCreatorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodCreatorInformer constructs a new informer for PodCreator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodCreatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MygroupV1alpha1().PodCreators(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MygroupV1alpha1().PodCreators(namespace).Watch(context.TODO(), options)
			},
		},
		&mygroupv1alpha1.PodCreator{},
		resyncPeriod,
		indexers,
	)
}

func (f *podCreatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodCreatorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podCreatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mygroupv1alpha1.PodCreator{}, f.defaultInformer)
}

func (f *podCreatorInformer) Lister() v1alpha1.PodCreatorLister {
	return v1alpha1.NewPodCreatorLister(f.Informer().GetIndexer())
}