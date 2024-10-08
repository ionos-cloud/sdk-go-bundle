/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature ### 1.2.0  - Added registry `apiSubnetAllowList` ### 1.2.1  - Amended `apiSubnetAllowList` Regex
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

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

// RegistriesApiService RegistriesApi service
type RegistriesApiService service

type ApiRegistriesDeleteRequest struct {
	ctx        _context.Context
	ApiService *RegistriesApiService
	registryId string
}

func (r ApiRegistriesDeleteRequest) Execute() (*shared.APIResponse, error) {
	return r.ApiService.RegistriesDeleteExecute(r)
}

/*
 * RegistriesDelete Delete registry
 *
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId The unique ID of the registry
 * @return ApiRegistriesDeleteRequest
 */
func (a *RegistriesApiService) RegistriesDelete(ctx _context.Context, registryId string) ApiRegistriesDeleteRequest {
	return ApiRegistriesDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
	}
}

/*
 * Execute executes the request
 */
func (a *RegistriesApiService) RegistriesDeleteExecute(r ApiRegistriesDeleteRequest) (*shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.RegistriesDelete")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return nil, gerr
	}

	localVarPath := localBasePath + "/registries/{registryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterValueToString(r.registryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
		Operation:   "RegistriesDelete",
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
		return localVarAPIResponse, newErr
	}

	return localVarAPIResponse, nil
}

type ApiRegistriesFindByIdRequest struct {
	ctx        _context.Context
	ApiService *RegistriesApiService
	registryId string
}

func (r ApiRegistriesFindByIdRequest) Execute() (RegistryResponse, *shared.APIResponse, error) {
	return r.ApiService.RegistriesFindByIdExecute(r)
}

/*
 * RegistriesFindById Get a registry
 * Get all information for a specific container registry
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param registryId The unique ID of the registry
 * @return ApiRegistriesFindByIdRequest
 */
func (a *RegistriesApiService) RegistriesFindById(ctx _context.Context, registryId string) ApiRegistriesFindByIdRequest {
	return ApiRegistriesFindByIdRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
	}
}

/*
 * Execute executes the request
 * @return RegistryResponse
 */
func (a *RegistriesApiService) RegistriesFindByIdExecute(r ApiRegistriesFindByIdRequest) (RegistryResponse, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegistryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.RegistriesFindById")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/registries/{registryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterValueToString(r.registryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
		Operation:   "RegistriesFindById",
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
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

type ApiRegistriesGetRequest struct {
	ctx             _context.Context
	ApiService      *RegistriesApiService
	filterName      *string
	limit           *string
	paginationToken *string
}

func (r ApiRegistriesGetRequest) FilterName(filterName string) ApiRegistriesGetRequest {
	r.filterName = &filterName
	return r
}
func (r ApiRegistriesGetRequest) Limit(limit string) ApiRegistriesGetRequest {
	r.limit = &limit
	return r
}
func (r ApiRegistriesGetRequest) PaginationToken(paginationToken string) ApiRegistriesGetRequest {
	r.paginationToken = &paginationToken
	return r
}

func (r ApiRegistriesGetRequest) Execute() (RegistriesResponse, *shared.APIResponse, error) {
	return r.ApiService.RegistriesGetExecute(r)
}

/*
 * RegistriesGet List all container registries
 * List all managed container registries for your account
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRegistriesGetRequest
 */
func (a *RegistriesApiService) RegistriesGet(ctx _context.Context) ApiRegistriesGetRequest {
	return ApiRegistriesGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return RegistriesResponse
 */
func (a *RegistriesApiService) RegistriesGetExecute(r ApiRegistriesGetRequest) (RegistriesResponse, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegistriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.RegistriesGet")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/registries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filterName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter.name", r.filterName, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.paginationToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagination.token", r.paginationToken, "")
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
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
		Operation:   "RegistriesGet",
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
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

type ApiRegistriesPatchRequest struct {
	ctx                _context.Context
	ApiService         *RegistriesApiService
	registryId         string
	patchRegistryInput *PatchRegistryInput
}

func (r ApiRegistriesPatchRequest) PatchRegistryInput(patchRegistryInput PatchRegistryInput) ApiRegistriesPatchRequest {
	r.patchRegistryInput = &patchRegistryInput
	return r
}

func (r ApiRegistriesPatchRequest) Execute() (RegistryResponse, *shared.APIResponse, error) {
	return r.ApiService.RegistriesPatchExecute(r)
}

/*
  - RegistriesPatch Update the properties of a registry
  - Update the properties of a registry

- "garbageCollectionSchedule" time and days of the week for runs
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryId The unique ID of the registry
  - @return ApiRegistriesPatchRequest
*/
func (a *RegistriesApiService) RegistriesPatch(ctx _context.Context, registryId string) ApiRegistriesPatchRequest {
	return ApiRegistriesPatchRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
	}
}

/*
 * Execute executes the request
 * @return RegistryResponse
 */
func (a *RegistriesApiService) RegistriesPatchExecute(r ApiRegistriesPatchRequest) (RegistryResponse, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegistryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.RegistriesPatch")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/registries/{registryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterValueToString(r.registryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.patchRegistryInput == nil {
		return localVarReturnValue, nil, reportError("patchRegistryInput is required and must be specified")
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
	// body params
	localVarPostBody = r.patchRegistryInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
		Operation:   "RegistriesPatch",
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
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

type ApiRegistriesPostRequest struct {
	ctx               _context.Context
	ApiService        *RegistriesApiService
	postRegistryInput *PostRegistryInput
}

func (r ApiRegistriesPostRequest) PostRegistryInput(postRegistryInput PostRegistryInput) ApiRegistriesPostRequest {
	r.postRegistryInput = &postRegistryInput
	return r
}

func (r ApiRegistriesPostRequest) Execute() (PostRegistryOutput, *shared.APIResponse, error) {
	return r.ApiService.RegistriesPostExecute(r)
}

/*
  - RegistriesPost Create container registry
  - Create a registry to hold container images or OCI compliant artifacts

- "name" must have passed validation
- "location" must be one of the available location IDs
- "garbageCollectionSchedule" time and days of the week for runs
- "features": "vulnerabilityScanning" default is enabled
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @return ApiRegistriesPostRequest
*/
func (a *RegistriesApiService) RegistriesPost(ctx _context.Context) ApiRegistriesPostRequest {
	return ApiRegistriesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return PostRegistryOutput
 */
func (a *RegistriesApiService) RegistriesPostExecute(r ApiRegistriesPostRequest) (PostRegistryOutput, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PostRegistryOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.RegistriesPost")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/registries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.postRegistryInput == nil {
		return localVarReturnValue, nil, reportError("postRegistryInput is required and must be specified")
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
	// body params
	localVarPostBody = r.postRegistryInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
		Operation:   "RegistriesPost",
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
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

type ApiRegistriesPutRequest struct {
	ctx              _context.Context
	ApiService       *RegistriesApiService
	registryId       string
	putRegistryInput *PutRegistryInput
}

func (r ApiRegistriesPutRequest) PutRegistryInput(putRegistryInput PutRegistryInput) ApiRegistriesPutRequest {
	r.putRegistryInput = &putRegistryInput
	return r
}

func (r ApiRegistriesPutRequest) Execute() (PutRegistryOutput, *shared.APIResponse, error) {
	return r.ApiService.RegistriesPutExecute(r)
}

/*
  - RegistriesPut Create or replace a container registry
  - Create/replace a registry to hold container images or OCI compliant

artifacts
**On create**
- "name" must have passed validation
- "location" must be one of the available location IDs
**On update**
- "name" cannot be changed
- "location" cannot be changed
**On create or update**
- "garbageCollectionSchedule": time and days of the week for runs

  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryId The unique ID of the registry
  - @return ApiRegistriesPutRequest
*/
func (a *RegistriesApiService) RegistriesPut(ctx _context.Context, registryId string) ApiRegistriesPutRequest {
	return ApiRegistriesPutRequest{
		ApiService: a,
		ctx:        ctx,
		registryId: registryId,
	}
}

/*
 * Execute executes the request
 * @return PutRegistryOutput
 */
func (a *RegistriesApiService) RegistriesPutExecute(r ApiRegistriesPutRequest) (PutRegistryOutput, *shared.APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PutRegistryOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistriesApiService.RegistriesPut")
	if err != nil {
		gerr := shared.GenericOpenAPIError{}
		gerr.SetError(err.Error())
		return localVarReturnValue, nil, gerr
	}

	localVarPath := localBasePath + "/registries/{registryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registryId"+"}", _neturl.PathEscape(parameterValueToString(r.registryId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.putRegistryInput == nil {
		return localVarReturnValue, nil, reportError("putRegistryInput is required and must be specified")
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
	// body params
	localVarPostBody = r.putRegistryInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(shared.ContextAPIKeys).(map[string]shared.APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
		Operation:   "RegistriesPut",
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarAPIResponse, newErr
			}
			newErr.SetModel(v)
		}
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
