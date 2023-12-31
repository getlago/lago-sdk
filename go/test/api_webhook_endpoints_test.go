/*
Lago API documentation

Testing WebhookEndpointsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package lagoapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/getlago/sdk/go"
)

func Test_lagoapi_WebhookEndpointsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebhookEndpointsAPIService CreateWebhookEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WebhookEndpointsAPI.CreateWebhookEndpoint(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookEndpointsAPIService DestroyWebhookEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lagoId string

		resp, httpRes, err := apiClient.WebhookEndpointsAPI.DestroyWebhookEndpoint(context.Background(), lagoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookEndpointsAPIService FindAllWebhookEndpoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WebhookEndpointsAPI.FindAllWebhookEndpoints(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookEndpointsAPIService FindWebhookEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lagoId string

		resp, httpRes, err := apiClient.WebhookEndpointsAPI.FindWebhookEndpoint(context.Background(), lagoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebhookEndpointsAPIService UpdateWebhookEndpoint", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lagoId string

		resp, httpRes, err := apiClient.WebhookEndpointsAPI.UpdateWebhookEndpoint(context.Background(), lagoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
