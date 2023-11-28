# PaginationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** | Current page. | 
**NextPage** | Pointer to **NullableInt32** | Next page. | [optional] 
**PrevPage** | Pointer to **NullableInt32** | Previous page. | [optional] 
**TotalPages** | **int32** | Total number of pages. | 
**TotalCount** | **int32** | Total number of records. | 

## Methods

### NewPaginationMeta

`func NewPaginationMeta(currentPage int32, totalPages int32, totalCount int32, ) *PaginationMeta`

NewPaginationMeta instantiates a new PaginationMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationMetaWithDefaults

`func NewPaginationMetaWithDefaults() *PaginationMeta`

NewPaginationMetaWithDefaults instantiates a new PaginationMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PaginationMeta) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PaginationMeta) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PaginationMeta) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetNextPage

`func (o *PaginationMeta) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PaginationMeta) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PaginationMeta) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PaginationMeta) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *PaginationMeta) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *PaginationMeta) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil
### GetPrevPage

`func (o *PaginationMeta) GetPrevPage() int32`

GetPrevPage returns the PrevPage field if non-nil, zero value otherwise.

### GetPrevPageOk

`func (o *PaginationMeta) GetPrevPageOk() (*int32, bool)`

GetPrevPageOk returns a tuple with the PrevPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPage

`func (o *PaginationMeta) SetPrevPage(v int32)`

SetPrevPage sets PrevPage field to given value.

### HasPrevPage

`func (o *PaginationMeta) HasPrevPage() bool`

HasPrevPage returns a boolean if a field has been set.

### SetPrevPageNil

`func (o *PaginationMeta) SetPrevPageNil(b bool)`

 SetPrevPageNil sets the value for PrevPage to be an explicit nil

### UnsetPrevPage
`func (o *PaginationMeta) UnsetPrevPage()`

UnsetPrevPage ensures that no value is present for PrevPage, not even an explicit nil
### GetTotalPages

`func (o *PaginationMeta) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginationMeta) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginationMeta) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetTotalCount

`func (o *PaginationMeta) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginationMeta) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginationMeta) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


