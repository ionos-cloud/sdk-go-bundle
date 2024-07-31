# ProvisioningMetadata

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

### NewProvisioningMetadata

`func NewProvisioningMetadata() *ProvisioningMetadata`

NewProvisioningMetadata instantiates a new ProvisioningMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningMetadataWithDefaults

`func NewProvisioningMetadataWithDefaults() *ProvisioningMetadata`

NewProvisioningMetadataWithDefaults instantiates a new ProvisioningMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *ProvisioningMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ProvisioningMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ProvisioningMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ProvisioningMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ProvisioningMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProvisioningMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProvisioningMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProvisioningMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ProvisioningMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ProvisioningMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ProvisioningMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ProvisioningMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetCreatedByUserUuid

`func (o *ProvisioningMetadata) GetCreatedByUserUuid() string`

GetCreatedByUserUuid returns the CreatedByUserUuid field if non-nil, zero value otherwise.

### GetCreatedByUserUuidOk

`func (o *ProvisioningMetadata) GetCreatedByUserUuidOk() (*string, bool)`

GetCreatedByUserUuidOk returns a tuple with the CreatedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserUuid

`func (o *ProvisioningMetadata) SetCreatedByUserUuid(v string)`

SetCreatedByUserUuid sets CreatedByUserUuid field to given value.

### HasCreatedByUserUuid

`func (o *ProvisioningMetadata) HasCreatedByUserUuid() bool`

HasCreatedByUserUuid returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *ProvisioningMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ProvisioningMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ProvisioningMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *ProvisioningMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *ProvisioningMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ProvisioningMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ProvisioningMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ProvisioningMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *ProvisioningMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *ProvisioningMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *ProvisioningMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *ProvisioningMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetLastModifiedByUserUuid

`func (o *ProvisioningMetadata) GetLastModifiedByUserUuid() string`

GetLastModifiedByUserUuid returns the LastModifiedByUserUuid field if non-nil, zero value otherwise.

### GetLastModifiedByUserUuidOk

`func (o *ProvisioningMetadata) GetLastModifiedByUserUuidOk() (*string, bool)`

GetLastModifiedByUserUuidOk returns a tuple with the LastModifiedByUserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserUuid

`func (o *ProvisioningMetadata) SetLastModifiedByUserUuid(v string)`

SetLastModifiedByUserUuid sets LastModifiedByUserUuid field to given value.

### HasLastModifiedByUserUuid

`func (o *ProvisioningMetadata) HasLastModifiedByUserUuid() bool`

HasLastModifiedByUserUuid returns a boolean if a field has been set.

### GetState

`func (o *ProvisioningMetadata) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProvisioningMetadata) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProvisioningMetadata) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProvisioningMetadata) HasState() bool`

HasState returns a boolean if a field has been set.


