/*


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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/openshift/hypershift/api/types/certificates/v1alpha1"
	certificatesv1alpha1 "github.com/openshift/hypershift/client/applyconfiguration/certificates/v1alpha1"
	scheme "github.com/openshift/hypershift/client/clientset/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CertificateSigningRequestApprovalsGetter has a method to return a CertificateSigningRequestApprovalInterface.
// A group's client should implement this interface.
type CertificateSigningRequestApprovalsGetter interface {
	CertificateSigningRequestApprovals(namespace string) CertificateSigningRequestApprovalInterface
}

// CertificateSigningRequestApprovalInterface has methods to work with CertificateSigningRequestApproval resources.
type CertificateSigningRequestApprovalInterface interface {
	Create(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.CreateOptions) (*v1alpha1.CertificateSigningRequestApproval, error)
	Update(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.UpdateOptions) (*v1alpha1.CertificateSigningRequestApproval, error)
	UpdateStatus(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.UpdateOptions) (*v1alpha1.CertificateSigningRequestApproval, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CertificateSigningRequestApproval, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CertificateSigningRequestApprovalList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateSigningRequestApproval, err error)
	Apply(ctx context.Context, certificateSigningRequestApproval *certificatesv1alpha1.CertificateSigningRequestApprovalApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error)
	ApplyStatus(ctx context.Context, certificateSigningRequestApproval *certificatesv1alpha1.CertificateSigningRequestApprovalApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error)
	CertificateSigningRequestApprovalExpansion
}

// certificateSigningRequestApprovals implements CertificateSigningRequestApprovalInterface
type certificateSigningRequestApprovals struct {
	client rest.Interface
	ns     string
}

// newCertificateSigningRequestApprovals returns a CertificateSigningRequestApprovals
func newCertificateSigningRequestApprovals(c *CertificatesV1alpha1Client, namespace string) *certificateSigningRequestApprovals {
	return &certificateSigningRequestApprovals{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the certificateSigningRequestApproval, and returns the corresponding certificateSigningRequestApproval object, and an error if there is any.
func (c *certificateSigningRequestApprovals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	result = &v1alpha1.CertificateSigningRequestApproval{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CertificateSigningRequestApprovals that match those selectors.
func (c *certificateSigningRequestApprovals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificateSigningRequestApprovalList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CertificateSigningRequestApprovalList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested certificateSigningRequestApprovals.
func (c *certificateSigningRequestApprovals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a certificateSigningRequestApproval and creates it.  Returns the server's representation of the certificateSigningRequestApproval, and an error, if there is any.
func (c *certificateSigningRequestApprovals) Create(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.CreateOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	result = &v1alpha1.CertificateSigningRequestApproval{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequestApproval).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a certificateSigningRequestApproval and updates it. Returns the server's representation of the certificateSigningRequestApproval, and an error, if there is any.
func (c *certificateSigningRequestApprovals) Update(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.UpdateOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	result = &v1alpha1.CertificateSigningRequestApproval{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		Name(certificateSigningRequestApproval.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequestApproval).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *certificateSigningRequestApprovals) UpdateStatus(ctx context.Context, certificateSigningRequestApproval *v1alpha1.CertificateSigningRequestApproval, opts v1.UpdateOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	result = &v1alpha1.CertificateSigningRequestApproval{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		Name(certificateSigningRequestApproval.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequestApproval).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the certificateSigningRequestApproval and deletes it. Returns an error if one occurs.
func (c *certificateSigningRequestApprovals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *certificateSigningRequestApprovals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched certificateSigningRequestApproval.
func (c *certificateSigningRequestApprovals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	result = &v1alpha1.CertificateSigningRequestApproval{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied certificateSigningRequestApproval.
func (c *certificateSigningRequestApprovals) Apply(ctx context.Context, certificateSigningRequestApproval *certificatesv1alpha1.CertificateSigningRequestApprovalApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	if certificateSigningRequestApproval == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(certificateSigningRequestApproval)
	if err != nil {
		return nil, err
	}
	name := certificateSigningRequestApproval.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval.Name must be provided to Apply")
	}
	result = &v1alpha1.CertificateSigningRequestApproval{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *certificateSigningRequestApprovals) ApplyStatus(ctx context.Context, certificateSigningRequestApproval *certificatesv1alpha1.CertificateSigningRequestApprovalApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.CertificateSigningRequestApproval, err error) {
	if certificateSigningRequestApproval == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(certificateSigningRequestApproval)
	if err != nil {
		return nil, err
	}

	name := certificateSigningRequestApproval.Name
	if name == nil {
		return nil, fmt.Errorf("certificateSigningRequestApproval.Name must be provided to Apply")
	}

	result = &v1alpha1.CertificateSigningRequestApproval{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("certificatesigningrequestapprovals").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
