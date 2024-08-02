# \LongWeekendAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LongWeekendLongWeekend**](LongWeekendAPI.md#LongWeekendLongWeekend) | **Get** /api/v3/LongWeekend/{year}/{countryCode} | Get long weekends for a given country



## LongWeekendLongWeekend

> []LongWeekendV3Dto LongWeekendLongWeekend(ctx, year, countryCode).Execute()

Get long weekends for a given country

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
	resp, r, err := apiClient.LongWeekendAPI.LongWeekendLongWeekend(context.Background(), year, countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongWeekendAPI.LongWeekendLongWeekend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LongWeekendLongWeekend`: []LongWeekendV3Dto
	fmt.Fprintf(os.Stdout, "Response from `LongWeekendAPI.LongWeekendLongWeekend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**year** | **int32** |  | 
**countryCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLongWeekendLongWeekendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]LongWeekendV3Dto**](LongWeekendV3Dto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

