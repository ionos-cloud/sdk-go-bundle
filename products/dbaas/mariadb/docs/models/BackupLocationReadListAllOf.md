# BackupLocationReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of BackupLocation resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of BackupLocation resources. | |
|**Items** | Pointer to [**[]BackupLocationRead**](BackupLocationRead.md) | The list of BackupLocation resources. | [optional] |

## Methods

### NewBackupLocationReadListAllOf

`func NewBackupLocationReadListAllOf(id string, type_ string, href string, ) *BackupLocationReadListAllOf`

NewBackupLocationReadListAllOf instantiates a new BackupLocationReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupLocationReadListAllOfWithDefaults

`func NewBackupLocationReadListAllOfWithDefaults() *BackupLocationReadListAllOf`

NewBackupLocationReadListAllOfWithDefaults instantiates a new BackupLocationReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupLocationReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupLocationReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupLocationReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BackupLocationReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupLocationReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupLocationReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *BackupLocationReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BackupLocationReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BackupLocationReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *BackupLocationReadListAllOf) GetItems() []BackupLocationRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BackupLocationReadListAllOf) GetItemsOk() (*[]BackupLocationRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BackupLocationReadListAllOf) SetItems(v []BackupLocationRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *BackupLocationReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


