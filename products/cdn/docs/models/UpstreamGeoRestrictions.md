# UpstreamGeoRestrictions

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**BlockList** | Pointer to **[]string** | Country codes, the format should be based on ISO 3166-1 alpha-2 codes standard. Those codes are used to either blacklist or whitelist countries in geoIPBlock.  | [optional] |
|**AllowList** | Pointer to **[]string** | Country codes, the format should be based on ISO 3166-1 alpha-2 codes standard. Those codes are used to either blacklist or whitelist countries in geoIPBlock.  | [optional] |

## Methods

### NewUpstreamGeoRestrictions

`func NewUpstreamGeoRestrictions() *UpstreamGeoRestrictions`

NewUpstreamGeoRestrictions instantiates a new UpstreamGeoRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamGeoRestrictionsWithDefaults

`func NewUpstreamGeoRestrictionsWithDefaults() *UpstreamGeoRestrictions`

NewUpstreamGeoRestrictionsWithDefaults instantiates a new UpstreamGeoRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockList

`func (o *UpstreamGeoRestrictions) GetBlockList() []string`

GetBlockList returns the BlockList field if non-nil, zero value otherwise.

### GetBlockListOk

`func (o *UpstreamGeoRestrictions) GetBlockListOk() (*[]string, bool)`

GetBlockListOk returns a tuple with the BlockList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockList

`func (o *UpstreamGeoRestrictions) SetBlockList(v []string)`

SetBlockList sets BlockList field to given value.

### HasBlockList

`func (o *UpstreamGeoRestrictions) HasBlockList() bool`

HasBlockList returns a boolean if a field has been set.

### GetAllowList

`func (o *UpstreamGeoRestrictions) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *UpstreamGeoRestrictions) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *UpstreamGeoRestrictions) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *UpstreamGeoRestrictions) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.


