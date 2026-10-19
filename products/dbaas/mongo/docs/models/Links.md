# Links

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Prev** | Pointer to **string** | URL (with offset and limit parameters) of the previous page; only present if offset is greater than 0.  | [optional] [readonly] |
|**Self** | Pointer to **string** | URL (with offset and limit parameters) of the current page.  | [optional] [readonly] |
|**Next** | Pointer to **string** | URL (with offset and limit parameters) of the next page; only present if offset + limit is less than the total number of elements.  | [optional] [readonly] |

## Methods

### NewLinks

`func NewLinks() *Links`

NewLinks instantiates a new Links object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksWithDefaults

`func NewLinksWithDefaults() *Links`

NewLinksWithDefaults instantiates a new Links object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrev

`func (o *Links) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *Links) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *Links) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *Links) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetSelf

`func (o *Links) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Links) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Links) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Links) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetNext

`func (o *Links) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Links) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Links) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Links) HasNext() bool`

HasNext returns a boolean if a field has been set.


