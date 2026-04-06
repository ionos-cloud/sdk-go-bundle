# Traffic

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**TrafficMetadata**](TrafficMetadata.md) |  | [optional] |
|**TrafficObj** | Pointer to [**TrafficTrafficObj**](TrafficTrafficObj.md) |  | [optional] |
|**TrafficArr** | Pointer to [**[][]interface{}**](array.md) |  | [optional] |
|**Traffic** | Pointer to **[]string** |  | [optional] |

## Methods

### NewTraffic

`func NewTraffic() *Traffic`

NewTraffic instantiates a new Traffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficWithDefaults

`func NewTrafficWithDefaults() *Traffic`

NewTrafficWithDefaults instantiates a new Traffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Traffic) GetMetadata() TrafficMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Traffic) GetMetadataOk() (*TrafficMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Traffic) SetMetadata(v TrafficMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Traffic) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTrafficObj

`func (o *Traffic) GetTrafficObj() TrafficTrafficObj`

GetTrafficObj returns the TrafficObj field if non-nil, zero value otherwise.

### GetTrafficObjOk

`func (o *Traffic) GetTrafficObjOk() (*TrafficTrafficObj, bool)`

GetTrafficObjOk returns a tuple with the TrafficObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficObj

`func (o *Traffic) SetTrafficObj(v TrafficTrafficObj)`

SetTrafficObj sets TrafficObj field to given value.

### HasTrafficObj

`func (o *Traffic) HasTrafficObj() bool`

HasTrafficObj returns a boolean if a field has been set.

### GetTrafficArr

`func (o *Traffic) GetTrafficArr() [][]interface{}`

GetTrafficArr returns the TrafficArr field if non-nil, zero value otherwise.

### GetTrafficArrOk

`func (o *Traffic) GetTrafficArrOk() (*[][]interface{}, bool)`

GetTrafficArrOk returns a tuple with the TrafficArr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficArr

`func (o *Traffic) SetTrafficArr(v [][]interface{})`

SetTrafficArr sets TrafficArr field to given value.

### HasTrafficArr

`func (o *Traffic) HasTrafficArr() bool`

HasTrafficArr returns a boolean if a field has been set.

### GetTraffic

`func (o *Traffic) GetTraffic() []string`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *Traffic) GetTrafficOk() (*[]string, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *Traffic) SetTraffic(v []string)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *Traffic) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


