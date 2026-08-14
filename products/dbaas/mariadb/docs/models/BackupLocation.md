# BackupLocation

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Location** | Pointer to **string** | The Object Storage location where the backup will be created. The BackupLocations provides a list of supported locations.  | [optional] |

## Methods

### NewBackupLocation

`func NewBackupLocation() *BackupLocation`

NewBackupLocation instantiates a new BackupLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupLocationWithDefaults

`func NewBackupLocationWithDefaults() *BackupLocation`

NewBackupLocationWithDefaults instantiates a new BackupLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *BackupLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BackupLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BackupLocation) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BackupLocation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


