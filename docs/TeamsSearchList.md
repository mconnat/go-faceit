# TeamsSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]TeamSearch**](TeamSearch.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewTeamsSearchList

`func NewTeamsSearchList() *TeamsSearchList`

NewTeamsSearchList instantiates a new TeamsSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsSearchListWithDefaults

`func NewTeamsSearchListWithDefaults() *TeamsSearchList`

NewTeamsSearchListWithDefaults instantiates a new TeamsSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *TeamsSearchList) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TeamsSearchList) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TeamsSearchList) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TeamsSearchList) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItems

`func (o *TeamsSearchList) GetItems() []TeamSearch`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TeamsSearchList) GetItemsOk() (*[]TeamSearch, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TeamsSearchList) SetItems(v []TeamSearch)`

SetItems sets Items field to given value.

### HasItems

`func (o *TeamsSearchList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStart

`func (o *TeamsSearchList) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TeamsSearchList) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TeamsSearchList) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *TeamsSearchList) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


