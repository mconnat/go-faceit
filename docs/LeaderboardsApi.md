# {{classname}}

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChampionshipGroupRanking**](LeaderboardsApi.md#GetChampionshipGroupRanking) | **Get** /leaderboards/championships/{championship_id}/groups/{group} | Retrieve group ranking of a championship
[**GetChampionshipLeaderboards**](LeaderboardsApi.md#GetChampionshipLeaderboards) | **Get** /leaderboards/championships/{championship_id} | Retrieve all leaderboards of a championship
[**GetHubLeaderboards**](LeaderboardsApi.md#GetHubLeaderboards) | **Get** /leaderboards/hubs/{hub_id} | Retrieve all leaderboards of a hub
[**GetHubRanking**](LeaderboardsApi.md#GetHubRanking) | **Get** /leaderboards/hubs/{hub_id}/general | Retrieve all time ranking of a hub
[**GetHubSeasonRanking**](LeaderboardsApi.md#GetHubSeasonRanking) | **Get** /leaderboards/hubs/{hub_id}/seasons/{season} | Retrieve seasonal ranking of a hub
[**GetLeaderboard**](LeaderboardsApi.md#GetLeaderboard) | **Get** /leaderboards/{leaderboard_id} | Retrieve ranking from a leaderboard id
[**GetPlayerRankingInLeaderboard**](LeaderboardsApi.md#GetPlayerRankingInLeaderboard) | **Get** /leaderboards/{leaderboard_id}/players/{player_id} | Retrieve a players ranking in a leaderboard

# **GetChampionshipGroupRanking**
> EntityRanking GetChampionshipGroupRanking(ctx, championshipId, group, optional)
Retrieve group ranking of a championship

Retrieve group ranking of a championship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **championshipId** | **string**| The id of the championship | 
  **group** | **int32**| A group of the championship | 
 **optional** | ***LeaderboardsApiGetChampionshipGroupRankingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeaderboardsApiGetChampionshipGroupRankingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChampionshipLeaderboards**
> LeaderboardsList GetChampionshipLeaderboards(ctx, championshipId, optional)
Retrieve all leaderboards of a championship

Retrieve all leaderboards of a championship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **championshipId** | **string**| The id of the championship | 
 **optional** | ***LeaderboardsApiGetChampionshipLeaderboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeaderboardsApiGetChampionshipLeaderboardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**LeaderboardsList**](LeaderboardsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubLeaderboards**
> LeaderboardsList GetHubLeaderboards(ctx, hubId, optional)
Retrieve all leaderboards of a hub

Retrieve all leaderboards of a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
 **optional** | ***LeaderboardsApiGetHubLeaderboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeaderboardsApiGetHubLeaderboardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**LeaderboardsList**](LeaderboardsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubRanking**
> EntityRanking GetHubRanking(ctx, hubId, optional)
Retrieve all time ranking of a hub

Retrieve all time ranking of a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
 **optional** | ***LeaderboardsApiGetHubRankingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeaderboardsApiGetHubRankingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubSeasonRanking**
> EntityRanking GetHubSeasonRanking(ctx, hubId, season, optional)
Retrieve seasonal ranking of a hub

Retrieve seasonal ranking of a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
  **season** | **int32**| A season of the hub | 
 **optional** | ***LeaderboardsApiGetHubSeasonRankingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeaderboardsApiGetHubSeasonRankingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeaderboard**
> EntityRanking GetLeaderboard(ctx, leaderboardId, optional)
Retrieve ranking from a leaderboard id

Retrieve ranking from a leaderboard id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **leaderboardId** | **string**| The id of the leaderboard | 
 **optional** | ***LeaderboardsApiGetLeaderboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeaderboardsApiGetLeaderboardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**EntityRanking**](EntityRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerRankingInLeaderboard**
> Ranking GetPlayerRankingInLeaderboard(ctx, leaderboardId, playerId)
Retrieve a players ranking in a leaderboard

Retrieve a players ranking in a leaderboard

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **leaderboardId** | **string**| The id of the leaderboard | 
  **playerId** | **string**| The id of the player | 

### Return type

[**Ranking**](Ranking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

