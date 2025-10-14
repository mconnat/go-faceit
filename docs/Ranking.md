# Ranking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentStreak** | Pointer to **int64** |  | [optional] 
**Draw** | Pointer to **int64** |  | [optional] 
**Lost** | Pointer to **int64** |  | [optional] 
**Played** | Pointer to **int64** |  | [optional] 
**Player** | Pointer to [**UserSimple**](UserSimple.md) |  | [optional] 
**Points** | Pointer to **int64** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**WinRate** | Pointer to **float64** |  | [optional] 
**Won** | Pointer to **int64** |  | [optional] 

## Methods

### NewRanking

`func NewRanking() *Ranking`

NewRanking instantiates a new Ranking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRankingWithDefaults

`func NewRankingWithDefaults() *Ranking`

NewRankingWithDefaults instantiates a new Ranking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentStreak

`func (o *Ranking) GetCurrentStreak() int64`

GetCurrentStreak returns the CurrentStreak field if non-nil, zero value otherwise.

### GetCurrentStreakOk

`func (o *Ranking) GetCurrentStreakOk() (*int64, bool)`

GetCurrentStreakOk returns a tuple with the CurrentStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStreak

`func (o *Ranking) SetCurrentStreak(v int64)`

SetCurrentStreak sets CurrentStreak field to given value.

### HasCurrentStreak

`func (o *Ranking) HasCurrentStreak() bool`

HasCurrentStreak returns a boolean if a field has been set.

### GetDraw

`func (o *Ranking) GetDraw() int64`

GetDraw returns the Draw field if non-nil, zero value otherwise.

### GetDrawOk

`func (o *Ranking) GetDrawOk() (*int64, bool)`

GetDrawOk returns a tuple with the Draw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraw

`func (o *Ranking) SetDraw(v int64)`

SetDraw sets Draw field to given value.

### HasDraw

`func (o *Ranking) HasDraw() bool`

HasDraw returns a boolean if a field has been set.

### GetLost

`func (o *Ranking) GetLost() int64`

GetLost returns the Lost field if non-nil, zero value otherwise.

### GetLostOk

`func (o *Ranking) GetLostOk() (*int64, bool)`

GetLostOk returns a tuple with the Lost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLost

`func (o *Ranking) SetLost(v int64)`

SetLost sets Lost field to given value.

### HasLost

`func (o *Ranking) HasLost() bool`

HasLost returns a boolean if a field has been set.

### GetPlayed

`func (o *Ranking) GetPlayed() int64`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *Ranking) GetPlayedOk() (*int64, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *Ranking) SetPlayed(v int64)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *Ranking) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetPlayer

`func (o *Ranking) GetPlayer() UserSimple`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *Ranking) GetPlayerOk() (*UserSimple, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *Ranking) SetPlayer(v UserSimple)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *Ranking) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetPoints

`func (o *Ranking) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *Ranking) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *Ranking) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *Ranking) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetPosition

`func (o *Ranking) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Ranking) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Ranking) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Ranking) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetWinRate

`func (o *Ranking) GetWinRate() float64`

GetWinRate returns the WinRate field if non-nil, zero value otherwise.

### GetWinRateOk

`func (o *Ranking) GetWinRateOk() (*float64, bool)`

GetWinRateOk returns a tuple with the WinRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinRate

`func (o *Ranking) SetWinRate(v float64)`

SetWinRate sets WinRate field to given value.

### HasWinRate

`func (o *Ranking) HasWinRate() bool`

HasWinRate returns a boolean if a field has been set.

### GetWon

`func (o *Ranking) GetWon() int64`

GetWon returns the Won field if non-nil, zero value otherwise.

### GetWonOk

`func (o *Ranking) GetWonOk() (*int64, bool)`

GetWonOk returns a tuple with the Won field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWon

`func (o *Ranking) SetWon(v int64)`

SetWon sets Won field to given value.

### HasWon

`func (o *Ranking) HasWon() bool`

HasWon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


