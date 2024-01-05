// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFeatureGates implements FeatureGateInterface
type FakeFeatureGates struct {
	Fake *FakeConfigV1
}

var featuregatesResource = v1.SchemeGroupVersion.WithResource("featuregates")

var featuregatesKind = v1.SchemeGroupVersion.WithKind("FeatureGate")

// Get takes name of the featureGate, and returns the corresponding featureGate object, and an error if there is any.
func (c *FakeFeatureGates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(featuregatesResource, name), &v1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.FeatureGate), err
}

// List takes label and field selectors, and returns the list of FeatureGates that match those selectors.
func (c *FakeFeatureGates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.FeatureGateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(featuregatesResource, featuregatesKind, opts), &v1.FeatureGateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.FeatureGateList{ListMeta: obj.(*v1.FeatureGateList).ListMeta}
	for _, item := range obj.(*v1.FeatureGateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested featureGates.
func (c *FakeFeatureGates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(featuregatesResource, opts))
}

// Create takes the representation of a featureGate and creates it.  Returns the server's representation of the featureGate, and an error, if there is any.
func (c *FakeFeatureGates) Create(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.CreateOptions) (result *v1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(featuregatesResource, featureGate), &v1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.FeatureGate), err
}

// Update takes the representation of a featureGate and updates it. Returns the server's representation of the featureGate, and an error, if there is any.
func (c *FakeFeatureGates) Update(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.UpdateOptions) (result *v1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(featuregatesResource, featureGate), &v1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.FeatureGate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFeatureGates) UpdateStatus(ctx context.Context, featureGate *v1.FeatureGate, opts metav1.UpdateOptions) (*v1.FeatureGate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(featuregatesResource, "status", featureGate), &v1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.FeatureGate), err
}

// Delete takes name of the featureGate and deletes it. Returns an error if one occurs.
func (c *FakeFeatureGates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(featuregatesResource, name, opts), &v1.FeatureGate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFeatureGates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(featuregatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.FeatureGateList{})
	return err
}

// Patch applies the patch and returns the patched featureGate.
func (c *FakeFeatureGates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FeatureGate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(featuregatesResource, name, pt, data, subresources...), &v1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.FeatureGate), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied featureGate.
func (c *FakeFeatureGates) Apply(ctx context.Context, featureGate *configv1.FeatureGateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.FeatureGate, err error) {
	if featureGate == nil {
		return nil, fmt.Errorf("featureGate provided to Apply must not be nil")
	}
	data, err := json.Marshal(featureGate)
	if err != nil {
		return nil, err
	}
	name := featureGate.Name
	if name == nil {
		return nil, fmt.Errorf("featureGate.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(featuregatesResource, *name, types.ApplyPatchType, data), &v1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.FeatureGate), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFeatureGates) ApplyStatus(ctx context.Context, featureGate *configv1.FeatureGateApplyConfiguration, opts metav1.ApplyOptions) (result *v1.FeatureGate, err error) {
	if featureGate == nil {
		return nil, fmt.Errorf("featureGate provided to Apply must not be nil")
	}
	data, err := json.Marshal(featureGate)
	if err != nil {
		return nil, err
	}
	name := featureGate.Name
	if name == nil {
		return nil, fmt.Errorf("featureGate.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(featuregatesResource, *name, types.ApplyPatchType, data, "status"), &v1.FeatureGate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.FeatureGate), err
}