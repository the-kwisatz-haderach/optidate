# \PublicHolidayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicHolidayIsTodayPublicHoliday**](PublicHolidayAPI.md#PublicHolidayIsTodayPublicHoliday) | **Get** /api/v3/IsTodayPublicHoliday/{countryCode} | Is today a public holiday
[**PublicHolidayNextPublicHolidays**](PublicHolidayAPI.md#PublicHolidayNextPublicHolidays) | **Get** /api/v3/NextPublicHolidays/{countryCode} | Returns the upcoming public holidays for the next 365 days for the given country
[**PublicHolidayNextPublicHolidaysWorldwide**](PublicHolidayAPI.md#PublicHolidayNextPublicHolidaysWorldwide) | **Get** /api/v3/NextPublicHolidaysWorldwide | Returns the upcoming public holidays for the next 7 days
[**PublicHolidayPublicHolidaysV3**](PublicHolidayAPI.md#PublicHolidayPublicHolidaysV3) | **Get** /api/v3/PublicHolidays/{year}/{countryCode} | Get public holidays



## PublicHolidayIsTodayPublicHoliday

> PublicHolidayIsTodayPublicHoliday(ctx, countryCode).CountyCode(countyCode).Offset(offset).Execute()

Is today a public holiday



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
	countryCode := "countryCode_example" // string | The Country Code
	countyCode := "countyCode_example" // string | Federal State Code (optional)
	offset := int32(56) // int32 | utc timezone offset (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicHolidayAPI.PublicHolidayIsTodayPublicHoliday(context.Background(), countryCode).CountyCode(countyCode).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicHolidayAPI.PublicHolidayIsTodayPublicHoliday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | The Country Code | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicHolidayIsTodayPublicHolidayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countyCode** | **string** | Federal State Code | 
 **offset** | **int32** | utc timezone offset | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicHolidayNextPublicHolidays

> []PublicHolidayV3Dto PublicHolidayNextPublicHolidays(ctx, countryCode).Execute()

Returns the upcoming public holidays for the next 365 days for the given country

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
	resp, r, err := apiClient.PublicHolidayAPI.PublicHolidayNextPublicHolidays(context.Background(), countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicHolidayAPI.PublicHolidayNextPublicHolidays``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicHolidayNextPublicHolidays`: []PublicHolidayV3Dto
	fmt.Fprintf(os.Stdout, "Response from `PublicHolidayAPI.PublicHolidayNextPublicHolidays`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicHolidayNextPublicHolidaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PublicHolidayV3Dto**](PublicHolidayV3Dto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicHolidayNextPublicHolidaysWorldwide

> []PublicHolidayV3Dto PublicHolidayNextPublicHolidaysWorldwide(ctx).Execute()

Returns the upcoming public holidays for the next 7 days

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
	resp, r, err := apiClient.PublicHolidayAPI.PublicHolidayNextPublicHolidaysWorldwide(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicHolidayAPI.PublicHolidayNextPublicHolidaysWorldwide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicHolidayNextPublicHolidaysWorldwide`: []PublicHolidayV3Dto
	fmt.Fprintf(os.Stdout, "Response from `PublicHolidayAPI.PublicHolidayNextPublicHolidaysWorldwide`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPublicHolidayNextPublicHolidaysWorldwideRequest struct via the builder pattern


### Return type

[**[]PublicHolidayV3Dto**](PublicHolidayV3Dto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublicHolidayPublicHolidaysV3

> []PublicHolidayV3Dto PublicHolidayPublicHolidaysV3(ctx, year, countryCode).Execute()

Get public holidays

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
	year := int32(56) // int32 | 
	countryCode := "countryCode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicHolidayAPI.PublicHolidayPublicHolidaysV3(context.Background(), year, countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicHolidayAPI.PublicHolidayPublicHolidaysV3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublicHolidayPublicHolidaysV3`: []PublicHolidayV3Dto
	fmt.Fprintf(os.Stdout, "Response from `PublicHolidayAPI.PublicHolidayPublicHolidaysV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**year** | **int32** |  | 
**countryCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublicHolidayPublicHolidaysV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PublicHolidayV3Dto**](PublicHolidayV3Dto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

