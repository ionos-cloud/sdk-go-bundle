# SnapshotLocation

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Location** | Pointer to **string** | The Object Storage location where snapshots will be stored. For added data safety, use a different location than the cluster. A list of supported locations is provided by the SnapshotLocations endpoint.  | [optional] |

## Methods

### NewSnapshotLocation

`func NewSnapshotLocation() *SnapshotLocation`

NewSnapshotLocation instantiates a new SnapshotLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotLocationWithDefaults

`func NewSnapshotLocationWithDefaults() *SnapshotLocation`

NewSnapshotLocationWithDefaults instantiates a new SnapshotLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SnapshotLocation) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SnapshotLocation) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SnapshotLocation) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SnapshotLocation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


