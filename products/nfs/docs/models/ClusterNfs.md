# ClusterNfs

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MinVersion** | Pointer to **string** | The version of the NFS cluster, that is supported at minimum.  Currently supported version: * &#x60;4.2&#x60; - NFSv4.2  | [optional] [default to "4.2"]|

## Methods

### NewClusterNfs

`func NewClusterNfs() *ClusterNfs`

NewClusterNfs instantiates a new ClusterNfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNfsWithDefaults

`func NewClusterNfsWithDefaults() *ClusterNfs`

NewClusterNfsWithDefaults instantiates a new ClusterNfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinVersion

`func (o *ClusterNfs) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *ClusterNfs) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *ClusterNfs) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *ClusterNfs) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.


