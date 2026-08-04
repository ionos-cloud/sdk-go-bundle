# BackupLocationRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the BackupLocation. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the BackupLocation. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**BackupLocation**](BackupLocation.md) |  | |

## Methods

### NewBackupLocationRead

`func NewBackupLocationRead(id string, type_ string, href string, metadata map[string]interface{}, properties BackupLocation, ) *BackupLocationRead`

NewBackupLocationRead instantiates a new BackupLocationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupLocationReadWithDefaults

`func NewBackupLocationReadWithDefaults() *BackupLocationRead`

NewBackupLocationReadWithDefaults instantiates a new BackupLocationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupLocationRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupLocationRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupLocationRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BackupLocationRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupLocationRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupLocationRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *BackupLocationRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupLocationRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupLocationRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *BackupLocationRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BackupLocationRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BackupLocationRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *BackupLocationRead) GetProperties() BackupLocation`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BackupLocationRead) GetPropertiesOk() (*BackupLocation, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BackupLocationRead) SetProperties(v BackupLocation)`

SetProperties sets Properties field to given value.



