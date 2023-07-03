# MetadataWithStateFqdnZoneId

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The date of the last change formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The date of creation of the zone formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
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



