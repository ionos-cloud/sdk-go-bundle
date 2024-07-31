# DeletedMetadata

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
|**State** | Pointer to **string** | The current state reported back by the pipeline. | [optional] |

## Methods

### NewDeletedMetadata

`func NewDeletedMetadata() *DeletedMetadata`

NewDeletedMetadata instantiates a new DeletedMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedMetadataWithDefaults

`func NewDeletedMetadataWithDefaults() *DeletedMetadata`

NewDeletedMetadataWithDefaults instantiates a new DeletedMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *DeletedMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DeletedMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DeletedMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DeletedMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DeletedMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeletedMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeletedMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeletedMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *DeletedMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *DeletedMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *DeletedMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *DeletedMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedByUserUuid

`func (o *DeletedMetadata) GetCreatedByUserUuid() string`

GetCreatedByUserUuid returns the CreatedByUserUuid field if non-nil, zero value otherwise.

### GetCreatedByUserUuidOk

`func (o *DeletedMetadata) GetCreatedByUserUuidOk() (*string, bool)`

GetCreatedByUserUuidOk returns a tuple with the CreatedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserUuid

`func (o *DeletedMetadata) SetCreatedByUserUuid(v string)`

SetCreatedByUserUuid sets CreatedByUserUuid field to given value.

### HasCreatedByUserUuid

`func (o *DeletedMetadata) HasCreatedByUserUuid() bool`

HasCreatedByUserUuid returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *DeletedMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *DeletedMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *DeletedMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *DeletedMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *DeletedMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *DeletedMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *DeletedMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *DeletedMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *DeletedMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *DeletedMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *DeletedMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *DeletedMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetLastModifiedByUserUuid

`func (o *DeletedMetadata) GetLastModifiedByUserUuid() string`

GetLastModifiedByUserUuid returns the LastModifiedByUserUuid field if non-nil, zero value otherwise.

### GetLastModifiedByUserUuidOk

`func (o *DeletedMetadata) GetLastModifiedByUserUuidOk() (*string, bool)`

GetLastModifiedByUserUuidOk returns a tuple with the LastModifiedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserUuid

`func (o *DeletedMetadata) SetLastModifiedByUserUuid(v string)`

SetLastModifiedByUserUuid sets LastModifiedByUserUuid field to given value.

### HasLastModifiedByUserUuid

`func (o *DeletedMetadata) HasLastModifiedByUserUuid() bool`

HasLastModifiedByUserUuid returns a boolean if a field has been set.

### GetState

`func (o *DeletedMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeletedMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeletedMetadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeletedMetadata) HasState() bool`

HasState returns a boolean if a field has been set.


