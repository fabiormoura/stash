package framework

import (
	"github.com/appscode/go/crypto/rand"
	scs "github.com/appscode/stash/client/clientset"
	clientset "k8s.io/client-go/kubernetes"
)

type Framework struct {
	kubeClient  clientset.Interface
	stashClient scs.ExtensionInterface
	namespace   string
}

func New(kubeClient clientset.Interface, extClient scs.ExtensionInterface) *Framework {
	return &Framework{
		kubeClient:  kubeClient,
		stashClient: extClient,
		namespace:   rand.WithUniqSuffix("test-stash"),
	}
}

func (f *Framework) Invoke() *Invocation {
	return &Invocation{
		Framework: f,
		app:       rand.WithUniqSuffix("stash-e2e"),
	}
}

type Invocation struct {
	*Framework
	app string
}
