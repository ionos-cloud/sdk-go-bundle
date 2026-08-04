# BackupReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Backup resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Backup resources. | |
|**Items** | Pointer to [**[]BackupRead**](BackupRead.md) | The list of Backup resources. | [optional] |

## Methods

### NewBackupReadListAllOf

`func NewBackupReadListAllOf(id string, type_ string, href string, ) *BackupReadListAllOf`

NewBackupReadListAllOf instantiates a new BackupReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupReadListAllOfWithDefaults

`func NewBackupReadListAllOfWithDefaults() *BackupReadListAllOf`

NewBackupReadListAllOfWithDefaults instantiates a new BackupReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BackupReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *BackupReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *BackupReadListAllOf) GetItems() []BackupRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BackupReadListAllOf) GetItemsOk() (*[]BackupRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BackupReadListAllOf) SetItems(v []BackupRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *BackupReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


