/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	_context "context"
	"fmt"
	"github.com/ionos-cloud/sdk-go-bundle/shared"
	"io"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// RequestsApiService RequestsApi service
type RequestsApiService service

type ApiRequestsFindByIdRequest struct {
	ctx             _context.Context
	ApiService      *RequestsApiService
	requestId       string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiRequestsFindByIdRequest) Pretty(pretty bool) ApiRequestsFindByIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiRequestsFindByIdRequest) Depth(depth int32) ApiRequestsFindByIdRequest {
	r.depth = &depth
	return r
}
func (r ApiRequestsFindByIdRequest) XContractNumber(xContractNumber int32) ApiRequestsFindByIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiRequestsFindByIdRequest) Execute() (Request, *shared.APIResponse, error) {
	return r.ApiService.RequestsFindByIdExecute(r)
}

/*
 * RequestsFindById Retrieve requests
 * Retrieve the properties of the specified request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param requestId The unique ID of the request.
 * @return ApiRequestsFindByIdRequest
 */
func (a *RequestsApiService) RequestsFindById(ctx _context.Context, requestId string) ApiRequestsFindByIdRequest {
	return ApiRequestsFindByIdRequest{
		ApiService: a,
		ctx:        ctx,
		requestId:  requestId,
	}
}

/*
 * Execute executes the request
 * @return Request
 */
func (a *RequestsApiService) RequestsFindByIdExecute(r ApiRequestsFindByIdRequest) (Request, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Request
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsApiService.RequestsFindById")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/requests/{requestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", _neturl.PathEscape(parameterValueToString(r.requestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pretty", r.pretty, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
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
	if r.xContractNumber != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Contract-Number", *r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &shared.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "RequestsFindById",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := shared.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)))
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.SetError(err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.SetModel(v)
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := shared.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRequestsGetRequest struct {
	ctx                 _context.Context
	ApiService          *RequestsApiService
	filters             _neturl.Values
	orderBy             *string
	maxResults          *int32
	pretty              *bool
	depth               *int32
	xContractNumber     *int32
	filterStatus        *string
	filterCreatedAfter  *string
	filterCreatedBefore *string
	filterCreatedDate   *string
	filterCreatedBy     *string
	filterEtag          *string
	filterRequestStatus *string
	filterMethod        *string
	filterHeaders       *string
	filterBody          *string
	filterUrl           *string
	offset              *int32
	limit               *int32
}

func (r ApiRequestsGetRequest) Pretty(pretty bool) ApiRequestsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiRequestsGetRequest) Depth(depth int32) ApiRequestsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiRequestsGetRequest) XContractNumber(xContractNumber int32) ApiRequestsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}
func (r ApiRequestsGetRequest) FilterStatus(filterStatus string) ApiRequestsGetRequest {
	r.filterStatus = &filterStatus
	return r
}
func (r ApiRequestsGetRequest) FilterCreatedAfter(filterCreatedAfter string) ApiRequestsGetRequest {
	r.filterCreatedAfter = &filterCreatedAfter
	return r
}
func (r ApiRequestsGetRequest) FilterCreatedBefore(filterCreatedBefore string) ApiRequestsGetRequest {
	r.filterCreatedBefore = &filterCreatedBefore
	return r
}
func (r ApiRequestsGetRequest) FilterCreatedDate(filterCreatedDate string) ApiRequestsGetRequest {
	r.filterCreatedDate = &filterCreatedDate
	return r
}
func (r ApiRequestsGetRequest) FilterCreatedBy(filterCreatedBy string) ApiRequestsGetRequest {
	r.filterCreatedBy = &filterCreatedBy
	return r
}
func (r ApiRequestsGetRequest) FilterEtag(filterEtag string) ApiRequestsGetRequest {
	r.filterEtag = &filterEtag
	return r
}
func (r ApiRequestsGetRequest) FilterRequestStatus(filterRequestStatus string) ApiRequestsGetRequest {
	r.filterRequestStatus = &filterRequestStatus
	return r
}
func (r ApiRequestsGetRequest) FilterMethod(filterMethod string) ApiRequestsGetRequest {
	r.filterMethod = &filterMethod
	return r
}
func (r ApiRequestsGetRequest) FilterHeaders(filterHeaders string) ApiRequestsGetRequest {
	r.filterHeaders = &filterHeaders
	return r
}
func (r ApiRequestsGetRequest) FilterBody(filterBody string) ApiRequestsGetRequest {
	r.filterBody = &filterBody
	return r
}
func (r ApiRequestsGetRequest) FilterUrl(filterUrl string) ApiRequestsGetRequest {
	r.filterUrl = &filterUrl
	return r
}
func (r ApiRequestsGetRequest) Offset(offset int32) ApiRequestsGetRequest {
	r.offset = &offset
	return r
}
func (r ApiRequestsGetRequest) Limit(limit int32) ApiRequestsGetRequest {
	r.limit = &limit
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiRequestsGetRequest) Filter(key string, value string) ApiRequestsGetRequest {
	filterKey := fmt.Sprintf("filter.%s", key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiRequestsGetRequest) OrderBy(orderBy string) ApiRequestsGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiRequestsGetRequest) MaxResults(maxResults int32) ApiRequestsGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiRequestsGetRequest) Execute() (Requests, *shared.APIResponse, error) {
	return r.ApiService.RequestsGetExecute(r)
}

/*
 * RequestsGet List requests
 * List all API requests.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRequestsGetRequest
 */
func (a *RequestsApiService) RequestsGet(ctx _context.Context) ApiRequestsGetRequest {
	return ApiRequestsGetRequest{
		ApiService: a,
		ctx:        ctx,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return Requests
 */
func (a *RequestsApiService) RequestsGetExecute(r ApiRequestsGetRequest) (Requests, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Requests
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsApiService.RequestsGet")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/requests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pretty", r.pretty, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
	}
	if r.filterStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.status", r.filterStatus, "")
	}
	if r.filterCreatedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdAfter", r.filterCreatedAfter, "")
	}
	if r.filterCreatedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdBefore", r.filterCreatedBefore, "")
	}
	if r.filterCreatedDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdDate", r.filterCreatedDate, "")
	}
	if r.filterCreatedBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.createdBy", r.filterCreatedBy, "")
	}
	if r.filterEtag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.etag", r.filterEtag, "")
	}
	if r.filterRequestStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.requestStatus", r.filterRequestStatus, "")
	}
	if r.filterMethod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.method", r.filterMethod, "")
	}
	if r.filterHeaders != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.headers", r.filterHeaders, "")
	}
	if r.filterBody != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.body", r.filterBody, "")
	}
	if r.filterUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.url", r.filterUrl, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("maxResults", parameterToString(*r.maxResults, ""))
	}
	if len(r.filters) > 0 {
		for k, v := range r.filters {
			for _, iv := range v {
				localVarQueryParams.Add(k, iv)
			}
		}
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
	if r.xContractNumber != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Contract-Number", *r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &shared.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "RequestsGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := shared.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)))
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.SetError(err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.SetModel(v)
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := shared.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiRequestsStatusGetRequest struct {
	ctx             _context.Context
	ApiService      *RequestsApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	requestId       string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiRequestsStatusGetRequest) Pretty(pretty bool) ApiRequestsStatusGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiRequestsStatusGetRequest) Depth(depth int32) ApiRequestsStatusGetRequest {
	r.depth = &depth
	return r
}
func (r ApiRequestsStatusGetRequest) XContractNumber(xContractNumber int32) ApiRequestsStatusGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiRequestsStatusGetRequest) Filter(key string, value string) ApiRequestsStatusGetRequest {
	filterKey := fmt.Sprintf("filter.%s", key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiRequestsStatusGetRequest) OrderBy(orderBy string) ApiRequestsStatusGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiRequestsStatusGetRequest) MaxResults(maxResults int32) ApiRequestsStatusGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiRequestsStatusGetRequest) Execute() (RequestStatus, *shared.APIResponse, error) {
	return r.ApiService.RequestsStatusGetExecute(r)
}

/*
 * RequestsStatusGet Retrieve request status
 * Retrieve the status of the specified request.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param requestId The unique ID of the request.
 * @return ApiRequestsStatusGetRequest
 */
func (a *RequestsApiService) RequestsStatusGet(ctx _context.Context, requestId string) ApiRequestsStatusGetRequest {
	return ApiRequestsStatusGetRequest{
		ApiService: a,
		ctx:        ctx,
		requestId:  requestId,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return RequestStatus
 */
func (a *RequestsApiService) RequestsStatusGetExecute(r ApiRequestsStatusGetRequest) (RequestStatus, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RequestStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestsApiService.RequestsStatusGet")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/requests/{requestId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"requestId"+"}", _neturl.PathEscape(parameterValueToString(r.requestId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pretty", r.pretty, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, ""))
	}
	if r.maxResults != nil {
		localVarQueryParams.Add("maxResults", parameterToString(*r.maxResults, ""))
	}
	if len(r.filters) > 0 {
		for k, v := range r.filters {
			for _, iv := range v {
				localVarQueryParams.Add(k, iv)
			}
		}
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
	if r.xContractNumber != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Contract-Number", *r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["TokenAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &shared.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "RequestsStatusGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := shared.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(fmt.Sprintf("%s: %s", localVarHTTPResponse.Status, string(localVarBody)))
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.SetError(err.Error())
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.SetModel(v)
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := shared.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}
