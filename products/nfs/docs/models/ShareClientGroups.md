# ShareClientGroups

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Description** | Pointer to **string** | Optional description for the clients groups.  | [optional] |
|**IpNetworks** | Pointer to **[]string** |  | [optional] |
|**Hosts** | Pointer to **[]string** |  | [optional] |
|**Nfs** | Pointer to [**ShareClientGroupsNfs**](ShareClientGroupsNfs.md) |  | [optional] |

## Methods

### NewShareClientGroups

`func NewShareClientGroups() *ShareClientGroups`

NewShareClientGroups instantiates a new ShareClientGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareClientGroupsWithDefaults

`func NewShareClientGroupsWithDefaults() *ShareClientGroups`

NewShareClientGroupsWithDefaults instantiates a new ShareClientGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ShareClientGroups) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShareClientGroups) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShareClientGroups) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShareClientGroups) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpNetworks

`func (o *ShareClientGroups) GetIpNetworks() []string`

GetIpNetworks returns the IpNetworks field if non-nil, zero value otherwise.

### GetIpNetworksOk

`func (o *ShareClientGroups) GetIpNetworksOk() (*[]string, bool)`

GetIpNetworksOk returns a tuple with the IpNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpNetworks

`func (o *ShareClientGroups) SetIpNetworks(v []string)`

SetIpNetworks sets IpNetworks field to given value.

### HasIpNetworks

`func (o *ShareClientGroups) HasIpNetworks() bool`

HasIpNetworks returns a boolean if a field has been set.

### GetHosts

`func (o *ShareClientGroups) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ShareClientGroups) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ShareClientGroups) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ShareClientGroups) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNfs

`func (o *ShareClientGroups) GetNfs() ShareClientGroupsNfs`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *ShareClientGroups) GetNfsOk() (*ShareClientGroupsNfs, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *ShareClientGroups) SetNfs(v ShareClientGroupsNfs)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *ShareClientGroups) HasNfs() bool`

HasNfs returns a boolean if a field has been set.


