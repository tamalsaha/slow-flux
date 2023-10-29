/*
Copyright AppsCode Inc. and Contributors

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

package versioned

import (
	"fmt"
	"net/http"

	autoscalingv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/autoscaling/v1alpha1"
	catalogv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/catalog/v1alpha1"
	configv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/config/v1alpha1"
	dashboardv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/dashboard/v1alpha1"
	kubedbv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1"
	kubedbv1alpha2 "kubedb.dev/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha2"
	opsv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/ops/v1alpha1"
	postgresv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/postgres/v1alpha1"
	schemav1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/schema/v1alpha1"
	uiv1alpha1 "kubedb.dev/apimachinery/client/clientset/versioned/typed/ui/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface
	CatalogV1alpha1() catalogv1alpha1.CatalogV1alpha1Interface
	ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface
	DashboardV1alpha1() dashboardv1alpha1.DashboardV1alpha1Interface
	KubedbV1alpha1() kubedbv1alpha1.KubedbV1alpha1Interface
	KubedbV1alpha2() kubedbv1alpha2.KubedbV1alpha2Interface
	OpsV1alpha1() opsv1alpha1.OpsV1alpha1Interface
	PostgresV1alpha1() postgresv1alpha1.PostgresV1alpha1Interface
	SchemaV1alpha1() schemav1alpha1.SchemaV1alpha1Interface
	UiV1alpha1() uiv1alpha1.UiV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	autoscalingV1alpha1 *autoscalingv1alpha1.AutoscalingV1alpha1Client
	catalogV1alpha1     *catalogv1alpha1.CatalogV1alpha1Client
	configV1alpha1      *configv1alpha1.ConfigV1alpha1Client
	dashboardV1alpha1   *dashboardv1alpha1.DashboardV1alpha1Client
	kubedbV1alpha1      *kubedbv1alpha1.KubedbV1alpha1Client
	kubedbV1alpha2      *kubedbv1alpha2.KubedbV1alpha2Client
	opsV1alpha1         *opsv1alpha1.OpsV1alpha1Client
	postgresV1alpha1    *postgresv1alpha1.PostgresV1alpha1Client
	schemaV1alpha1      *schemav1alpha1.SchemaV1alpha1Client
	uiV1alpha1          *uiv1alpha1.UiV1alpha1Client
}

// AutoscalingV1alpha1 retrieves the AutoscalingV1alpha1Client
func (c *Clientset) AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface {
	return c.autoscalingV1alpha1
}

// CatalogV1alpha1 retrieves the CatalogV1alpha1Client
func (c *Clientset) CatalogV1alpha1() catalogv1alpha1.CatalogV1alpha1Interface {
	return c.catalogV1alpha1
}

// ConfigV1alpha1 retrieves the ConfigV1alpha1Client
func (c *Clientset) ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface {
	return c.configV1alpha1
}

// DashboardV1alpha1 retrieves the DashboardV1alpha1Client
func (c *Clientset) DashboardV1alpha1() dashboardv1alpha1.DashboardV1alpha1Interface {
	return c.dashboardV1alpha1
}

// KubedbV1alpha1 retrieves the KubedbV1alpha1Client
func (c *Clientset) KubedbV1alpha1() kubedbv1alpha1.KubedbV1alpha1Interface {
	return c.kubedbV1alpha1
}

// KubedbV1alpha2 retrieves the KubedbV1alpha2Client
func (c *Clientset) KubedbV1alpha2() kubedbv1alpha2.KubedbV1alpha2Interface {
	return c.kubedbV1alpha2
}

// OpsV1alpha1 retrieves the OpsV1alpha1Client
func (c *Clientset) OpsV1alpha1() opsv1alpha1.OpsV1alpha1Interface {
	return c.opsV1alpha1
}

// PostgresV1alpha1 retrieves the PostgresV1alpha1Client
func (c *Clientset) PostgresV1alpha1() postgresv1alpha1.PostgresV1alpha1Interface {
	return c.postgresV1alpha1
}

// SchemaV1alpha1 retrieves the SchemaV1alpha1Client
func (c *Clientset) SchemaV1alpha1() schemav1alpha1.SchemaV1alpha1Interface {
	return c.schemaV1alpha1
}

// UiV1alpha1 retrieves the UiV1alpha1Client
func (c *Clientset) UiV1alpha1() uiv1alpha1.UiV1alpha1Interface {
	return c.uiV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.autoscalingV1alpha1, err = autoscalingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.catalogV1alpha1, err = catalogv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.configV1alpha1, err = configv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.dashboardV1alpha1, err = dashboardv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kubedbV1alpha1, err = kubedbv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kubedbV1alpha2, err = kubedbv1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.opsV1alpha1, err = opsv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.postgresV1alpha1, err = postgresv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.schemaV1alpha1, err = schemav1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.uiV1alpha1, err = uiv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.autoscalingV1alpha1 = autoscalingv1alpha1.New(c)
	cs.catalogV1alpha1 = catalogv1alpha1.New(c)
	cs.configV1alpha1 = configv1alpha1.New(c)
	cs.dashboardV1alpha1 = dashboardv1alpha1.New(c)
	cs.kubedbV1alpha1 = kubedbv1alpha1.New(c)
	cs.kubedbV1alpha2 = kubedbv1alpha2.New(c)
	cs.opsV1alpha1 = opsv1alpha1.New(c)
	cs.postgresV1alpha1 = postgresv1alpha1.New(c)
	cs.schemaV1alpha1 = schemav1alpha1.New(c)
	cs.uiV1alpha1 = uiv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}