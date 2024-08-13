# SecondaryZoneRecordReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The resource&#39;s unique identifier. | [readonly] |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Metadata** | [**SecondaryZoneRecordReadListMetadata**](SecondaryZoneRecordReadListMetadata.md) |  | |
|**Items** | [**[]SecondaryZoneRecordRead**](SecondaryZoneRecordRead.md) |  | |
|**Offset** | **float32** | Pagination offset. | [readonly] |
|**Limit** | **float32** | Pagination limit. | [readonly] |
|**Links** | [**Links**](Links.md) |  | |

## Methods

### NewSecondaryZoneRecordReadList

`func NewSecondaryZoneRecordReadList(id string, type_ string, href string, metadata SecondaryZoneRecordReadListMetadata, items []SecondaryZoneRecordRead, offset float32, limit float32, links Links, ) *SecondaryZoneRecordReadList`

NewSecondaryZoneRecordReadList instantiates a new SecondaryZoneRecordReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryZoneRecordReadListWithDefaults

`func NewSecondaryZoneRecordReadListWithDefaults() *SecondaryZoneRecordReadList`

NewSecondaryZoneRecordReadListWithDefaults instantiates a new SecondaryZoneRecordReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecondaryZoneRecordReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecondaryZoneRecordReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecondaryZoneRecordReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SecondaryZoneRecordReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecondaryZoneRecordReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecondaryZoneRecordReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SecondaryZoneRecordReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SecondaryZoneRecordReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SecondaryZoneRecordReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *SecondaryZoneRecordReadList) GetMetadata() SecondaryZoneRecordReadListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SecondaryZoneRecordReadList) GetMetadataOk() (*SecondaryZoneRecordReadListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SecondaryZoneRecordReadList) SetMetadata(v SecondaryZoneRecordReadListMetadata)`

SetMetadata sets Metadata field to given value.


### GetItems

`func (o *SecondaryZoneRecordReadList) GetItems() []SecondaryZoneRecordRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecondaryZoneRecordReadList) GetItemsOk() (*[]SecondaryZoneRecordRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecondaryZoneRecordReadList) SetItems(v []SecondaryZoneRecordRead)`

SetItems sets Items field to given value.


### GetOffset

`func (o *SecondaryZoneRecordReadList) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SecondaryZoneRecordReadList) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SecondaryZoneRecordReadList) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *SecondaryZoneRecordReadList) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SecondaryZoneRecordReadList) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SecondaryZoneRecordReadList) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *SecondaryZoneRecordReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecondaryZoneRecordReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecondaryZoneRecordReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.



