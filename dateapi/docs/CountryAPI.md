# \CountryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountryAvailableCountries**](CountryAPI.md#CountryAvailableCountries) | **Get** /api/v3/AvailableCountries | Get all available countries
[**CountryCountryInfo**](CountryAPI.md#CountryCountryInfo) | **Get** /api/v3/CountryInfo/{countryCode} | Get country info for the given country



## CountryAvailableCountries

> []CountryV3Dto CountryAvailableCountries(ctx).Execute()

Get all available countries

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountryAPI.CountryAvailableCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.CountryAvailableCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountryAvailableCountries`: []CountryV3Dto
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.CountryAvailableCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCountryAvailableCountriesRequest struct via the builder pattern


### Return type

[**[]CountryV3Dto**](CountryV3Dto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountryCountryInfo

> CountryInfoDto CountryCountryInfo(ctx, countryCode).Execute()

Get country info for the given country

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	countryCode := "countryCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountryAPI.CountryCountryInfo(context.Background(), countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryAPI.CountryCountryInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountryCountryInfo`: CountryInfoDto
	fmt.Fprintf(os.Stdout, "Response from `CountryAPI.CountryCountryInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountryCountryInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountryInfoDto**](CountryInfoDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

