# ZoneTransferPrimaryIpsStatus

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | **string** |  | |
|**Items** | [**[]ZoneTransferPrimaryIpStatus**](ZoneTransferPrimaryIpStatus.md) |  | |

## Methods

### NewZoneTransferPrimaryIpsStatus

`func NewZoneTransferPrimaryIpsStatus(type_ string, items []ZoneTransferPrimaryIpStatus, ) *ZoneTransferPrimaryIpsStatus`

NewZoneTransferPrimaryIpsStatus instantiates a new ZoneTransferPrimaryIpsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneTransferPrimaryIpsStatusWithDefaults

`func NewZoneTransferPrimaryIpsStatusWithDefaults() *ZoneTransferPrimaryIpsStatus`

NewZoneTransferPrimaryIpsStatusWithDefaults instantiates a new ZoneTransferPrimaryIpsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ZoneTransferPrimaryIpsStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneTransferPrimaryIpsStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneTransferPrimaryIpsStatus) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *ZoneTransferPrimaryIpsStatus) GetItems() []ZoneTransferPrimaryIpStatus`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ZoneTransferPrimaryIpsStatus) GetItemsOk() (*[]ZoneTransferPrimaryIpStatus, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ZoneTransferPrimaryIpsStatus) SetItems(v []ZoneTransferPrimaryIpStatus)`

SetItems sets Items field to given value.



