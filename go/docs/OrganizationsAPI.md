# \OrganizationsAPI

All URIs are relative to *https://api.getlago.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateOrganization**](OrganizationsAPI.md#UpdateOrganization) | **Put** /organizations | Update your organization



## UpdateOrganization

> Organization UpdateOrganization(ctx).OrganizationUpdateInput(organizationUpdateInput).Execute()

Update your organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/getlago/sdk/go"
)

func main() {
    organizationUpdateInput := *openapiclient.NewOrganizationUpdateInput(*openapiclient.NewOrganizationUpdateInputOrganization()) // OrganizationUpdateInput | Update an existing organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsAPI.UpdateOrganization(context.Background()).OrganizationUpdateInput(organizationUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.UpdateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.UpdateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationUpdateInput** | [**OrganizationUpdateInput**](OrganizationUpdateInput.md) | Update an existing organization | 

### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

