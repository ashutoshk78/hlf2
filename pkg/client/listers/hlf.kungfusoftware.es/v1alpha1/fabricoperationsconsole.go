// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.pkuidlabs.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricOperationsConsoleLister helps list FabricOperationsConsoles.
// All objects returned here must be treated as read-only.
type FabricOperationsConsoleLister interface {
	// List lists all FabricOperationsConsoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOperationsConsole, err error)
	// FabricOperationsConsoles returns an object that can list and get FabricOperationsConsoles.
	FabricOperationsConsoles(namespace string) FabricOperationsConsoleNamespaceLister
	FabricOperationsConsoleListerExpansion
}

// fabricOperationsConsoleLister implements the FabricOperationsConsoleLister interface.
type fabricOperationsConsoleLister struct {
	indexer cache.Indexer
}

// NewFabricOperationsConsoleLister returns a new FabricOperationsConsoleLister.
func NewFabricOperationsConsoleLister(indexer cache.Indexer) FabricOperationsConsoleLister {
	return &fabricOperationsConsoleLister{indexer: indexer}
}

// List lists all FabricOperationsConsoles in the indexer.
func (s *fabricOperationsConsoleLister) List(selector labels.Selector) (ret []*v1alpha1.FabricOperationsConsole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricOperationsConsole))
	})
	return ret, err
}

// FabricOperationsConsoles returns an object that can list and get FabricOperationsConsoles.
func (s *fabricOperationsConsoleLister) FabricOperationsConsoles(namespace string) FabricOperationsConsoleNamespaceLister {
	return fabricOperationsConsoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricOperationsConsoleNamespaceLister helps list and get FabricOperationsConsoles.
// All objects returned here must be treated as read-only.
type FabricOperationsConsoleNamespaceLister interface {
	// List lists all FabricOperationsConsoles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOperationsConsole, err error)
	// Get retrieves the FabricOperationsConsole from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricOperationsConsole, error)
	FabricOperationsConsoleNamespaceListerExpansion
}

// fabricOperationsConsoleNamespaceLister implements the FabricOperationsConsoleNamespaceLister
// interface.
type fabricOperationsConsoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricOperationsConsoles in the indexer for a given namespace.
func (s fabricOperationsConsoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricOperationsConsole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricOperationsConsole))
	})
	return ret, err
}

// Get retrieves the FabricOperationsConsole from the indexer for a given namespace and name.
func (s fabricOperationsConsoleNamespaceLister) Get(name string) (*v1alpha1.FabricOperationsConsole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricoperationsconsole"), name)
	}
	return obj.(*v1alpha1.FabricOperationsConsole), nil
}
