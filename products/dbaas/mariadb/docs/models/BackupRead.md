# BackupRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Backup. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Backup. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**Backup**](Backup.md) |  | |

## Methods

### NewBackupRead

`func NewBackupRead(id string, type_ string, href string, metadata map[string]interface{}, properties Backup, ) *BackupRead`

NewBackupRead instantiates a new BackupRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupReadWithDefaults

`func NewBackupReadWithDefaults() *BackupRead`

NewBackupReadWithDefaults instantiates a new BackupRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BackupRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *BackupRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *BackupRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BackupRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BackupRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *BackupRead) GetProperties() Backup`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BackupRead) GetPropertiesOk() (*Backup, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BackupRead) SetProperties(v Backup)`

SetProperties sets Properties field to given value.



