# ClusterBackupListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]BackupResponse**](BackupResponse.md) |  | [optional] |

## Methods

### NewClusterBackupListAllOf

`func NewClusterBackupListAllOf() *ClusterBackupListAllOf`

NewClusterBackupListAllOf instantiates a new ClusterBackupListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterBackupListAllOfWithDefaults

`func NewClusterBackupListAllOfWithDefaults() *ClusterBackupListAllOf`

NewClusterBackupListAllOfWithDefaults instantiates a new ClusterBackupListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterBackupListAllOf) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterBackupListAllOf) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterBackupListAllOf) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterBackupListAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ClusterBackupListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterBackupListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterBackupListAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterBackupListAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *ClusterBackupListAllOf) GetItems() []BackupResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterBackupListAllOf) GetItemsOk() (*[]BackupResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterBackupListAllOf) SetItems(v []BackupResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterBackupListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


