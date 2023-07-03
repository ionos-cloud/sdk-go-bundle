# RecordRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The record ID (UUID). | [readonly] |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Metadata** | [**MetadataWithStateFqdnZoneId**](MetadataWithStateFqdnZoneId.md) |  | |
|**Properties** | [**Record**](Record.md) |  | |

## Methods

### NewRecordRead

`func NewRecordRead(id string, type_ string, href string, metadata MetadataWithStateFqdnZoneId, properties Record, ) *RecordRead`

NewRecordRead instantiates a new RecordRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordReadWithDefaults

`func NewRecordReadWithDefaults() *RecordRead`

NewRecordReadWithDefaults instantiates a new RecordRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecordRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RecordRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RecordRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RecordRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RecordRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *RecordRead) GetMetadata() MetadataWithStateFqdnZoneId`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecordRead) GetMetadataOk() (*MetadataWithStateFqdnZoneId, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecordRead) SetMetadata(v MetadataWithStateFqdnZoneId)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *RecordRead) GetProperties() Record`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RecordRead) GetPropertiesOk() (*Record, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RecordRead) SetProperties(v Record)`

SetProperties sets Properties field to given value.



