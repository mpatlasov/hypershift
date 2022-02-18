package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InboundSecurityRuleClient is the network Client
type InboundSecurityRuleClient struct {
	BaseClient
}

// NewInboundSecurityRuleClient creates an instance of the InboundSecurityRuleClient client.
func NewInboundSecurityRuleClient(subscriptionID string) InboundSecurityRuleClient {
	return NewInboundSecurityRuleClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInboundSecurityRuleClientWithBaseURI creates an instance of the InboundSecurityRuleClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewInboundSecurityRuleClientWithBaseURI(baseURI string, subscriptionID string) InboundSecurityRuleClient {
	return InboundSecurityRuleClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the specified Network Virtual Appliance Inbound Security Rules.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkVirtualApplianceName - the name of the Network Virtual Appliance.
// ruleCollectionName - the name of security rule collection.
// parameters - parameters supplied to the create or update Network Virtual Appliance Inbound Security Rules
// operation.
func (client InboundSecurityRuleClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule) (result InboundSecurityRuleCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InboundSecurityRuleClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkVirtualApplianceName, ruleCollectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InboundSecurityRuleClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InboundSecurityRuleClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InboundSecurityRuleClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, ruleCollectionName string, parameters InboundSecurityRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkVirtualApplianceName": autorest.Encode("path", networkVirtualApplianceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"ruleCollectionName":          autorest.Encode("path", ruleCollectionName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/inboundSecurityRules/{ruleCollectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InboundSecurityRuleClient) CreateOrUpdateSender(req *http.Request) (future InboundSecurityRuleCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client InboundSecurityRuleClient) CreateOrUpdateResponder(resp *http.Response) (result InboundSecurityRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
