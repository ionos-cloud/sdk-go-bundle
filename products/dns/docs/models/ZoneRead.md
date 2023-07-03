# ZoneRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The zone ID (UUID). | |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Metadata** | [**MetadataWithStateNameservers**](MetadataWithStateNameservers.md) |  | |
|**Properties** | [**Zone**](Zone.md) |  | |

## Methods

### NewZoneRead

`func NewZoneRead(id string, type_ string, href string, metadata MetadataWithStateNameservers, properties Zone, ) *ZoneRead`

NewZoneRead instantiates a new ZoneRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneReadWithDefaults

`func NewZoneReadWithDefaults() *ZoneRead`

NewZoneReadWithDefaults instantiates a new ZoneRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ZoneRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ZoneRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ZoneRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ZoneRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *ZoneRead) GetMetadata() MetadataWithStateNameservers`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ZoneRead) GetMetadataOk() (*MetadataWithStateNameservers, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ZoneRead) SetMetadata(v MetadataWithStateNameservers)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *ZoneRead) GetProperties() Zone`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ZoneRead) GetPropertiesOk() (*Zone, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ZoneRead) SetProperties(v Zone)`

SetProperties sets Properties field to given value.



