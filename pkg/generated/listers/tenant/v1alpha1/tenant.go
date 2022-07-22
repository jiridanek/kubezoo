/*
Copyright 2022 The KubeZoo Authors.

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
	v1alpha1 "github.com/kubewharf/kubezoo/pkg/apis/tenant/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TenantLister helps list Tenants.
type TenantLister interface {
	// List lists all Tenants in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Tenant, err error)
	// Get retrieves the Tenant from the index for a given name.
	Get(name string) (*v1alpha1.Tenant, error)
	TenantListerExpansion
}

// tenantLister implements the TenantLister interface.
type tenantLister struct {
	indexer cache.Indexer
}

// NewTenantLister returns a new TenantLister.
func NewTenantLister(indexer cache.Indexer) TenantLister {
	return &tenantLister{indexer: indexer}
}

// List lists all Tenants in the indexer.
func (s *tenantLister) List(selector labels.Selector) (ret []*v1alpha1.Tenant, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Tenant))
	})
	return ret, err
}

// Get retrieves the Tenant from the index for a given name.
func (s *tenantLister) Get(name string) (*v1alpha1.Tenant, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tenant"), name)
	}
	return obj.(*v1alpha1.Tenant), nil
}
