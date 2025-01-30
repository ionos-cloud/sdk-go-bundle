# BackupResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Properties** | Pointer to [**Backup**](Backup.md) |  | [optional] |

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

### GetProperties

`func (o *BackupResponse) GetProperties() Backup`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BackupResponse) GetPropertiesOk() (*Backup, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BackupResponse) SetProperties(v Backup)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BackupResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


