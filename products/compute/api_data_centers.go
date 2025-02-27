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

// DataCentersApiService DataCentersApi service
type DataCentersApiService service

type ApiDatacentersDeleteRequest struct {
	ctx             _context.Context
	ApiService      *DataCentersApiService
	datacenterId    string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersDeleteRequest) Pretty(pretty bool) ApiDatacentersDeleteRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersDeleteRequest) Depth(depth int32) ApiDatacentersDeleteRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersDeleteRequest) XContractNumber(xContractNumber int32) ApiDatacentersDeleteRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersDeleteRequest) Execute() (*shared.APIResponse, error) {
	return r.ApiService.DatacentersDeleteExecute(r)
}

/*
 * DatacentersDelete Delete data centers
 * Delete the specified data center and all the elements it contains. This method is destructive and should be used carefully.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the data center.
 * @return ApiDatacentersDeleteRequest
 */
func (a *DataCentersApiService) DatacentersDelete(ctx _context.Context, datacenterId string) ApiDatacentersDeleteRequest {
	return ApiDatacentersDeleteRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 */
func (a *DataCentersApiService) DatacentersDeleteExecute(r ApiDatacentersDeleteRequest) (*shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCentersApiService.DatacentersDelete")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterValueToString(r.datacenterId, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, httpRequestTime, err := a.client.callAPI(req)

	localVarAPIResponse := &shared.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "DatacentersDelete",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarAPIResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarAPIResponse, err
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
			return localVarAPIResponse, newErr
		}
		newErr.SetModel(v)
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}

type ApiDatacentersFindByIdRequest struct {
	ctx             _context.Context
	ApiService      *DataCentersApiService
	datacenterId    string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersFindByIdRequest) Pretty(pretty bool) ApiDatacentersFindByIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersFindByIdRequest) Depth(depth int32) ApiDatacentersFindByIdRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersFindByIdRequest) XContractNumber(xContractNumber int32) ApiDatacentersFindByIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersFindByIdRequest) Execute() (Datacenter, *shared.APIResponse, error) {
	return r.ApiService.DatacentersFindByIdExecute(r)
}

/*
 * DatacentersFindById Retrieve data centers
 * Retrieve data centers by resource ID. This value is in the response body when the data center is created, and in the list of the data centers, returned by GET.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the data center.
 * @return ApiDatacentersFindByIdRequest
 */
func (a *DataCentersApiService) DatacentersFindById(ctx _context.Context, datacenterId string) ApiDatacentersFindByIdRequest {
	return ApiDatacentersFindByIdRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCentersApiService) DatacentersFindByIdExecute(r ApiDatacentersFindByIdRequest) (Datacenter, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCentersApiService.DatacentersFindById")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterValueToString(r.datacenterId, "")), -1)

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
		Operation:   "DatacentersFindById",
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

type ApiDatacentersGetRequest struct {
	ctx             _context.Context
	ApiService      *DataCentersApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	pretty          *bool
	depth           *int32
	xContractNumber *int32
	offset          *int32
	limit           *int32
}

func (r ApiDatacentersGetRequest) Pretty(pretty bool) ApiDatacentersGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersGetRequest) Depth(depth int32) ApiDatacentersGetRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersGetRequest) XContractNumber(xContractNumber int32) ApiDatacentersGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}
func (r ApiDatacentersGetRequest) Offset(offset int32) ApiDatacentersGetRequest {
	r.offset = &offset
	return r
}
func (r ApiDatacentersGetRequest) Limit(limit int32) ApiDatacentersGetRequest {
	r.limit = &limit
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiDatacentersGetRequest) Filter(key string, value string) ApiDatacentersGetRequest {
	filterKey := fmt.Sprintf("filter.%s", key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiDatacentersGetRequest) OrderBy(orderBy string) ApiDatacentersGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiDatacentersGetRequest) MaxResults(maxResults int32) ApiDatacentersGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiDatacentersGetRequest) Execute() (Datacenters, *shared.APIResponse, error) {
	return r.ApiService.DatacentersGetExecute(r)
}

/*
 * DatacentersGet List your data centers
 * List the data centers for your account. Default limit is the first 100 items; use pagination query parameters for listing more items.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDatacentersGetRequest
 */
func (a *DataCentersApiService) DatacentersGet(ctx _context.Context) ApiDatacentersGetRequest {
	return ApiDatacentersGetRequest{
		ApiService: a,
		ctx:        ctx,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return Datacenters
 */
func (a *DataCentersApiService) DatacentersGetExecute(r ApiDatacentersGetRequest) (Datacenters, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenters
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCentersApiService.DatacentersGet")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/datacenters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pretty", r.pretty, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
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
		Operation:   "DatacentersGet",
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

type ApiDatacentersPatchRequest struct {
	ctx             _context.Context
	ApiService      *DataCentersApiService
	datacenterId    string
	datacenter      *DatacenterPropertiesPut
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersPatchRequest) Datacenter(datacenter DatacenterPropertiesPut) ApiDatacentersPatchRequest {
	r.datacenter = &datacenter
	return r
}
func (r ApiDatacentersPatchRequest) Pretty(pretty bool) ApiDatacentersPatchRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersPatchRequest) Depth(depth int32) ApiDatacentersPatchRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersPatchRequest) XContractNumber(xContractNumber int32) ApiDatacentersPatchRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersPatchRequest) Execute() (Datacenter, *shared.APIResponse, error) {
	return r.ApiService.DatacentersPatchExecute(r)
}

/*
 * DatacentersPatch Partially modify a Data Center by ID
 * Updates the properties of the specified data center, rename it, or change the description.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the data center.
 * @return ApiDatacentersPatchRequest
 */
func (a *DataCentersApiService) DatacentersPatch(ctx _context.Context, datacenterId string) ApiDatacentersPatchRequest {
	return ApiDatacentersPatchRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCentersApiService) DatacentersPatchExecute(r ApiDatacentersPatchRequest) (Datacenter, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCentersApiService.DatacentersPatch")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterValueToString(r.datacenterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.datacenter == nil {
		return localVarReturnValue, nil, reportError("datacenter is required and must be specified")
	}

	if r.pretty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pretty", r.pretty, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.datacenter
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
		Operation:   "DatacentersPatch",
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

type ApiDatacentersPostRequest struct {
	ctx             _context.Context
	ApiService      *DataCentersApiService
	datacenter      *DatacenterPost
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersPostRequest) Datacenter(datacenter DatacenterPost) ApiDatacentersPostRequest {
	r.datacenter = &datacenter
	return r
}
func (r ApiDatacentersPostRequest) Pretty(pretty bool) ApiDatacentersPostRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersPostRequest) Depth(depth int32) ApiDatacentersPostRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersPostRequest) XContractNumber(xContractNumber int32) ApiDatacentersPostRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersPostRequest) Execute() (Datacenter, *shared.APIResponse, error) {
	return r.ApiService.DatacentersPostExecute(r)
}

/*
  - DatacentersPost Create a Data Center
  - Creates new data centers, and data centers that already contain elements, such as servers and storage volumes.

Virtual data centers are the foundation of the platform; they act as logical containers for all other objects you create, such as servers and storage volumes. You can provision as many data centers as needed. Data centers have their own private networks and are logically segmented from each other to create isolation.
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @return ApiDatacentersPostRequest
*/
func (a *DataCentersApiService) DatacentersPost(ctx _context.Context) ApiDatacentersPostRequest {
	return ApiDatacentersPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCentersApiService) DatacentersPostExecute(r ApiDatacentersPostRequest) (Datacenter, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCentersApiService.DatacentersPost")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/datacenters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.datacenter == nil {
		return localVarReturnValue, nil, reportError("datacenter is required and must be specified")
	}

	if r.pretty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pretty", r.pretty, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.datacenter
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
		Operation:   "DatacentersPost",
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

type ApiDatacentersPutRequest struct {
	ctx             _context.Context
	ApiService      *DataCentersApiService
	datacenterId    string
	datacenter      *DatacenterPut
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiDatacentersPutRequest) Datacenter(datacenter DatacenterPut) ApiDatacentersPutRequest {
	r.datacenter = &datacenter
	return r
}
func (r ApiDatacentersPutRequest) Pretty(pretty bool) ApiDatacentersPutRequest {
	r.pretty = &pretty
	return r
}
func (r ApiDatacentersPutRequest) Depth(depth int32) ApiDatacentersPutRequest {
	r.depth = &depth
	return r
}
func (r ApiDatacentersPutRequest) XContractNumber(xContractNumber int32) ApiDatacentersPutRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiDatacentersPutRequest) Execute() (Datacenter, *shared.APIResponse, error) {
	return r.ApiService.DatacentersPutExecute(r)
}

/*
 * DatacentersPut Modify a Data Center by ID
 * Modifies the properties of the specified data center, rename it, or change the description.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param datacenterId The unique ID of the data center.
 * @return ApiDatacentersPutRequest
 */
func (a *DataCentersApiService) DatacentersPut(ctx _context.Context, datacenterId string) ApiDatacentersPutRequest {
	return ApiDatacentersPutRequest{
		ApiService:   a,
		ctx:          ctx,
		datacenterId: datacenterId,
	}
}

/*
 * Execute executes the request
 * @return Datacenter
 */
func (a *DataCentersApiService) DatacentersPutExecute(r ApiDatacentersPutRequest) (Datacenter, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Datacenter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataCentersApiService.DatacentersPut")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/datacenters/{datacenterId}"
	localVarPath = strings.Replace(localVarPath, "{"+"datacenterId"+"}", _neturl.PathEscape(parameterValueToString(r.datacenterId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.datacenter == nil {
		return localVarReturnValue, nil, reportError("datacenter is required and must be specified")
	}

	if r.pretty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pretty", r.pretty, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.datacenter
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
		Operation:   "DatacentersPut",
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
