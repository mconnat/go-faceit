# LeaderboardConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPlayers** | **int64** | Max players in the leaderboard. | [optional] [default to null]
**PointsPerLoss** | **int64** | User will lose this amount of points if they lose a match | [optional] [default to null]
**PointsPerWin** | **int64** | User will gain this amount of points if they win a match. When not configured, it&#39;s using the global value which is 3 | [optional] [default to null]
**Promotion** | [***Promotion**](Promotion.md) |  | [optional] [default to null]
**Relegation** | [***Relegation**](Relegation.md) |  | [optional] [default to null]
**StartingPoints** | **int64** | Starting points for a player. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


