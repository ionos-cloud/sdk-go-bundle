# CommonZoneRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The zone ID (UUID). | |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Metadata** | [**MetadataWithStateNameservers**](MetadataWithStateNameservers.md) |  | |

## Methods

### NewCommonZoneRead

`func NewCommonZoneRead(id string, type_ string, href string, metadata MetadataWithStateNameservers, ) *CommonZoneRead`

NewCommonZoneRead instantiates a new CommonZoneRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonZoneReadWithDefaults

`func NewCommonZoneReadWithDefaults() *CommonZoneRead`

NewCommonZoneReadWithDefaults instantiates a new CommonZoneRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonZoneRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonZoneRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonZoneRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CommonZoneRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommonZoneRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommonZoneRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *CommonZoneRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CommonZoneRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CommonZoneRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *CommonZoneRead) GetMetadata() MetadataWithStateNameservers`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CommonZoneRead) GetMetadataOk() (*MetadataWithStateNameservers, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CommonZoneRead) SetMetadata(v MetadataWithStateNameservers)`

SetMetadata sets Metadata field to given value.



