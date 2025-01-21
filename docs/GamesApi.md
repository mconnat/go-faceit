# \GamesApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllGames**](GamesApi.md#GetAllGames) | **Get** /games | Retrieve details of all games on FACEIT
[**GetGame**](GamesApi.md#GetGame) | **Get** /games/{game_id} | Retrieve game details
[**GetGameMatchmakings**](GamesApi.md#GetGameMatchmakings) | **Get** /games/{gameId}/matchmakings | Retrieve details of all matchmakings of a game on FACEIT
[**GetParentGame**](GamesApi.md#GetParentGame) | **Get** /games/{game_id}/parent | Retrieve the details of the parent game, if the game is region-specific
[**GetQueueBans**](GamesApi.md#GetQueueBans) | **Get** /games/{game_id}/queues/{queue_id}/bans | Retrieve queue bans on FACEIT
[**GetQueueById**](GamesApi.md#GetQueueById) | **Get** /games/{game_id}/queues/{queue_id} | Retrieve details of a queue on FACEIT
[**GetQueuesByEntityFilters**](GamesApi.md#GetQueuesByEntityFilters) | **Get** /games/{game_id}/queues | Retrieve queues by filters on FACEIT
[**GetQueuesByRegion**](GamesApi.md#GetQueuesByRegion) | **Get** /games/{game_id}/regions/{region_id}/queues | Retrieve queues by region on FACEIT


# **GetAllGames**
> GamesList GetAllGames(ctx, optional)
Retrieve details of all games on FACEIT

Retrieve details of all games on FACEIT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GamesApiGetAllGamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetAllGamesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**GamesList**](GamesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGame**
> Game GetGame(ctx, gameId)
Retrieve game details

Retrieve game details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 

### Return type

[**Game**](Game.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGameMatchmakings**
> MatchmakingList GetGameMatchmakings(ctx, gameId, optional)
Retrieve details of all matchmakings of a game on FACEIT

Retrieve details of all matchmakings of a game on FACEIT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 
 **optional** | ***GamesApiGetGameMatchmakingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetGameMatchmakingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **optional.String**| The region of the matchmakings | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**MatchmakingList**](MatchmakingList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParentGame**
> Game GetParentGame(ctx, gameId)
Retrieve the details of the parent game, if the game is region-specific

Retrieve the details of the parent game, if the game is region-specific

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 

### Return type

[**Game**](Game.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueueBans**
> QueueBansList GetQueueBans(ctx, gameId, queueId, optional)
Retrieve queue bans on FACEIT

Retrieve queue bans on FACEIT. Available only for game or queue owners(organizers)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 
  **queueId** | **string**| The id of the queue | 
 **optional** | ***GamesApiGetQueueBansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetQueueBansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**QueueBansList**](QueueBansList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueueById**
> Queue GetQueueById(ctx, gameId, queueId)
Retrieve details of a queue on FACEIT

Retrieve details of a queue on FACEIT. Available only for game or queue owners(organizers)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 
  **queueId** | **string**| The id of the queue | 

### Return type

[**Queue**](Queue.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueuesByEntityFilters**
> QueuesList GetQueuesByEntityFilters(ctx, gameId, entityType, entityId, optional)
Retrieve queues by filters on FACEIT

Retrieve queues by filters on FACEIT. Available only for game developers and queue owners(organizers)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 
  **entityType** | **string**| The type of the entity | 
  **entityId** | **string**| The id of the entity | 
 **optional** | ***GamesApiGetQueuesByEntityFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetQueuesByEntityFiltersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**QueuesList**](QueuesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQueuesByRegion**
> QueuesList GetQueuesByRegion(ctx, gameId, regionId, optional)
Retrieve queues by region on FACEIT

Retrieve queues by region on FACEIT. Available only for game developers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **string**| The id of the game | 
  **regionId** | **string**| The id of the region | 
 **optional** | ***GamesApiGetQueuesByRegionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetQueuesByRegionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**QueuesList**](QueuesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

