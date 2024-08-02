/*
Nager.Date API - V3

Nager.Date is open source software. If you would like to support the project you can award a GitHub star ⭐ or much better <a href='https://github.com/sponsors/nager'>start a sponsorship</a>

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dateapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PublicHolidayAPIService PublicHolidayAPI service
type PublicHolidayAPIService service

type ApiPublicHolidayIsTodayPublicHolidayRequest struct {
	ctx context.Context
	ApiService *PublicHolidayAPIService
	countryCode string
	countyCode *string
	offset *int32
}

// Federal State Code
func (r ApiPublicHolidayIsTodayPublicHolidayRequest) CountyCode(countyCode string) ApiPublicHolidayIsTodayPublicHolidayRequest {
	r.countyCode = &countyCode
	return r
}

// utc timezone offset
func (r ApiPublicHolidayIsTodayPublicHolidayRequest) Offset(offset int32) ApiPublicHolidayIsTodayPublicHolidayRequest {
	r.offset = &offset
	return r
}

func (r ApiPublicHolidayIsTodayPublicHolidayRequest) Execute() (*http.Response, error) {
	return r.ApiService.PublicHolidayIsTodayPublicHolidayExecute(r)
}

/*
PublicHolidayIsTodayPublicHoliday Is today a public holiday

The calculation is made on the basis of UTC time to adjust the time please use the offset.

This is a special endpoint for `curl`


200 = Today is a public holiday

204 = Today is not a public holiday


`STATUSCODE=$(curl --silent --output /dev/stderr --write-out "%{http_code}" https://date.nager.at/Api/v3/IsTodayPublicHoliday/AT)`


`if [ $STATUSCODE -ne 200 ]; then # error handling; fi`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param countryCode The Country Code
 @return ApiPublicHolidayIsTodayPublicHolidayRequest
*/
func (a *PublicHolidayAPIService) PublicHolidayIsTodayPublicHoliday(ctx context.Context, countryCode string) ApiPublicHolidayIsTodayPublicHolidayRequest {
	return ApiPublicHolidayIsTodayPublicHolidayRequest{
		ApiService: a,
		ctx: ctx,
		countryCode: countryCode,
	}
}

// Execute executes the request
func (a *PublicHolidayAPIService) PublicHolidayIsTodayPublicHolidayExecute(r ApiPublicHolidayIsTodayPublicHolidayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicHolidayAPIService.PublicHolidayIsTodayPublicHoliday")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/IsTodayPublicHoliday/{countryCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"countryCode"+"}", url.PathEscape(parameterValueToString(r.countryCode, "countryCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.countyCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "countyCode", r.countyCode, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPublicHolidayNextPublicHolidaysRequest struct {
	ctx context.Context
	ApiService *PublicHolidayAPIService
	countryCode string
}

func (r ApiPublicHolidayNextPublicHolidaysRequest) Execute() ([]PublicHolidayV3Dto, *http.Response, error) {
	return r.ApiService.PublicHolidayNextPublicHolidaysExecute(r)
}

/*
PublicHolidayNextPublicHolidays Returns the upcoming public holidays for the next 365 days for the given country

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param countryCode
 @return ApiPublicHolidayNextPublicHolidaysRequest
*/
func (a *PublicHolidayAPIService) PublicHolidayNextPublicHolidays(ctx context.Context, countryCode string) ApiPublicHolidayNextPublicHolidaysRequest {
	return ApiPublicHolidayNextPublicHolidaysRequest{
		ApiService: a,
		ctx: ctx,
		countryCode: countryCode,
	}
}

// Execute executes the request
//  @return []PublicHolidayV3Dto
func (a *PublicHolidayAPIService) PublicHolidayNextPublicHolidaysExecute(r ApiPublicHolidayNextPublicHolidaysRequest) ([]PublicHolidayV3Dto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PublicHolidayV3Dto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicHolidayAPIService.PublicHolidayNextPublicHolidays")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/NextPublicHolidays/{countryCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"countryCode"+"}", url.PathEscape(parameterValueToString(r.countryCode, "countryCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPublicHolidayNextPublicHolidaysWorldwideRequest struct {
	ctx context.Context
	ApiService *PublicHolidayAPIService
}

func (r ApiPublicHolidayNextPublicHolidaysWorldwideRequest) Execute() ([]PublicHolidayV3Dto, *http.Response, error) {
	return r.ApiService.PublicHolidayNextPublicHolidaysWorldwideExecute(r)
}

/*
PublicHolidayNextPublicHolidaysWorldwide Returns the upcoming public holidays for the next 7 days

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPublicHolidayNextPublicHolidaysWorldwideRequest
*/
func (a *PublicHolidayAPIService) PublicHolidayNextPublicHolidaysWorldwide(ctx context.Context) ApiPublicHolidayNextPublicHolidaysWorldwideRequest {
	return ApiPublicHolidayNextPublicHolidaysWorldwideRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PublicHolidayV3Dto
func (a *PublicHolidayAPIService) PublicHolidayNextPublicHolidaysWorldwideExecute(r ApiPublicHolidayNextPublicHolidaysWorldwideRequest) ([]PublicHolidayV3Dto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PublicHolidayV3Dto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicHolidayAPIService.PublicHolidayNextPublicHolidaysWorldwide")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/NextPublicHolidaysWorldwide"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPublicHolidayPublicHolidaysV3Request struct {
	ctx context.Context
	ApiService *PublicHolidayAPIService
	year int32
	countryCode string
}

func (r ApiPublicHolidayPublicHolidaysV3Request) Execute() ([]PublicHolidayV3Dto, *http.Response, error) {
	return r.ApiService.PublicHolidayPublicHolidaysV3Execute(r)
}

/*
PublicHolidayPublicHolidaysV3 Get public holidays

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param year 
 @param countryCode 
 @return ApiPublicHolidayPublicHolidaysV3Request
*/
func (a *PublicHolidayAPIService) PublicHolidayPublicHolidaysV3(ctx context.Context, year int32, countryCode string) ApiPublicHolidayPublicHolidaysV3Request {
	return ApiPublicHolidayPublicHolidaysV3Request{
		ApiService: a,
		ctx: ctx,
		year: year,
		countryCode: countryCode,
	}
}

// Execute executes the request
//  @return []PublicHolidayV3Dto
func (a *PublicHolidayAPIService) PublicHolidayPublicHolidaysV3Execute(r ApiPublicHolidayPublicHolidaysV3Request) ([]PublicHolidayV3Dto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PublicHolidayV3Dto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicHolidayAPIService.PublicHolidayPublicHolidaysV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v3/PublicHolidays/{year}/{countryCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"year"+"}", url.PathEscape(parameterValueToString(r.year, "year")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"countryCode"+"}", url.PathEscape(parameterValueToString(r.countryCode, "countryCode")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
