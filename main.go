package main

import (
	"context"
	"fmt"

	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2/klogr"
	"sigs.k8s.io/cli-utils/pkg/flowcontrol"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// https://github.com/fluxcd/pkg/blob/3e343b975578c194de040afaa5e811a3d9f8b715/runtime/client/client.go#L75-L82

// GetConfigOrDie wraps ctrl.GetConfigOrDie and checks if the Kubernetes apiserver
// has PriorityAndFairness flow control filter enabled. If true, it returns a rest.Config
// with client side throttling disabled. Otherwise, it returns a modified rest.Config
// configured with the provided Options.
func GetConfigOrDie() *rest.Config {
	config := ctrl.GetConfigOrDie()
	enabled, err := flowcontrol.IsEnabled(context.Background(), config)
	if err == nil && enabled {
		// A negative QPS and Burst indicates that the client should not have a rate limiter.
		// Ref: https://github.com/kubernetes/kubernetes/blob/v1.24.0/staging/src/k8s.io/client-go/rest/config.go#L354-L364
		config.QPS = -1
		config.Burst = -1
		return config
	}
	config.QPS = 100
	config.Burst = 100
	return config
}

func NewClient() (client.Client, error) {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)

	ctrl.SetLogger(klogr.New())
	cfg := GetConfigOrDie()

	hc, err := rest.HTTPClientFor(cfg)
	if err != nil {
		return nil, err
	}

	mapper, err := apiutil.NewDynamicRESTMapper(cfg, hc)
	if err != nil {
		return nil, err
	}

	return client.New(cfg, client.Options{
		Scheme: scheme,
		Mapper: mapper,
		//Opts: client.WarningHandlerOptions{
		//	SuppressWarnings:   false,
		//	AllowDuplicateLogs: false,
		//},
	})
}

func main() {
	if err := useKubebuilderClient(); err != nil {
		panic(err)
	}
}

func useKubebuilderClient() error {
	fmt.Println("Using kubebuilder client")
	kc, err := NewClient()
	if err != nil {
		return err
	}

	var pglist core.NodeList
	err = kc.List(context.TODO(), &pglist)
	if err != nil {
		return err
	}
	for _, db := range pglist.Items {
		fmt.Println(client.ObjectKeyFromObject(&db))
	}
	return nil
}
