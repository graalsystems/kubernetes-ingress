// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	configurationv1 "github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1"
	scheme "github.com/nginxinc/kubernetes-ingress/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// GlobalConfigurationsGetter has a method to return a GlobalConfigurationInterface.
// A group's client should implement this interface.
type GlobalConfigurationsGetter interface {
	GlobalConfigurations(namespace string) GlobalConfigurationInterface
}

// GlobalConfigurationInterface has methods to work with GlobalConfiguration resources.
type GlobalConfigurationInterface interface {
	Create(ctx context.Context, globalConfiguration *configurationv1.GlobalConfiguration, opts metav1.CreateOptions) (*configurationv1.GlobalConfiguration, error)
	Update(ctx context.Context, globalConfiguration *configurationv1.GlobalConfiguration, opts metav1.UpdateOptions) (*configurationv1.GlobalConfiguration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*configurationv1.GlobalConfiguration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*configurationv1.GlobalConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *configurationv1.GlobalConfiguration, err error)
	GlobalConfigurationExpansion
}

// globalConfigurations implements GlobalConfigurationInterface
type globalConfigurations struct {
	*gentype.ClientWithList[*configurationv1.GlobalConfiguration, *configurationv1.GlobalConfigurationList]
}

// newGlobalConfigurations returns a GlobalConfigurations
func newGlobalConfigurations(c *K8sV1Client, namespace string) *globalConfigurations {
	return &globalConfigurations{
		gentype.NewClientWithList[*configurationv1.GlobalConfiguration, *configurationv1.GlobalConfigurationList](
			"globalconfigurations",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *configurationv1.GlobalConfiguration { return &configurationv1.GlobalConfiguration{} },
			func() *configurationv1.GlobalConfigurationList { return &configurationv1.GlobalConfigurationList{} },
		),
	}
}
