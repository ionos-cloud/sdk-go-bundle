# PaginationLinks

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Next** | **string** |  | |
|**Previous** | **string** |  | |
|**Self** | **string** |  | |

## Methods

### NewPaginationLinks

`func NewPaginationLinks(next string, previous string, self string, ) *PaginationLinks`

NewPaginationLinks instantiates a new PaginationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationLinksWithDefaults

`func NewPaginationLinksWithDefaults() *PaginationLinks`

NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginationLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationLinks) SetNext(v string)`

SetNext sets Next field to given value.


### GetPrevious

`func (o *PaginationLinks) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginationLinks) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginationLinks) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### GetSelf

`func (o *PaginationLinks) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PaginationLinks) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PaginationLinks) SetSelf(v string)`

SetSelf sets Self field to given value.



