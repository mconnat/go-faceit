# \RankingsApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalRanking**](RankingsApi.md#GetGlobalRanking) | **Get** /rankings/games/{game_id}/regions/{region} | Retrieve global ranking of a game
[**GetPlayerRanking**](RankingsApi.md#GetPlayerRanking) | **Get** /rankings/games/{game_id}/regions/{region}/players/{player_id} | Retrieve user position in the global ranking of a game


# **GetGlobalRanking**
> GlobalRankingList GetGlobalRanking(ctx, gameId, region, optional)
Retrieve global ranking of a game

Retrieve global ranking of a game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 
  **region** | **string**| A region of a game | 
 **optional** | ***RankingsApiGetGlobalRankingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RankingsApiGetGlobalRankingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **country** | **optional.String**| A country code (ISO 3166-1) | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**GlobalRankingList**](GlobalRankingList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerRanking**
> PlayerGlobalRanking GetPlayerRanking(ctx, gameId, region, playerId, optional)
Retrieve user position in the global ranking of a game

Retrieve user position in the global ranking of a game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 
  **region** | **string**| A region of a game | 
  **playerId** | **string**| The id of a player | 
 **optional** | ***RankingsApiGetPlayerRankingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RankingsApiGetPlayerRankingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **country** | **optional.String**| A country code (ISO 3166-1) | 
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**PlayerGlobalRanking**](PlayerGlobalRanking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

