// Copyright 2020 The Amadeus Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kafkaconnectv1alpha1 "github.com/AmadeusITGroup/Kubernetes-Kafka-Connect-Operator/pkg/apis/kafkaconnect/v1alpha1"
	versioned "github.com/AmadeusITGroup/Kubernetes-Kafka-Connect-Operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/AmadeusITGroup/Kubernetes-Kafka-Connect-Operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/AmadeusITGroup/Kubernetes-Kafka-Connect-Operator/pkg/generated/listers/kafkaconnect/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KafkaConnectInformer provides access to a shared informer and lister for
// KafkaConnects.
type KafkaConnectInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KafkaConnectLister
}

type kafkaConnectInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKafkaConnectInformer constructs a new informer for KafkaConnect type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKafkaConnectInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKafkaConnectInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKafkaConnectInformer constructs a new informer for KafkaConnect type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKafkaConnectInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafkaconnectV1alpha1().KafkaConnects(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KafkaconnectV1alpha1().KafkaConnects(namespace).Watch(options)
			},
		},
		&kafkaconnectv1alpha1.KafkaConnect{},
		resyncPeriod,
		indexers,
	)
}

func (f *kafkaConnectInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKafkaConnectInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kafkaConnectInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kafkaconnectv1alpha1.KafkaConnect{}, f.defaultInformer)
}

func (f *kafkaConnectInformer) Lister() v1alpha1.KafkaConnectLister {
	return v1alpha1.NewKafkaConnectLister(f.Informer().GetIndexer())
}