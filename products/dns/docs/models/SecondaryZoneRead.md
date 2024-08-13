# SecondaryZoneRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The zone ID (UUID). | |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Metadata** | [**MetadataWithStateNameservers**](MetadataWithStateNameservers.md) |  | |
|**Properties** | [**SecondaryZone**](SecondaryZone.md) |  | |

## Methods

### NewSecondaryZoneRead

`func NewSecondaryZoneRead(id string, type_ string, href string, metadata MetadataWithStateNameservers, properties SecondaryZone, ) *SecondaryZoneRead`

NewSecondaryZoneRead instantiates a new SecondaryZoneRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryZoneReadWithDefaults

`func NewSecondaryZoneReadWithDefaults() *SecondaryZoneRead`

NewSecondaryZoneReadWithDefaults instantiates a new SecondaryZoneRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecondaryZoneRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecondaryZoneRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecondaryZoneRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SecondaryZoneRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecondaryZoneRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecondaryZoneRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SecondaryZoneRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SecondaryZoneRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SecondaryZoneRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *SecondaryZoneRead) GetMetadata() MetadataWithStateNameservers`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecondaryZoneRead) GetMetadataOk() (*MetadataWithStateNameservers, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecondaryZoneRead) SetMetadata(v MetadataWithStateNameservers)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *SecondaryZoneRead) GetProperties() SecondaryZone`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SecondaryZoneRead) GetPropertiesOk() (*SecondaryZone, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SecondaryZoneRead) SetProperties(v SecondaryZone)`

SetProperties sets Properties field to given value.



