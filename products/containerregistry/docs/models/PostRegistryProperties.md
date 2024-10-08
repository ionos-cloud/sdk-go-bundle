# PostRegistryProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**GarbageCollectionSchedule** | Pointer to [**WeeklySchedule**](WeeklySchedule.md) |  | [optional] |
|**Location** | **string** |  | |
|**Name** | **string** |  | |
|**Features** | Pointer to [**RegistryFeatures**](RegistryFeatures.md) |  | [optional] |
|**ApiSubnetAllowList** | Pointer to **[]string** | Subnets and IPs that are allowed to access the registry API, supports IPv4 and IPv6. Maximum of 25 items may be specified. If no CIDR is given /32 and /128 are assumed for IPv4 and IPv6 respectively. 0.0.0.0/0 can be used to deny all traffic. __Note__: If this list is empty or not set, there are no restrictions.  | [optional] |

## Methods

### NewPostRegistryProperties

`func NewPostRegistryProperties(location string, name string, ) *PostRegistryProperties`

NewPostRegistryProperties instantiates a new PostRegistryProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRegistryPropertiesWithDefaults

`func NewPostRegistryPropertiesWithDefaults() *PostRegistryProperties`

NewPostRegistryPropertiesWithDefaults instantiates a new PostRegistryProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGarbageCollectionSchedule

`func (o *PostRegistryProperties) GetGarbageCollectionSchedule() WeeklySchedule`

GetGarbageCollectionSchedule returns the GarbageCollectionSchedule field if non-nil, zero value otherwise.

### GetGarbageCollectionScheduleOk

`func (o *PostRegistryProperties) GetGarbageCollectionScheduleOk() (*WeeklySchedule, bool)`

GetGarbageCollectionScheduleOk returns a tuple with the GarbageCollectionSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageCollectionSchedule

`func (o *PostRegistryProperties) SetGarbageCollectionSchedule(v WeeklySchedule)`

SetGarbageCollectionSchedule sets GarbageCollectionSchedule field to given value.

### HasGarbageCollectionSchedule

`func (o *PostRegistryProperties) HasGarbageCollectionSchedule() bool`

HasGarbageCollectionSchedule returns a boolean if a field has been set.

### GetLocation

`func (o *PostRegistryProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostRegistryProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostRegistryProperties) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *PostRegistryProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostRegistryProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostRegistryProperties) SetName(v string)`

SetName sets Name field to given value.


### GetFeatures

`func (o *PostRegistryProperties) GetFeatures() RegistryFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PostRegistryProperties) GetFeaturesOk() (*RegistryFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PostRegistryProperties) SetFeatures(v RegistryFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PostRegistryProperties) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetApiSubnetAllowList

`func (o *PostRegistryProperties) GetApiSubnetAllowList() []string`

GetApiSubnetAllowList returns the ApiSubnetAllowList field if non-nil, zero value otherwise.

### GetApiSubnetAllowListOk

`func (o *PostRegistryProperties) GetApiSubnetAllowListOk() (*[]string, bool)`

GetApiSubnetAllowListOk returns a tuple with the ApiSubnetAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSubnetAllowList

`func (o *PostRegistryProperties) SetApiSubnetAllowList(v []string)`

SetApiSubnetAllowList sets ApiSubnetAllowList field to given value.

### HasApiSubnetAllowList

`func (o *PostRegistryProperties) HasApiSubnetAllowList() bool`

HasApiSubnetAllowList returns a boolean if a field has been set.


