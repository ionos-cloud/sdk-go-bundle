# BackupResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Metadata** | Pointer to [**BackupMetadata**](BackupMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**ClusterBackup**](ClusterBackup.md) |  | [optional] |

## Methods

### NewBackupResponse

`func NewBackupResponse() *BackupResponse`

NewBackupResponse instantiates a new BackupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupResponseWithDefaults

`func NewBackupResponseWithDefaults() *BackupResponse`

NewBackupResponseWithDefaults instantiates a new BackupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackupResponse) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupResponse) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupResponse) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *BackupResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *BackupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *BackupResponse) GetMetadata() BackupMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BackupResponse) GetMetadataOk() (*BackupMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BackupResponse) SetMetadata(v BackupMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BackupResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *BackupResponse) GetProperties() ClusterBackup`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BackupResponse) GetPropertiesOk() (*ClusterBackup, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BackupResponse) SetProperties(v ClusterBackup)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BackupResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


