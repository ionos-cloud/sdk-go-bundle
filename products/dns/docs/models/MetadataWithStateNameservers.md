# MetadataWithStateNameservers

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
|**Nameservers** | **[]string** | The list of nameservers associated to the zone.  Nameservers are different for primary and secondary zones. For primary zones it would be: - ns-ic.ui-dns.com - ns-ic.ui-dns.de - ns-ic.ui-dns.org - ns-ic.ui-dns.biz  And for secondary zones: - nscs.ui-dns.com - nscs.ui-dns.de - nscs.ui-dns.org - nscs.ui-dns.biz  | |

## Methods

### NewMetadataWithStateNameservers

`func NewMetadataWithStateNameservers(state ProvisioningState, nameservers []string, ) *MetadataWithStateNameservers`

NewMetadataWithStateNameservers instantiates a new MetadataWithStateNameservers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStateNameserversWithDefaults

`func NewMetadataWithStateNameserversWithDefaults() *MetadataWithStateNameservers`

NewMetadataWithStateNameserversWithDefaults instantiates a new MetadataWithStateNameservers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *MetadataWithStateNameservers) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *MetadataWithStateNameservers) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *MetadataWithStateNameservers) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *MetadataWithStateNameservers) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MetadataWithStateNameservers) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MetadataWithStateNameservers) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MetadataWithStateNameservers) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MetadataWithStateNameservers) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *MetadataWithStateNameservers) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *MetadataWithStateNameservers) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *MetadataWithStateNameservers) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *MetadataWithStateNameservers) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *MetadataWithStateNameservers) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *MetadataWithStateNameservers) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *MetadataWithStateNameservers) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *MetadataWithStateNameservers) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *MetadataWithStateNameservers) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MetadataWithStateNameservers) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MetadataWithStateNameservers) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MetadataWithStateNameservers) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *MetadataWithStateNameservers) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *MetadataWithStateNameservers) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *MetadataWithStateNameservers) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *MetadataWithStateNameservers) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetState

`func (o *MetadataWithStateNameservers) GetState() ProvisioningState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithStateNameservers) GetStateOk() (*ProvisioningState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithStateNameservers) SetState(v ProvisioningState)`

SetState sets State field to given value.


### GetNameservers

`func (o *MetadataWithStateNameservers) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *MetadataWithStateNameservers) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *MetadataWithStateNameservers) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.



