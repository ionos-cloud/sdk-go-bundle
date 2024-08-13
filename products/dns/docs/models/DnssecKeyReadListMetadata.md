# DnssecKeyReadListMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ZoneId** | Pointer to **string** | The ID (UUID) of the DNS zone of which record belongs to. | [optional] [readonly] |
|**Items** | Pointer to [**[]DnssecKey**](DnssecKey.md) |  | [optional] |

## Methods

### NewDnssecKeyReadListMetadata

`func NewDnssecKeyReadListMetadata() *DnssecKeyReadListMetadata`

NewDnssecKeyReadListMetadata instantiates a new DnssecKeyReadListMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecKeyReadListMetadataWithDefaults

`func NewDnssecKeyReadListMetadataWithDefaults() *DnssecKeyReadListMetadata`

NewDnssecKeyReadListMetadataWithDefaults instantiates a new DnssecKeyReadListMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *DnssecKeyReadListMetadata) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnssecKeyReadListMetadata) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnssecKeyReadListMetadata) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnssecKeyReadListMetadata) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetItems

`func (o *DnssecKeyReadListMetadata) GetItems() []DnssecKey`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DnssecKeyReadListMetadata) GetItemsOk() (*[]DnssecKey, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DnssecKeyReadListMetadata) SetItems(v []DnssecKey)`

SetItems sets Items field to given value.

### HasItems

`func (o *DnssecKeyReadListMetadata) HasItems() bool`

HasItems returns a boolean if a field has been set.


