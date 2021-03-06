// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	garden "github.com/gardener/gardener/pkg/apis/garden"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackupInfrastructures implements BackupInfrastructureInterface
type FakeBackupInfrastructures struct {
	Fake *FakeGarden
	ns   string
}

var backupinfrastructuresResource = schema.GroupVersionResource{Group: "garden.sapcloud.io", Version: "", Resource: "backupinfrastructures"}

var backupinfrastructuresKind = schema.GroupVersionKind{Group: "garden.sapcloud.io", Version: "", Kind: "BackupInfrastructure"}

// Get takes name of the backupInfrastructure, and returns the corresponding backupInfrastructure object, and an error if there is any.
func (c *FakeBackupInfrastructures) Get(name string, options v1.GetOptions) (result *garden.BackupInfrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backupinfrastructuresResource, c.ns, name), &garden.BackupInfrastructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*garden.BackupInfrastructure), err
}

// List takes label and field selectors, and returns the list of BackupInfrastructures that match those selectors.
func (c *FakeBackupInfrastructures) List(opts v1.ListOptions) (result *garden.BackupInfrastructureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backupinfrastructuresResource, backupinfrastructuresKind, c.ns, opts), &garden.BackupInfrastructureList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &garden.BackupInfrastructureList{ListMeta: obj.(*garden.BackupInfrastructureList).ListMeta}
	for _, item := range obj.(*garden.BackupInfrastructureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupInfrastructures.
func (c *FakeBackupInfrastructures) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backupinfrastructuresResource, c.ns, opts))

}

// Create takes the representation of a backupInfrastructure and creates it.  Returns the server's representation of the backupInfrastructure, and an error, if there is any.
func (c *FakeBackupInfrastructures) Create(backupInfrastructure *garden.BackupInfrastructure) (result *garden.BackupInfrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backupinfrastructuresResource, c.ns, backupInfrastructure), &garden.BackupInfrastructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*garden.BackupInfrastructure), err
}

// Update takes the representation of a backupInfrastructure and updates it. Returns the server's representation of the backupInfrastructure, and an error, if there is any.
func (c *FakeBackupInfrastructures) Update(backupInfrastructure *garden.BackupInfrastructure) (result *garden.BackupInfrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backupinfrastructuresResource, c.ns, backupInfrastructure), &garden.BackupInfrastructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*garden.BackupInfrastructure), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupInfrastructures) UpdateStatus(backupInfrastructure *garden.BackupInfrastructure) (*garden.BackupInfrastructure, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backupinfrastructuresResource, "status", c.ns, backupInfrastructure), &garden.BackupInfrastructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*garden.BackupInfrastructure), err
}

// Delete takes name of the backupInfrastructure and deletes it. Returns an error if one occurs.
func (c *FakeBackupInfrastructures) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backupinfrastructuresResource, c.ns, name), &garden.BackupInfrastructure{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupInfrastructures) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backupinfrastructuresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &garden.BackupInfrastructureList{})
	return err
}

// Patch applies the patch and returns the patched backupInfrastructure.
func (c *FakeBackupInfrastructures) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.BackupInfrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backupinfrastructuresResource, c.ns, name, pt, data, subresources...), &garden.BackupInfrastructure{})

	if obj == nil {
		return nil, err
	}
	return obj.(*garden.BackupInfrastructure), err
}
