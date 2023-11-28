/*
Lago API documentation

Testing FeesAPIService

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

func Test_lagoapi_FeesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeesAPIService FindAllFees", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeesAPI.FindAllFees(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeesAPIService FindFee", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lagoId string

		resp, httpRes, err := apiClient.FeesAPI.FindFee(context.Background(), lagoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeesAPIService UpdateFee", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var lagoId string

		resp, httpRes, err := apiClient.FeesAPI.UpdateFee(context.Background(), lagoId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
