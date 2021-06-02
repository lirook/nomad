/*
 * nomad
 *
 * Documentation of the Nomad v1 API.
 *
 * API version: 1.0.0
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiv1

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// JobsApiService JobsApi service
type JobsApiService service

type ApiGetJobAllocationsRequest struct {
	ctx _context.Context
	ApiService *JobsApiService
	jobName string
	all *int32
	region *string
	stale *string
	prefix *string
	namespace *string
	perPage *int32
	nextToken *string
}

func (r ApiGetJobAllocationsRequest) All(all int32) ApiGetJobAllocationsRequest {
	r.all = &all
	return r
}
func (r ApiGetJobAllocationsRequest) Region(region string) ApiGetJobAllocationsRequest {
	r.region = &region
	return r
}
func (r ApiGetJobAllocationsRequest) Stale(stale string) ApiGetJobAllocationsRequest {
	r.stale = &stale
	return r
}
func (r ApiGetJobAllocationsRequest) Prefix(prefix string) ApiGetJobAllocationsRequest {
	r.prefix = &prefix
	return r
}
func (r ApiGetJobAllocationsRequest) Namespace(namespace string) ApiGetJobAllocationsRequest {
	r.namespace = &namespace
	return r
}
func (r ApiGetJobAllocationsRequest) PerPage(perPage int32) ApiGetJobAllocationsRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetJobAllocationsRequest) NextToken(nextToken string) ApiGetJobAllocationsRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiGetJobAllocationsRequest) Execute() ([]AllocListStub, *_nethttp.Response, error) {
	return r.ApiService.GetJobAllocationsExecute(r)
}

/*
 * GetJobAllocations Gets information about a single job's allocations. See [documentation](https://www.nomadproject.io/api-docs/allocations#list-allocations). 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jobName The job identifier.
 * @return ApiGetJobAllocationsRequest
 */
func (a *JobsApiService) GetJobAllocations(ctx _context.Context, jobName string) ApiGetJobAllocationsRequest {
	return ApiGetJobAllocationsRequest{
		ApiService: a,
		ctx: ctx,
		jobName: jobName,
	}
}

/*
 * Execute executes the request
 * @return []AllocListStub
 */
func (a *JobsApiService) GetJobAllocationsExecute(r ApiGetJobAllocationsRequest) ([]AllocListStub, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AllocListStub
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobsApiService.GetJobAllocations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobName}/allocations"
	localVarPath = strings.Replace(localVarPath, "{"+"jobName"+"}", _neturl.PathEscape(parameterToString(r.jobName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.all != nil {
		localVarQueryParams.Add("all", parameterToString(*r.all, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.stale != nil {
		localVarQueryParams.Add("stale", parameterToString(*r.stale, ""))
	}
	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Nomad-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
