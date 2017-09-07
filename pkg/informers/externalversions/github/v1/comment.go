/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1

import (
	github_v1 "github.com/nikhita/kube-custom-controller/pkg/apis/github/v1"
	client "github.com/nikhita/kube-custom-controller/pkg/client"
	internalinterfaces "github.com/nikhita/kube-custom-controller/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/nikhita/kube-custom-controller/pkg/listers/github/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CommentInformer provides access to a shared informer and lister for
// Comments.
type CommentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CommentLister
}

type commentInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewCommentInformer constructs a new informer for Comment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCommentInformer(client client.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.GithubV1().Comments(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.GithubV1().Comments(namespace).Watch(options)
			},
		},
		&github_v1.Comment{},
		resyncPeriod,
		indexers,
	)
}

func defaultCommentInformer(client client.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewCommentInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *commentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&github_v1.Comment{}, defaultCommentInformer)
}

func (f *commentInformer) Lister() v1.CommentLister {
	return v1.NewCommentLister(f.Informer().GetIndexer())
}
