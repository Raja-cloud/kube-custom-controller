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

// This file was automatically generated by lister-gen

package internalversion

import (
	github "github.com/nikhita/kube-custom-controller/pkg/apis/github"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CommentLister helps list Comments.
type CommentLister interface {
	// List lists all Comments in the indexer.
	List(selector labels.Selector) (ret []*github.Comment, err error)
	// Comments returns an object that can list and get Comments.
	Comments(namespace string) CommentNamespaceLister
	CommentListerExpansion
}

// commentLister implements the CommentLister interface.
type commentLister struct {
	indexer cache.Indexer
}

// NewCommentLister returns a new CommentLister.
func NewCommentLister(indexer cache.Indexer) CommentLister {
	return &commentLister{indexer: indexer}
}

// List lists all Comments in the indexer.
func (s *commentLister) List(selector labels.Selector) (ret []*github.Comment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*github.Comment))
	})
	return ret, err
}

// Comments returns an object that can list and get Comments.
func (s *commentLister) Comments(namespace string) CommentNamespaceLister {
	return commentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CommentNamespaceLister helps list and get Comments.
type CommentNamespaceLister interface {
	// List lists all Comments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*github.Comment, err error)
	// Get retrieves the Comment from the indexer for a given namespace and name.
	Get(name string) (*github.Comment, error)
	CommentNamespaceListerExpansion
}

// commentNamespaceLister implements the CommentNamespaceLister
// interface.
type commentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Comments in the indexer for a given namespace.
func (s commentNamespaceLister) List(selector labels.Selector) (ret []*github.Comment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*github.Comment))
	})
	return ret, err
}

// Get retrieves the Comment from the indexer for a given namespace and name.
func (s commentNamespaceLister) Get(name string) (*github.Comment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(github.Resource("comment"), name)
	}
	return obj.(*github.Comment), nil
}
