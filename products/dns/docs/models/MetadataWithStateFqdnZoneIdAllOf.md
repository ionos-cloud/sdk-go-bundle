# MetadataWithStateFqdnZoneIdAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**State** | [**ProvisioningState**](ProvisioningState.md) |  | |
|**Fqdn** | **string** | A fully qualified domain name. FQDN consists of two parts - the hostname and the domain name. | [readonly] |
|**ZoneId** | **string** | The ID (UUID) of the DNS zone of which record belongs to. | [readonly] |

## Methods

### NewMetadataWithStateFqdnZoneIdAllOf

`func NewMetadataWithStateFqdnZoneIdAllOf(state ProvisioningState, fqdn string, zoneId string, ) *MetadataWithStateFqdnZoneIdAllOf`

NewMetadataWithStateFqdnZoneIdAllOf instantiates a new MetadataWithStateFqdnZoneIdAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStateFqdnZoneIdAllOfWithDefaults

`func NewMetadataWithStateFqdnZoneIdAllOfWithDefaults() *MetadataWithStateFqdnZoneIdAllOf`

NewMetadataWithStateFqdnZoneIdAllOfWithDefaults instantiates a new MetadataWithStateFqdnZoneIdAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MetadataWithStateFqdnZoneIdAllOf) GetState() ProvisioningState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithStateFqdnZoneIdAllOf) GetStateOk() (*ProvisioningState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithStateFqdnZoneIdAllOf) SetState(v ProvisioningState)`

SetState sets State field to given value.


### GetFqdn

`func (o *MetadataWithStateFqdnZoneIdAllOf) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *MetadataWithStateFqdnZoneIdAllOf) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *MetadataWithStateFqdnZoneIdAllOf) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetZoneId

`func (o *MetadataWithStateFqdnZoneIdAllOf) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *MetadataWithStateFqdnZoneIdAllOf) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *MetadataWithStateFqdnZoneIdAllOf) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.



