# BackupProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Location** | Pointer to **string** | The location where the cluster backups will be stored. If not set, the backup is stored in the nearest location of the cluster.  | [optional] |

## Methods

### NewBackupProperties

`func NewBackupProperties() *BackupProperties`

NewBackupProperties instantiates a new BackupProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPropertiesWithDefaults

`func NewBackupPropertiesWithDefaults() *BackupProperties`

NewBackupPropertiesWithDefaults instantiates a new BackupProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *BackupProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BackupProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BackupProperties) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BackupProperties) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


