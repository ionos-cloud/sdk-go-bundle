# Metadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 creation timestamp. | [optional] |
|**CreatedBy** | Pointer to **string** |  | [optional] |
|**CreatedByUserId** | Pointer to **string** |  | [optional] |
|**CreatedByUserUuid** | Pointer to **string** |  | [optional] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 modified timestamp. | [optional] |
|**LastModifiedBy** | Pointer to **string** |  | [optional] |
|**LastModifiedByUserId** | Pointer to **string** |  | [optional] |
|**LastModifiedByUserUuid** | Pointer to **string** |  | [optional] |
|**Status** | Pointer to **string** | The current status reported back by the pipeline. | [optional] |

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *Metadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *Metadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *Metadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *Metadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Metadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Metadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Metadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Metadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *Metadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *Metadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *Metadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *Metadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedByUserUuid

`func (o *Metadata) GetCreatedByUserUuid() string`

GetCreatedByUserUuid returns the CreatedByUserUuid field if non-nil, zero value otherwise.

### GetCreatedByUserUuidOk

`func (o *Metadata) GetCreatedByUserUuidOk() (*string, bool)`

GetCreatedByUserUuidOk returns a tuple with the CreatedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserUuid

`func (o *Metadata) SetCreatedByUserUuid(v string)`

SetCreatedByUserUuid sets CreatedByUserUuid field to given value.

### HasCreatedByUserUuid

`func (o *Metadata) HasCreatedByUserUuid() bool`

HasCreatedByUserUuid returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *Metadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *Metadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *Metadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *Metadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *Metadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *Metadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *Metadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *Metadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *Metadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *Metadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *Metadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *Metadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetLastModifiedByUserUuid

`func (o *Metadata) GetLastModifiedByUserUuid() string`

GetLastModifiedByUserUuid returns the LastModifiedByUserUuid field if non-nil, zero value otherwise.

### GetLastModifiedByUserUuidOk

`func (o *Metadata) GetLastModifiedByUserUuidOk() (*string, bool)`

GetLastModifiedByUserUuidOk returns a tuple with the LastModifiedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserUuid

`func (o *Metadata) SetLastModifiedByUserUuid(v string)`

SetLastModifiedByUserUuid sets LastModifiedByUserUuid field to given value.

### HasLastModifiedByUserUuid

`func (o *Metadata) HasLastModifiedByUserUuid() bool`

HasLastModifiedByUserUuid returns a boolean if a field has been set.

### GetStatus

`func (o *Metadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Metadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Metadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Metadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


