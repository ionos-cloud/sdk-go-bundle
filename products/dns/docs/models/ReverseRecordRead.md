# ReverseRecordRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The reverse DNS record ID (UUID). | |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Metadata** | [**Metadata**](Metadata.md) |  | |
|**Properties** | [**ReverseRecord**](ReverseRecord.md) |  | |

## Methods

### NewReverseRecordRead

`func NewReverseRecordRead(id string, type_ string, href string, metadata Metadata, properties ReverseRecord, ) *ReverseRecordRead`

NewReverseRecordRead instantiates a new ReverseRecordRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseRecordReadWithDefaults

`func NewReverseRecordReadWithDefaults() *ReverseRecordRead`

NewReverseRecordReadWithDefaults instantiates a new ReverseRecordRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReverseRecordRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReverseRecordRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReverseRecordRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ReverseRecordRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReverseRecordRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReverseRecordRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ReverseRecordRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ReverseRecordRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ReverseRecordRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *ReverseRecordRead) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReverseRecordRead) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReverseRecordRead) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *ReverseRecordRead) GetProperties() ReverseRecord`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReverseRecordRead) GetPropertiesOk() (*ReverseRecord, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReverseRecordRead) SetProperties(v ReverseRecord)`

SetProperties sets Properties field to given value.



