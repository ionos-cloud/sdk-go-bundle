/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	_context "context"
	"fmt"
	"github.com/ionos-cloud/sdk-go-bundle/common"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// TargetGroupsApiService TargetGroupsApi service
type TargetGroupsApiService service

type ApiTargetGroupsDeleteRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroupId   string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetGroupsDeleteRequest) Pretty(pretty bool) ApiTargetGroupsDeleteRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetGroupsDeleteRequest) Depth(depth int32) ApiTargetGroupsDeleteRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetGroupsDeleteRequest) XContractNumber(xContractNumber int32) ApiTargetGroupsDeleteRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetGroupsDeleteRequest) Execute() (*common.APIResponse, error) {
	return r.ApiService.TargetGroupsDeleteExecute(r)
}

/*
 * TargetGroupsDelete Remove target groups
 * Remove the specified target group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetGroupsDeleteRequest
 */
func (a *TargetGroupsApiService) TargetGroupsDelete(ctx _context.Context, targetGroupId string) ApiTargetGroupsDeleteRequest {
	return ApiTargetGroupsDeleteRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 */
func (a *TargetGroupsApiService) TargetGroupsDeleteExecute(r ApiTargetGroupsDeleteRequest) (*common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetGroupsDelete")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
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
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarAPIResponse := &common.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "TargetGroupsDelete",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{}
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

type ApiTargetgroupsFindByTargetGroupIdRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroupId   string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetgroupsFindByTargetGroupIdRequest) Pretty(pretty bool) ApiTargetgroupsFindByTargetGroupIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsFindByTargetGroupIdRequest) Depth(depth int32) ApiTargetgroupsFindByTargetGroupIdRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsFindByTargetGroupIdRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsFindByTargetGroupIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsFindByTargetGroupIdRequest) Execute() (TargetGroup, *common.APIResponse, error) {
	return r.ApiService.TargetgroupsFindByTargetGroupIdExecute(r)
}

/*
 * TargetgroupsFindByTargetGroupId Retrieve target groups
 * Retrieve the properties of the specified target group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetgroupsFindByTargetGroupIdRequest
 */
func (a *TargetGroupsApiService) TargetgroupsFindByTargetGroupId(ctx _context.Context, targetGroupId string) ApiTargetgroupsFindByTargetGroupIdRequest {
	return ApiTargetgroupsFindByTargetGroupIdRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsFindByTargetGroupIdExecute(r ApiTargetgroupsFindByTargetGroupIdRequest) (TargetGroup, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsFindByTargetGroupId")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
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
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarAPIResponse := &common.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "TargetgroupsFindByTargetGroupId",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{}
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
		newErr := common.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsGetRequest struct {
	ctx        _context.Context
	ApiService *TargetGroupsApiService
	Params
	pretty          *bool
	depth           *int32
	xContractNumber *int32
	offset          *int32
	limit           *int32
}

func (r ApiTargetgroupsGetRequest) Pretty(pretty bool) ApiTargetgroupsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsGetRequest) Depth(depth int32) ApiTargetgroupsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsGetRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}
func (r ApiTargetgroupsGetRequest) Offset(offset int32) ApiTargetgroupsGetRequest {
	r.offset = &offset
	return r
}
func (r ApiTargetgroupsGetRequest) Limit(limit int32) ApiTargetgroupsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiTargetgroupsGetRequest) Execute() (TargetGroups, *common.APIResponse, error) {
	return r.ApiService.TargetgroupsGetExecute(r)
}

/*
 * TargetgroupsGet List target groups
 * List all target groups.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTargetgroupsGetRequest
 */
func (a *TargetGroupsApiService) TargetgroupsGet(ctx _context.Context) ApiTargetgroupsGetRequest {
	return ApiTargetgroupsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return TargetGroups
 */
func (a *TargetGroupsApiService) TargetgroupsGetExecute(r ApiTargetgroupsGetRequest) (TargetGroups, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroups
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsGet")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/targetgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarAPIResponse := &common.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "TargetgroupsGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{}
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
		newErr := common.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsPatchRequest struct {
	ctx                   _context.Context
	ApiService            *TargetGroupsApiService
	targetGroupId         string
	targetGroupProperties *TargetGroupProperties
	pretty                *bool
	depth                 *int32
	xContractNumber       *int32
}

func (r ApiTargetgroupsPatchRequest) TargetGroupProperties(targetGroupProperties TargetGroupProperties) ApiTargetgroupsPatchRequest {
	r.targetGroupProperties = &targetGroupProperties
	return r
}
func (r ApiTargetgroupsPatchRequest) Pretty(pretty bool) ApiTargetgroupsPatchRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsPatchRequest) Depth(depth int32) ApiTargetgroupsPatchRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsPatchRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsPatchRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsPatchRequest) Execute() (TargetGroup, *common.APIResponse, error) {
	return r.ApiService.TargetgroupsPatchExecute(r)
}

/*
 * TargetgroupsPatch Partially modify target groups
 * Update the properties of the specified target group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetgroupsPatchRequest
 */
func (a *TargetGroupsApiService) TargetgroupsPatch(ctx _context.Context, targetGroupId string) ApiTargetgroupsPatchRequest {
	return ApiTargetgroupsPatchRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsPatchExecute(r ApiTargetgroupsPatchRequest) (TargetGroup, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsPatch")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.targetGroupProperties == nil {
		return localVarReturnValue, nil, reportError("targetGroupProperties is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
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
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.targetGroupProperties
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarAPIResponse := &common.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "TargetgroupsPatch",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{}
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
		newErr := common.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsPostRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroup     *TargetGroup
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetgroupsPostRequest) TargetGroup(targetGroup TargetGroup) ApiTargetgroupsPostRequest {
	r.targetGroup = &targetGroup
	return r
}
func (r ApiTargetgroupsPostRequest) Pretty(pretty bool) ApiTargetgroupsPostRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsPostRequest) Depth(depth int32) ApiTargetgroupsPostRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsPostRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsPostRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsPostRequest) Execute() (TargetGroup, *common.APIResponse, error) {
	return r.ApiService.TargetgroupsPostExecute(r)
}

/*
 * TargetgroupsPost Create target groups
 * Create a target group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTargetgroupsPostRequest
 */
func (a *TargetGroupsApiService) TargetgroupsPost(ctx _context.Context) ApiTargetgroupsPostRequest {
	return ApiTargetgroupsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsPostExecute(r ApiTargetgroupsPostRequest) (TargetGroup, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsPost")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/targetgroups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.targetGroup == nil {
		return localVarReturnValue, nil, reportError("targetGroup is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
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
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.targetGroup
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarAPIResponse := &common.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "TargetgroupsPost",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{}
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
		newErr := common.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiTargetgroupsPutRequest struct {
	ctx             _context.Context
	ApiService      *TargetGroupsApiService
	targetGroupId   string
	targetGroup     *TargetGroupPut
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiTargetgroupsPutRequest) TargetGroup(targetGroup TargetGroupPut) ApiTargetgroupsPutRequest {
	r.targetGroup = &targetGroup
	return r
}
func (r ApiTargetgroupsPutRequest) Pretty(pretty bool) ApiTargetgroupsPutRequest {
	r.pretty = &pretty
	return r
}
func (r ApiTargetgroupsPutRequest) Depth(depth int32) ApiTargetgroupsPutRequest {
	r.depth = &depth
	return r
}
func (r ApiTargetgroupsPutRequest) XContractNumber(xContractNumber int32) ApiTargetgroupsPutRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiTargetgroupsPutRequest) Execute() (TargetGroup, *common.APIResponse, error) {
	return r.ApiService.TargetgroupsPutExecute(r)
}

/*
 * TargetgroupsPut Modify target groups
 * Modify the properties of the specified target group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param targetGroupId The unique ID of the target group.
 * @return ApiTargetgroupsPutRequest
 */
func (a *TargetGroupsApiService) TargetgroupsPut(ctx _context.Context, targetGroupId string) ApiTargetgroupsPutRequest {
	return ApiTargetgroupsPutRequest{
		ApiService:    a,
		ctx:           ctx,
		targetGroupId: targetGroupId,
	}
}

/*
 * Execute executes the request
 * @return TargetGroup
 */
func (a *TargetGroupsApiService) TargetgroupsPutExecute(r ApiTargetgroupsPutRequest) (TargetGroup, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TargetGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TargetGroupsApiService.TargetgroupsPut")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/targetgroups/{targetGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"targetGroupId"+"}", _neturl.PathEscape(parameterToString(r.targetGroupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.targetGroup == nil {
		return localVarReturnValue, nil, reportError("targetGroup is required and must be specified")
	}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
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
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	// body params
	localVarPostBody = r.targetGroup
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
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

	localVarAPIResponse := &common.APIResponse{
		Response:    localVarHTTPResponse,
		Method:      localVarHTTPMethod,
		RequestTime: httpRequestTime,
		RequestURL:  localVarPath,
		Operation:   "TargetgroupsPut",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{}
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
		newErr := common.GenericOpenAPIError{}
		newErr.SetStatusCode(localVarHTTPResponse.StatusCode)
		newErr.SetBody(localVarBody)
		newErr.SetError(err.Error())
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}
