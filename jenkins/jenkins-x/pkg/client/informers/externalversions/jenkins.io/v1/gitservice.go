// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkins_io_v1 "github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx/pkg/client/listers/jenkins.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GitServiceInformer provides access to a shared informer and lister for
// GitServices.
type GitServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.GitServiceLister
}

type gitServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGitServiceInformer constructs a new informer for GitService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGitServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGitServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGitServiceInformer constructs a new informer for GitService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGitServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().GitServices(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().GitServices(namespace).Watch(options)
			},
		},
		&jenkins_io_v1.GitService{},
		resyncPeriod,
		indexers,
	)
}

func (f *gitServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGitServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gitServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkins_io_v1.GitService{}, f.defaultInformer)
}

func (f *gitServiceInformer) Lister() v1.GitServiceLister {
	return v1.NewGitServiceLister(f.Informer().GetIndexer())
}
