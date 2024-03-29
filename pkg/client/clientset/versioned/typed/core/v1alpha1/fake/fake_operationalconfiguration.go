/* Copyright Microsoft */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/microsoft/scylla/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperationalConfigurations implements OperationalConfigurationInterface
type FakeOperationalConfigurations struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var operationalconfigurationsResource = schema.GroupVersionResource{Group: "core.hydra.io", Version: "v1alpha1", Resource: "operationalconfigurations"}

var operationalconfigurationsKind = schema.GroupVersionKind{Group: "core.hydra.io", Version: "v1alpha1", Kind: "OperationalConfiguration"}

// Get takes name of the operationalConfiguration, and returns the corresponding operationalConfiguration object, and an error if there is any.
func (c *FakeOperationalConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.OperationalConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(operationalconfigurationsResource, c.ns, name), &v1alpha1.OperationalConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperationalConfiguration), err
}

// List takes label and field selectors, and returns the list of OperationalConfigurations that match those selectors.
func (c *FakeOperationalConfigurations) List(opts v1.ListOptions) (result *v1alpha1.OperationalConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(operationalconfigurationsResource, operationalconfigurationsKind, c.ns, opts), &v1alpha1.OperationalConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OperationalConfigurationList{ListMeta: obj.(*v1alpha1.OperationalConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.OperationalConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operationalConfigurations.
func (c *FakeOperationalConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(operationalconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a operationalConfiguration and creates it.  Returns the server's representation of the operationalConfiguration, and an error, if there is any.
func (c *FakeOperationalConfigurations) Create(operationalConfiguration *v1alpha1.OperationalConfiguration) (result *v1alpha1.OperationalConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(operationalconfigurationsResource, c.ns, operationalConfiguration), &v1alpha1.OperationalConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperationalConfiguration), err
}

// Update takes the representation of a operationalConfiguration and updates it. Returns the server's representation of the operationalConfiguration, and an error, if there is any.
func (c *FakeOperationalConfigurations) Update(operationalConfiguration *v1alpha1.OperationalConfiguration) (result *v1alpha1.OperationalConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(operationalconfigurationsResource, c.ns, operationalConfiguration), &v1alpha1.OperationalConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperationalConfiguration), err
}

// Delete takes name of the operationalConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeOperationalConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(operationalconfigurationsResource, c.ns, name), &v1alpha1.OperationalConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperationalConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(operationalconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OperationalConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched operationalConfiguration.
func (c *FakeOperationalConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OperationalConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(operationalconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OperationalConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OperationalConfiguration), err
}
