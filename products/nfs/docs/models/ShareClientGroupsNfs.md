# ShareClientGroupsNfs

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Squash** | Pointer to **string** | The squash mode for the export. The squash mode can be: * &#x60;none&#x60; - No squash mode. no mapping (no_all_squash,no_root_squash). * &#x60;root-anonymous&#x60; - Map root user to anonymous uid (root_squash,anonuid&#x3D;&lt;uid&gt;,anongid&#x3D;&lt;gid&gt;). * &#x60;all-anonymous&#x60; - Map all users to anonymous uid (all_squash,anonuid&#x3D;&lt;uid&gt;,anongid&#x3D;&lt;gid&gt;).  | [optional] [default to "none"]|

## Methods

### NewShareClientGroupsNfs

`func NewShareClientGroupsNfs() *ShareClientGroupsNfs`

NewShareClientGroupsNfs instantiates a new ShareClientGroupsNfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareClientGroupsNfsWithDefaults

`func NewShareClientGroupsNfsWithDefaults() *ShareClientGroupsNfs`

NewShareClientGroupsNfsWithDefaults instantiates a new ShareClientGroupsNfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSquash

`func (o *ShareClientGroupsNfs) GetSquash() string`

GetSquash returns the Squash field if non-nil, zero value otherwise.

### GetSquashOk

`func (o *ShareClientGroupsNfs) GetSquashOk() (*string, bool)`

GetSquashOk returns a tuple with the Squash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquash

`func (o *ShareClientGroupsNfs) SetSquash(v string)`

SetSquash sets Squash field to given value.

### HasSquash

`func (o *ShareClientGroupsNfs) HasSquash() bool`

HasSquash returns a boolean if a field has been set.


