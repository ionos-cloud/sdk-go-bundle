# DnssecKeyReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**Type** | Pointer to **string** |  | [optional] |
|**Href** | Pointer to **string** |  | [optional] [readonly] |
|**Metadata** | Pointer to [**DnssecKeyReadListMetadata**](DnssecKeyReadListMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**DnssecKeyReadListProperties**](DnssecKeyReadListProperties.md) |  | [optional] |

## Methods

### NewDnssecKeyReadList

`func NewDnssecKeyReadList() *DnssecKeyReadList`

NewDnssecKeyReadList instantiates a new DnssecKeyReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecKeyReadListWithDefaults

`func NewDnssecKeyReadListWithDefaults() *DnssecKeyReadList`

NewDnssecKeyReadListWithDefaults instantiates a new DnssecKeyReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DnssecKeyReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DnssecKeyReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DnssecKeyReadList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DnssecKeyReadList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DnssecKeyReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnssecKeyReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnssecKeyReadList) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DnssecKeyReadList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *DnssecKeyReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DnssecKeyReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DnssecKeyReadList) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DnssecKeyReadList) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *DnssecKeyReadList) GetMetadata() DnssecKeyReadListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DnssecKeyReadList) GetMetadataOk() (*DnssecKeyReadListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DnssecKeyReadList) SetMetadata(v DnssecKeyReadListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DnssecKeyReadList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *DnssecKeyReadList) GetProperties() DnssecKeyReadListProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DnssecKeyReadList) GetPropertiesOk() (*DnssecKeyReadListProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DnssecKeyReadList) SetProperties(v DnssecKeyReadListProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DnssecKeyReadList) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


