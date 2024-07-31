# CentralLoggingResponseMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 modified timestamp. | [optional] |
|**LastModifiedBy** | Pointer to **string** |  | [optional] |
|**LastModifiedByUserId** | Pointer to **string** |  | [optional] |
|**LastModifiedByUserUuid** | Pointer to **string** |  | [optional] |
|**GrafanaEndpoint** | Pointer to **string** |  | [optional] |

## Methods

### NewCentralLoggingResponseMetadata

`func NewCentralLoggingResponseMetadata() *CentralLoggingResponseMetadata`

NewCentralLoggingResponseMetadata instantiates a new CentralLoggingResponseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralLoggingResponseMetadataWithDefaults

`func NewCentralLoggingResponseMetadataWithDefaults() *CentralLoggingResponseMetadata`

NewCentralLoggingResponseMetadataWithDefaults instantiates a new CentralLoggingResponseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModifiedDate

`func (o *CentralLoggingResponseMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *CentralLoggingResponseMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *CentralLoggingResponseMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *CentralLoggingResponseMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *CentralLoggingResponseMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *CentralLoggingResponseMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *CentralLoggingResponseMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *CentralLoggingResponseMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *CentralLoggingResponseMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *CentralLoggingResponseMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *CentralLoggingResponseMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *CentralLoggingResponseMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetLastModifiedByUserUuid

`func (o *CentralLoggingResponseMetadata) GetLastModifiedByUserUuid() string`

GetLastModifiedByUserUuid returns the LastModifiedByUserUuid field if non-nil, zero value otherwise.

### GetLastModifiedByUserUuidOk

`func (o *CentralLoggingResponseMetadata) GetLastModifiedByUserUuidOk() (*string, bool)`

GetLastModifiedByUserUuidOk returns a tuple with the LastModifiedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserUuid

`func (o *CentralLoggingResponseMetadata) SetLastModifiedByUserUuid(v string)`

SetLastModifiedByUserUuid sets LastModifiedByUserUuid field to given value.

### HasLastModifiedByUserUuid

`func (o *CentralLoggingResponseMetadata) HasLastModifiedByUserUuid() bool`

HasLastModifiedByUserUuid returns a boolean if a field has been set.

### GetGrafanaEndpoint

`func (o *CentralLoggingResponseMetadata) GetGrafanaEndpoint() string`

GetGrafanaEndpoint returns the GrafanaEndpoint field if non-nil, zero value otherwise.

### GetGrafanaEndpointOk

`func (o *CentralLoggingResponseMetadata) GetGrafanaEndpointOk() (*string, bool)`

GetGrafanaEndpointOk returns a tuple with the GrafanaEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaEndpoint

`func (o *CentralLoggingResponseMetadata) SetGrafanaEndpoint(v string)`

SetGrafanaEndpoint sets GrafanaEndpoint field to given value.

### HasGrafanaEndpoint

`func (o *CentralLoggingResponseMetadata) HasGrafanaEndpoint() bool`

HasGrafanaEndpoint returns a boolean if a field has been set.


