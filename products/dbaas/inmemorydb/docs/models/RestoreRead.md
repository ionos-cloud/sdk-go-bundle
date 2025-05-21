# RestoreRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Restore. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Restore. | |
|**Metadata** | [**RestoreMetadata**](RestoreMetadata.md) |  | |
|**Properties** | [**Restore**](Restore.md) |  | |

## Methods

### NewRestoreRead

`func NewRestoreRead(id string, type_ string, href string, metadata RestoreMetadata, properties Restore, ) *RestoreRead`

NewRestoreRead instantiates a new RestoreRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreReadWithDefaults

`func NewRestoreReadWithDefaults() *RestoreRead`

NewRestoreReadWithDefaults instantiates a new RestoreRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestoreRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RestoreRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RestoreRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RestoreRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RestoreRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *RestoreRead) GetMetadata() RestoreMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RestoreRead) GetMetadataOk() (*RestoreMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RestoreRead) SetMetadata(v RestoreMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *RestoreRead) GetProperties() Restore`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RestoreRead) GetPropertiesOk() (*Restore, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RestoreRead) SetProperties(v Restore)`

SetProperties sets Properties field to given value.



