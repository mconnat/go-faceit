# LeaderboardConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPlayers** | Pointer to **int64** | Max players in the leaderboard. | [optional] 
**PointsPerLoss** | Pointer to **int64** | User will lose this amount of points if they lose a match | [optional] 
**PointsPerWin** | Pointer to **int64** | User will gain this amount of points if they win a match. When not configured, it&#39;s using the global value which is 3 | [optional] 
**Promotion** | Pointer to [**Promotion**](Promotion.md) |  | [optional] 
**Relegation** | Pointer to [**Relegation**](Relegation.md) |  | [optional] 
**StartingPoints** | Pointer to **int64** | Starting points for a player. | [optional] 

## Methods

### NewLeaderboardConfig

`func NewLeaderboardConfig() *LeaderboardConfig`

NewLeaderboardConfig instantiates a new LeaderboardConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardConfigWithDefaults

`func NewLeaderboardConfigWithDefaults() *LeaderboardConfig`

NewLeaderboardConfigWithDefaults instantiates a new LeaderboardConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPlayers

`func (o *LeaderboardConfig) GetMaxPlayers() int64`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *LeaderboardConfig) GetMaxPlayersOk() (*int64, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *LeaderboardConfig) SetMaxPlayers(v int64)`

SetMaxPlayers sets MaxPlayers field to given value.

### HasMaxPlayers

`func (o *LeaderboardConfig) HasMaxPlayers() bool`

HasMaxPlayers returns a boolean if a field has been set.

### GetPointsPerLoss

`func (o *LeaderboardConfig) GetPointsPerLoss() int64`

GetPointsPerLoss returns the PointsPerLoss field if non-nil, zero value otherwise.

### GetPointsPerLossOk

`func (o *LeaderboardConfig) GetPointsPerLossOk() (*int64, bool)`

GetPointsPerLossOk returns a tuple with the PointsPerLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsPerLoss

`func (o *LeaderboardConfig) SetPointsPerLoss(v int64)`

SetPointsPerLoss sets PointsPerLoss field to given value.

### HasPointsPerLoss

`func (o *LeaderboardConfig) HasPointsPerLoss() bool`

HasPointsPerLoss returns a boolean if a field has been set.

### GetPointsPerWin

`func (o *LeaderboardConfig) GetPointsPerWin() int64`

GetPointsPerWin returns the PointsPerWin field if non-nil, zero value otherwise.

### GetPointsPerWinOk

`func (o *LeaderboardConfig) GetPointsPerWinOk() (*int64, bool)`

GetPointsPerWinOk returns a tuple with the PointsPerWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsPerWin

`func (o *LeaderboardConfig) SetPointsPerWin(v int64)`

SetPointsPerWin sets PointsPerWin field to given value.

### HasPointsPerWin

`func (o *LeaderboardConfig) HasPointsPerWin() bool`

HasPointsPerWin returns a boolean if a field has been set.

### GetPromotion

`func (o *LeaderboardConfig) GetPromotion() Promotion`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *LeaderboardConfig) GetPromotionOk() (*Promotion, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *LeaderboardConfig) SetPromotion(v Promotion)`

SetPromotion sets Promotion field to given value.

### HasPromotion

`func (o *LeaderboardConfig) HasPromotion() bool`

HasPromotion returns a boolean if a field has been set.

### GetRelegation

`func (o *LeaderboardConfig) GetRelegation() Relegation`

GetRelegation returns the Relegation field if non-nil, zero value otherwise.

### GetRelegationOk

`func (o *LeaderboardConfig) GetRelegationOk() (*Relegation, bool)`

GetRelegationOk returns a tuple with the Relegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelegation

`func (o *LeaderboardConfig) SetRelegation(v Relegation)`

SetRelegation sets Relegation field to given value.

### HasRelegation

`func (o *LeaderboardConfig) HasRelegation() bool`

HasRelegation returns a boolean if a field has been set.

### GetStartingPoints

`func (o *LeaderboardConfig) GetStartingPoints() int64`

GetStartingPoints returns the StartingPoints field if non-nil, zero value otherwise.

### GetStartingPointsOk

`func (o *LeaderboardConfig) GetStartingPointsOk() (*int64, bool)`

GetStartingPointsOk returns a tuple with the StartingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingPoints

`func (o *LeaderboardConfig) SetStartingPoints(v int64)`

SetStartingPoints sets StartingPoints field to given value.

### HasStartingPoints

`func (o *LeaderboardConfig) HasStartingPoints() bool`

HasStartingPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


