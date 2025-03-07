# Share

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The directory being exported | |
|**Quota** | Pointer to **int32** | The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using &#x60;0&#x60;.  | [optional] [default to 0]|
|**Gid** | Pointer to **int32** | The group ID that will own the exported directory and be used as anongid in squash modes root-anonymous and all-anonymous.  | [optional] [default to 65534]|
|**Uid** | Pointer to **int32** | The user ID that will own the exported directory and be used as anonuid in squash modes root-anonymous and all-anonymous.  | [optional] [default to 65534]|
|**ClientGroups** | [**[]ShareClientGroups**](ShareClientGroups.md) | The groups of clients are the systems connecting to the Network File Storage cluster.  | |

## Methods

### NewShare

`func NewShare(name string, clientGroups []ShareClientGroups, ) *Share`

NewShare instantiates a new Share object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareWithDefaults

`func NewShareWithDefaults() *Share`

NewShareWithDefaults instantiates a new Share object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Share) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Share) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Share) SetName(v string)`

SetName sets Name field to given value.


### GetQuota

`func (o *Share) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Share) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Share) SetQuota(v int32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Share) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetGid

`func (o *Share) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *Share) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *Share) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *Share) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetUid

`func (o *Share) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Share) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Share) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Share) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetClientGroups

`func (o *Share) GetClientGroups() []ShareClientGroups`

GetClientGroups returns the ClientGroups field if non-nil, zero value otherwise.

### GetClientGroupsOk

`func (o *Share) GetClientGroupsOk() (*[]ShareClientGroups, bool)`

GetClientGroupsOk returns a tuple with the ClientGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientGroups

`func (o *Share) SetClientGroups(v []ShareClientGroups)`

SetClientGroups sets ClientGroups field to given value.



