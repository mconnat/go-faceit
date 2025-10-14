# ItemList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewItemList

`func NewItemList() *ItemList`

NewItemList instantiates a new ItemList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemListWithDefaults

`func NewItemListWithDefaults() *ItemList`

NewItemListWithDefaults instantiates a new ItemList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ItemList) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ItemList) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ItemList) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ItemList) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItems

`func (o *ItemList) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ItemList) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ItemList) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *ItemList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStart

`func (o *ItemList) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ItemList) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ItemList) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ItemList) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


