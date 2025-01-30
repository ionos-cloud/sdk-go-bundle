# BackupListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]BackupResponse**](BackupResponse.md) |  | [optional] |

## Methods

### NewBackupListAllOf

`func NewBackupListAllOf() *BackupListAllOf`

NewBackupListAllOf instantiates a new BackupListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupListAllOfWithDefaults

`func NewBackupListAllOfWithDefaults() *BackupListAllOf`

NewBackupListAllOfWithDefaults instantiates a new BackupListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupListAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupListAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *BackupListAllOf) GetItems() []BackupResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BackupListAllOf) GetItemsOk() (*[]BackupResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BackupListAllOf) SetItems(v []BackupResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *BackupListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


