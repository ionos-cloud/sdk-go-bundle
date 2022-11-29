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

// PrivateCrossConnectsApiService PrivateCrossConnectsApi service
type PrivateCrossConnectsApiService service

type ApiPccsDeleteRequest struct {
	ctx             _context.Context
	ApiService      *PrivateCrossConnectsApiService
	pccId           string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiPccsDeleteRequest) Pretty(pretty bool) ApiPccsDeleteRequest {
	r.pretty = &pretty
	return r
}
func (r ApiPccsDeleteRequest) Depth(depth int32) ApiPccsDeleteRequest {
	r.depth = &depth
	return r
}
func (r ApiPccsDeleteRequest) XContractNumber(xContractNumber int32) ApiPccsDeleteRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiPccsDeleteRequest) Execute() (*common.APIResponse, error) {
	return r.ApiService.PccsDeleteExecute(r)
}

/*
 * PccsDelete Delete private Cross-Connects
 * Remove the specified private Cross-Connect (only if not connected to any data centers).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pccId The unique ID of the private Cross-Connect.
 * @return ApiPccsDeleteRequest
 */
func (a *PrivateCrossConnectsApiService) PccsDelete(ctx _context.Context, pccId string) ApiPccsDeleteRequest {
	return ApiPccsDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		pccId:      pccId,
	}
}

/*
 * Execute executes the request
 */
func (a *PrivateCrossConnectsApiService) PccsDeleteExecute(r ApiPccsDeleteRequest) (*common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateCrossConnectsApiService.PccsDelete")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/pccs/{pccId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pccId"+"}", _neturl.PathEscape(parameterToString(r.pccId, "")), -1)

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
		Operation:   "PccsDelete",
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

type ApiPccsFindByIdRequest struct {
	ctx             _context.Context
	ApiService      *PrivateCrossConnectsApiService
	pccId           string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiPccsFindByIdRequest) Pretty(pretty bool) ApiPccsFindByIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiPccsFindByIdRequest) Depth(depth int32) ApiPccsFindByIdRequest {
	r.depth = &depth
	return r
}
func (r ApiPccsFindByIdRequest) XContractNumber(xContractNumber int32) ApiPccsFindByIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiPccsFindByIdRequest) Execute() (PrivateCrossConnect, *common.APIResponse, error) {
	return r.ApiService.PccsFindByIdExecute(r)
}

/*
 * PccsFindById Retrieve private Cross-Connects
 * Retrieve a private Cross-Connect by the resource ID. Cross-Connect ID is in the response body when the private Cross-Connect is created, and in the list of private Cross-Connects, returned by GET.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pccId The unique ID of the private Cross-Connect.
 * @return ApiPccsFindByIdRequest
 */
func (a *PrivateCrossConnectsApiService) PccsFindById(ctx _context.Context, pccId string) ApiPccsFindByIdRequest {
	return ApiPccsFindByIdRequest{
		ApiService: a,
		ctx:        ctx,
		pccId:      pccId,
	}
}

/*
 * Execute executes the request
 * @return PrivateCrossConnect
 */
func (a *PrivateCrossConnectsApiService) PccsFindByIdExecute(r ApiPccsFindByIdRequest) (PrivateCrossConnect, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrivateCrossConnect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateCrossConnectsApiService.PccsFindById")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/pccs/{pccId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pccId"+"}", _neturl.PathEscape(parameterToString(r.pccId, "")), -1)

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
		Operation:   "PccsFindById",
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

type ApiPccsGetRequest struct {
	ctx        _context.Context
	ApiService *PrivateCrossConnectsApiService
	Params
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiPccsGetRequest) Pretty(pretty bool) ApiPccsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiPccsGetRequest) Depth(depth int32) ApiPccsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiPccsGetRequest) XContractNumber(xContractNumber int32) ApiPccsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiPccsGetRequest) Execute() (PrivateCrossConnects, *common.APIResponse, error) {
	return r.ApiService.PccsGetExecute(r)
}

/*
 * PccsGet List private Cross-Connects
 * List all private Cross-Connects for your account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPccsGetRequest
 */
func (a *PrivateCrossConnectsApiService) PccsGet(ctx _context.Context) ApiPccsGetRequest {
	return ApiPccsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return PrivateCrossConnects
 */
func (a *PrivateCrossConnectsApiService) PccsGetExecute(r ApiPccsGetRequest) (PrivateCrossConnects, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrivateCrossConnects
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateCrossConnectsApiService.PccsGet")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/pccs"

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
		Operation:   "PccsGet",
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

type ApiPccsPatchRequest struct {
	ctx             _context.Context
	ApiService      *PrivateCrossConnectsApiService
	pccId           string
	pcc             *PrivateCrossConnectProperties
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiPccsPatchRequest) Pcc(pcc PrivateCrossConnectProperties) ApiPccsPatchRequest {
	r.pcc = &pcc
	return r
}
func (r ApiPccsPatchRequest) Pretty(pretty bool) ApiPccsPatchRequest {
	r.pretty = &pretty
	return r
}
func (r ApiPccsPatchRequest) Depth(depth int32) ApiPccsPatchRequest {
	r.depth = &depth
	return r
}
func (r ApiPccsPatchRequest) XContractNumber(xContractNumber int32) ApiPccsPatchRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiPccsPatchRequest) Execute() (PrivateCrossConnect, *common.APIResponse, error) {
	return r.ApiService.PccsPatchExecute(r)
}

/*
 * PccsPatch Partially modify private Cross-Connects
 * Update the properties of the specified private Cross-Connect.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pccId The unique ID of the private Cross-Connect.
 * @return ApiPccsPatchRequest
 */
func (a *PrivateCrossConnectsApiService) PccsPatch(ctx _context.Context, pccId string) ApiPccsPatchRequest {
	return ApiPccsPatchRequest{
		ApiService: a,
		ctx:        ctx,
		pccId:      pccId,
	}
}

/*
 * Execute executes the request
 * @return PrivateCrossConnect
 */
func (a *PrivateCrossConnectsApiService) PccsPatchExecute(r ApiPccsPatchRequest) (PrivateCrossConnect, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrivateCrossConnect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateCrossConnectsApiService.PccsPatch")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/pccs/{pccId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pccId"+"}", _neturl.PathEscape(parameterToString(r.pccId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pcc == nil {
		return localVarReturnValue, nil, reportError("pcc is required and must be specified")
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
	localVarPostBody = r.pcc
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
		Operation:   "PccsPatch",
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

type ApiPccsPostRequest struct {
	ctx             _context.Context
	ApiService      *PrivateCrossConnectsApiService
	pcc             *PrivateCrossConnect
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiPccsPostRequest) Pcc(pcc PrivateCrossConnect) ApiPccsPostRequest {
	r.pcc = &pcc
	return r
}
func (r ApiPccsPostRequest) Pretty(pretty bool) ApiPccsPostRequest {
	r.pretty = &pretty
	return r
}
func (r ApiPccsPostRequest) Depth(depth int32) ApiPccsPostRequest {
	r.depth = &depth
	return r
}
func (r ApiPccsPostRequest) XContractNumber(xContractNumber int32) ApiPccsPostRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiPccsPostRequest) Execute() (PrivateCrossConnect, *common.APIResponse, error) {
	return r.ApiService.PccsPostExecute(r)
}

/*
 * PccsPost Create private Cross-Connects
 * Create a private Cross-Connect.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPccsPostRequest
 */
func (a *PrivateCrossConnectsApiService) PccsPost(ctx _context.Context) ApiPccsPostRequest {
	return ApiPccsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return PrivateCrossConnect
 */
func (a *PrivateCrossConnectsApiService) PccsPostExecute(r ApiPccsPostRequest) (PrivateCrossConnect, *common.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PrivateCrossConnect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateCrossConnectsApiService.PccsPost")
	if err != nil {
		gerr := common.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/pccs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.pcc == nil {
		return localVarReturnValue, nil, reportError("pcc is required and must be specified")
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
	localVarPostBody = r.pcc
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
		Operation:   "PccsPost",
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