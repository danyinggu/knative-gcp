// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"github.com/google/knative-gcp/pkg/pubsub/adapter"
	"github.com/google/knative-gcp/pkg/pubsub/adapter/converters"
	"github.com/google/knative-gcp/pkg/utils/clients"
)

// Injectors from wire.go:

func InitializeAdapter(ctx context.Context, maxConnsPerHost clients.MaxConnsPerHost, projectID clients.ProjectID, subscriptionID adapter.SubscriptionID, namespace adapter.Namespace, name adapter.Name, resourceGroup adapter.ResourceGroup, args *adapter.AdapterArgs) (*adapter.Adapter, error) {
	client, err := clients.NewPubsubClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	subscription := adapter.NewPubSubSubscription(ctx, client, subscriptionID)
	httpClient := clients.NewHTTPClient(ctx, maxConnsPerHost)
	converter := converters.NewPubSubConverter()
	statsReporter, err := adapter.NewStatsReporter(name, namespace, resourceGroup)
	if err != nil {
		return nil, err
	}
	adapterAdapter := adapter.NewAdapter(ctx, projectID, namespace, name, resourceGroup, subscription, httpClient, converter, statsReporter, args)
	return adapterAdapter, nil
}
