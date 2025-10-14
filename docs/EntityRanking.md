# EntityRanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]Ranking**](Ranking.md) |  | [optional] 
**Leaderboard** | Pointer to [**Leaderboard**](Leaderboard.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewEntityRanking

`func NewEntityRanking() *EntityRanking`

NewEntityRanking instantiates a new EntityRanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRankingWithDefaults

`func NewEntityRankingWithDefaults() *EntityRanking`

NewEntityRankingWithDefaults instantiates a new EntityRanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *EntityRanking) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EntityRanking) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EntityRanking) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EntityRanking) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItems

`func (o *EntityRanking) GetItems() []Ranking`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EntityRanking) GetItemsOk() (*[]Ranking, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EntityRanking) SetItems(v []Ranking)`

SetItems sets Items field to given value.

### HasItems

`func (o *EntityRanking) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLeaderboard

`func (o *EntityRanking) GetLeaderboard() Leaderboard`

GetLeaderboard returns the Leaderboard field if non-nil, zero value otherwise.

### GetLeaderboardOk

`func (o *EntityRanking) GetLeaderboardOk() (*Leaderboard, bool)`

GetLeaderboardOk returns a tuple with the Leaderboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboard

`func (o *EntityRanking) SetLeaderboard(v Leaderboard)`

SetLeaderboard sets Leaderboard field to given value.

### HasLeaderboard

`func (o *EntityRanking) HasLeaderboard() bool`

HasLeaderboard returns a boolean if a field has been set.

### GetStart

`func (o *EntityRanking) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EntityRanking) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EntityRanking) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *EntityRanking) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


