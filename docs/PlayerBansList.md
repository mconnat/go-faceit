# PlayerBansList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]PlayerBan**](PlayerBan.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewPlayerBansList

`func NewPlayerBansList() *PlayerBansList`

NewPlayerBansList instantiates a new PlayerBansList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerBansListWithDefaults

`func NewPlayerBansListWithDefaults() *PlayerBansList`

NewPlayerBansListWithDefaults instantiates a new PlayerBansList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *PlayerBansList) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PlayerBansList) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PlayerBansList) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PlayerBansList) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItems

`func (o *PlayerBansList) GetItems() []PlayerBan`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PlayerBansList) GetItemsOk() (*[]PlayerBan, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PlayerBansList) SetItems(v []PlayerBan)`

SetItems sets Items field to given value.

### HasItems

`func (o *PlayerBansList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStart

`func (o *PlayerBansList) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PlayerBansList) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PlayerBansList) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *PlayerBansList) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


