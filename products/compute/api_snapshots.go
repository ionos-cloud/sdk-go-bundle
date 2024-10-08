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

// SnapshotsApiService SnapshotsApi service
type SnapshotsApiService service

type ApiSnapshotsDeleteRequest struct {
	ctx             _context.Context
	ApiService      *SnapshotsApiService
	snapshotId      string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiSnapshotsDeleteRequest) Pretty(pretty bool) ApiSnapshotsDeleteRequest {
	r.pretty = &pretty
	return r
}
func (r ApiSnapshotsDeleteRequest) Depth(depth int32) ApiSnapshotsDeleteRequest {
	r.depth = &depth
	return r
}
func (r ApiSnapshotsDeleteRequest) XContractNumber(xContractNumber int32) ApiSnapshotsDeleteRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiSnapshotsDeleteRequest) Execute() (*shared.APIResponse, error) {
	return r.ApiService.SnapshotsDeleteExecute(r)
}

/*
 * SnapshotsDelete Delete snapshots
 * Deletes the specified snapshot.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param snapshotId The unique ID of the snapshot.
 * @return ApiSnapshotsDeleteRequest
 */
func (a *SnapshotsApiService) SnapshotsDelete(ctx _context.Context, snapshotId string) ApiSnapshotsDeleteRequest {
	return ApiSnapshotsDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		snapshotId: snapshotId,
	}
}

/*
 * Execute executes the request
 */
func (a *SnapshotsApiService) SnapshotsDeleteExecute(r ApiSnapshotsDeleteRequest) (*shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.SnapshotsDelete")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", _neturl.PathEscape(parameterValueToString(r.snapshotId, "")), -1)

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
		Operation:   "SnapshotsDelete",
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

type ApiSnapshotsFindByIdRequest struct {
	ctx             _context.Context
	ApiService      *SnapshotsApiService
	snapshotId      string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiSnapshotsFindByIdRequest) Pretty(pretty bool) ApiSnapshotsFindByIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiSnapshotsFindByIdRequest) Depth(depth int32) ApiSnapshotsFindByIdRequest {
	r.depth = &depth
	return r
}
func (r ApiSnapshotsFindByIdRequest) XContractNumber(xContractNumber int32) ApiSnapshotsFindByIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiSnapshotsFindByIdRequest) Execute() (Snapshot, *shared.APIResponse, error) {
	return r.ApiService.SnapshotsFindByIdExecute(r)
}

/*
 * SnapshotsFindById Retrieve snapshots by ID
 * Retrieve the properties of the specified snapshot.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param snapshotId The unique ID of the snapshot.
 * @return ApiSnapshotsFindByIdRequest
 */
func (a *SnapshotsApiService) SnapshotsFindById(ctx _context.Context, snapshotId string) ApiSnapshotsFindByIdRequest {
	return ApiSnapshotsFindByIdRequest{
		ApiService: a,
		ctx:        ctx,
		snapshotId: snapshotId,
	}
}

/*
 * Execute executes the request
 * @return Snapshot
 */
func (a *SnapshotsApiService) SnapshotsFindByIdExecute(r ApiSnapshotsFindByIdRequest) (Snapshot, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.SnapshotsFindById")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", _neturl.PathEscape(parameterValueToString(r.snapshotId, "")), -1)

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
		Operation:   "SnapshotsFindById",
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

type ApiSnapshotsGetRequest struct {
	ctx             _context.Context
	ApiService      *SnapshotsApiService
	filters         _neturl.Values
	orderBy         *string
	maxResults      *int32
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiSnapshotsGetRequest) Pretty(pretty bool) ApiSnapshotsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiSnapshotsGetRequest) Depth(depth int32) ApiSnapshotsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiSnapshotsGetRequest) XContractNumber(xContractNumber int32) ApiSnapshotsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

// Filters query parameters limit results to those containing a matching value for a specific property.
func (r ApiSnapshotsGetRequest) Filter(key string, value string) ApiSnapshotsGetRequest {
	filterKey := fmt.Sprintf("filter.%s", key)
	r.filters[filterKey] = append(r.filters[filterKey], value)
	return r
}

// OrderBy query param sorts the results alphanumerically in ascending order based on the specified property.
func (r ApiSnapshotsGetRequest) OrderBy(orderBy string) ApiSnapshotsGetRequest {
	r.orderBy = &orderBy
	return r
}

// MaxResults query param limits the number of results returned.
func (r ApiSnapshotsGetRequest) MaxResults(maxResults int32) ApiSnapshotsGetRequest {
	r.maxResults = &maxResults
	return r
}

func (r ApiSnapshotsGetRequest) Execute() (Snapshots, *shared.APIResponse, error) {
	return r.ApiService.SnapshotsGetExecute(r)
}

/*
 * SnapshotsGet List snapshots
 * List all available snapshots.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSnapshotsGetRequest
 */
func (a *SnapshotsApiService) SnapshotsGet(ctx _context.Context) ApiSnapshotsGetRequest {
	return ApiSnapshotsGetRequest{
		ApiService: a,
		ctx:        ctx,
		filters:    _neturl.Values{},
	}
}

/*
 * Execute executes the request
 * @return Snapshots
 */
func (a *SnapshotsApiService) SnapshotsGetExecute(r ApiSnapshotsGetRequest) (Snapshots, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Snapshots
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.SnapshotsGet")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/snapshots"

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
		Operation:   "SnapshotsGet",
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

type ApiSnapshotsPatchRequest struct {
	ctx             _context.Context
	ApiService      *SnapshotsApiService
	snapshotId      string
	snapshot        *SnapshotProperties
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiSnapshotsPatchRequest) Snapshot(snapshot SnapshotProperties) ApiSnapshotsPatchRequest {
	r.snapshot = &snapshot
	return r
}
func (r ApiSnapshotsPatchRequest) Pretty(pretty bool) ApiSnapshotsPatchRequest {
	r.pretty = &pretty
	return r
}
func (r ApiSnapshotsPatchRequest) Depth(depth int32) ApiSnapshotsPatchRequest {
	r.depth = &depth
	return r
}
func (r ApiSnapshotsPatchRequest) XContractNumber(xContractNumber int32) ApiSnapshotsPatchRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiSnapshotsPatchRequest) Execute() (Snapshot, *shared.APIResponse, error) {
	return r.ApiService.SnapshotsPatchExecute(r)
}

/*
 * SnapshotsPatch Partially modify snapshots
 * Update the properties of the specified snapshot.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param snapshotId The unique ID of the snapshot.
 * @return ApiSnapshotsPatchRequest
 */
func (a *SnapshotsApiService) SnapshotsPatch(ctx _context.Context, snapshotId string) ApiSnapshotsPatchRequest {
	return ApiSnapshotsPatchRequest{
		ApiService: a,
		ctx:        ctx,
		snapshotId: snapshotId,
	}
}

/*
 * Execute executes the request
 * @return Snapshot
 */
func (a *SnapshotsApiService) SnapshotsPatchExecute(r ApiSnapshotsPatchRequest) (Snapshot, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.SnapshotsPatch")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", _neturl.PathEscape(parameterValueToString(r.snapshotId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.snapshot == nil {
		return localVarReturnValue, nil, reportError("snapshot is required and must be specified")
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
	localVarPostBody = r.snapshot
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
		Operation:   "SnapshotsPatch",
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

type ApiSnapshotsPutRequest struct {
	ctx             _context.Context
	ApiService      *SnapshotsApiService
	snapshotId      string
	snapshot        *Snapshot
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiSnapshotsPutRequest) Snapshot(snapshot Snapshot) ApiSnapshotsPutRequest {
	r.snapshot = &snapshot
	return r
}
func (r ApiSnapshotsPutRequest) Pretty(pretty bool) ApiSnapshotsPutRequest {
	r.pretty = &pretty
	return r
}
func (r ApiSnapshotsPutRequest) Depth(depth int32) ApiSnapshotsPutRequest {
	r.depth = &depth
	return r
}
func (r ApiSnapshotsPutRequest) XContractNumber(xContractNumber int32) ApiSnapshotsPutRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiSnapshotsPutRequest) Execute() (Snapshot, *shared.APIResponse, error) {
	return r.ApiService.SnapshotsPutExecute(r)
}

/*
 * SnapshotsPut Modify a Snapshot by ID
 * Modifies the properties of the specified snapshot.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param snapshotId The unique ID of the snapshot.
 * @return ApiSnapshotsPutRequest
 */
func (a *SnapshotsApiService) SnapshotsPut(ctx _context.Context, snapshotId string) ApiSnapshotsPutRequest {
	return ApiSnapshotsPutRequest{
		ApiService: a,
		ctx:        ctx,
		snapshotId: snapshotId,
	}
}

/*
 * Execute executes the request
 * @return Snapshot
 */
func (a *SnapshotsApiService) SnapshotsPutExecute(r ApiSnapshotsPutRequest) (Snapshot, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Snapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SnapshotsApiService.SnapshotsPut")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", _neturl.PathEscape(parameterValueToString(r.snapshotId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.snapshot == nil {
		return localVarReturnValue, nil, reportError("snapshot is required and must be specified")
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
	localVarPostBody = r.snapshot
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
		Operation:   "SnapshotsPut",
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
