# MetadataWithStateFqdnZoneId

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The creation date formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
|**CreatedBy** | Pointer to **string** | Unique name of the identity that created the resource. | [optional] [readonly] |
|**CreatedByUserId** | Pointer to **string** | The unique ID of the user who created the resource. | [optional] [readonly] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The date of the last change formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
|**LastModifiedBy** | Pointer to **string** | Unique name of the identity that created the resource. | [optional] [readonly] |
|**LastModifiedByUserId** | Pointer to **string** | The unique ID of the user who last modified the resource. | [optional] [readonly] |
|**State** | [**ProvisioningState**](ProvisioningState.md) |  | |
|**Fqdn** | **string** | A fully qualified domain name. FQDN consists of two parts - the hostname and the domain name. | [readonly] |
|**ZoneId** | **string** | The ID (UUID) of the DNS zone of which record belongs to. | [readonly] |

## Methods

### NewMetadataWithStateFqdnZoneId

`func NewMetadataWithStateFqdnZoneId(state ProvisioningState, fqdn string, zoneId string, ) *MetadataWithStateFqdnZoneId`

NewMetadataWithStateFqdnZoneId instantiates a new MetadataWithStateFqdnZoneId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStateFqdnZoneIdWithDefaults

`func NewMetadataWithStateFqdnZoneIdWithDefaults() *MetadataWithStateFqdnZoneId`

NewMetadataWithStateFqdnZoneIdWithDefaults instantiates a new MetadataWithStateFqdnZoneId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *MetadataWithStateFqdnZoneId) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *MetadataWithStateFqdnZoneId) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *MetadataWithStateFqdnZoneId) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *MetadataWithStateFqdnZoneId) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MetadataWithStateFqdnZoneId) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MetadataWithStateFqdnZoneId) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MetadataWithStateFqdnZoneId) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MetadataWithStateFqdnZoneId) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *MetadataWithStateFqdnZoneId) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *MetadataWithStateFqdnZoneId) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *MetadataWithStateFqdnZoneId) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *MetadataWithStateFqdnZoneId) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *MetadataWithStateFqdnZoneId) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *MetadataWithStateFqdnZoneId) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *MetadataWithStateFqdnZoneId) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *MetadataWithStateFqdnZoneId) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *MetadataWithStateFqdnZoneId) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MetadataWithStateFqdnZoneId) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MetadataWithStateFqdnZoneId) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MetadataWithStateFqdnZoneId) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *MetadataWithStateFqdnZoneId) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *MetadataWithStateFqdnZoneId) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *MetadataWithStateFqdnZoneId) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *MetadataWithStateFqdnZoneId) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetState

`func (o *MetadataWithStateFqdnZoneId) GetState() ProvisioningState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithStateFqdnZoneId) GetStateOk() (*ProvisioningState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithStateFqdnZoneId) SetState(v ProvisioningState)`

SetState sets State field to given value.


### GetFqdn

`func (o *MetadataWithStateFqdnZoneId) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *MetadataWithStateFqdnZoneId) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *MetadataWithStateFqdnZoneId) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetZoneId

`func (o *MetadataWithStateFqdnZoneId) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *MetadataWithStateFqdnZoneId) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *MetadataWithStateFqdnZoneId) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.



