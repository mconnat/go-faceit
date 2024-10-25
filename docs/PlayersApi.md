# {{classname}}

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlayer**](PlayersApi.md#GetPlayer) | **Get** /players/{player_id} | Retrieve player details
[**GetPlayerBans**](PlayersApi.md#GetPlayerBans) | **Get** /players/{player_id}/bans | Retrieve all bans of a player
[**GetPlayerFromLookup**](PlayersApi.md#GetPlayerFromLookup) | **Get** /players | Retrieve player details
[**GetPlayerHistory**](PlayersApi.md#GetPlayerHistory) | **Get** /players/{player_id}/history | Retrieve all matches of a player
[**GetPlayerHubs**](PlayersApi.md#GetPlayerHubs) | **Get** /players/{player_id}/hubs | Retrieve all hubs of a player
[**GetPlayerStats**](PlayersApi.md#GetPlayerStats) | **Get** /players/{player_id}/games/{game_id}/stats | Retrieve statistics of a player for a given amount of matches
[**GetPlayerStats_0**](PlayersApi.md#GetPlayerStats_0) | **Get** /players/{player_id}/stats/{game_id} | Retrieve statistics of a player
[**GetPlayerTeams**](PlayersApi.md#GetPlayerTeams) | **Get** /players/{player_id}/teams | Retrieve all teams of a player
[**GetPlayerTournaments**](PlayersApi.md#GetPlayerTournaments) | **Get** /players/{player_id}/tournaments | Retrieve all tournaments of a player

# **GetPlayer**
> Player GetPlayer(ctx, playerId)
Retrieve player details

Retrieve player details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 

### Return type

[**Player**](Player.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerBans**
> PlayerBansList GetPlayerBans(ctx, playerId, optional)
Retrieve all bans of a player

Retrieve all bans of a player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 
 **optional** | ***PlayersApiGetPlayerBansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerBansOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**PlayerBansList**](PlayerBansList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerFromLookup**
> Player GetPlayerFromLookup(ctx, optional)
Retrieve player details

Retrieve player details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlayersApiGetPlayerFromLookupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerFromLookupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nickname** | **optional.String**| The nickname of the player on FACEIT | 
 **game** | **optional.String**| A game on FACEIT | 
 **gamePlayerId** | **optional.String**| The ID of a player on game&#x27;s platform | 

### Return type

[**Player**](Player.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerHistory**
> MatchHistoryList GetPlayerHistory(ctx, playerId, game, optional)
Retrieve all matches of a player

Retrieve all matches of a player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 
  **game** | **string**| A game on FACEIT | 
 **optional** | ***PlayersApiGetPlayerHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Int32**| The timestamp (Unix time) as lower bound of the query. 1 month ago if not specified | 
 **to** | **optional.Int32**| The timestamp (Unix time) as higher bound of the query. Current timestamp if not specified | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**MatchHistoryList**](MatchHistoryList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerHubs**
> HubsList GetPlayerHubs(ctx, playerId, optional)
Retrieve all hubs of a player

Retrieve all hubs of a player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 
 **optional** | ***PlayersApiGetPlayerHubsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerHubsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 50]

### Return type

[**HubsList**](HubsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerStats**
> PlayerStatsForMatchesList GetPlayerStats(ctx, playerId, gameId, optional)
Retrieve statistics of a player for a given amount of matches

Retrieve statistics of a player for a given amount of matches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 
  **gameId** | **string**| A game on FACEIT | 
 **optional** | ***PlayersApiGetPlayerStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]
 **from** | **optional.Int32**| Used to filter the dataset by date (minimum). Expected value is date (\&quot;items.stats.Match Finished At\&quot;) in epoch milliseconds.  | 
 **to** | **optional.Int32**| Used to filter the dataset by date (maximum). Expected value is date (\&quot;items.stats.Match Finished At\&quot;) in epoch milliseconds.  | 

### Return type

[**PlayerStatsForMatchesList**](PlayerStatsForMatchesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerStats_0**
> PlayerStats GetPlayerStats_0(ctx, playerId, gameId)
Retrieve statistics of a player

Retrieve statistics of a player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 
  **gameId** | **string**| A game on FACEIT | 

### Return type

[**PlayerStats**](PlayerStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerTeams**
> TeamList GetPlayerTeams(ctx, playerId, optional)
Retrieve all teams of a player

Retrieve all teams of a player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 
 **optional** | ***PlayersApiGetPlayerTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerTeamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**TeamList**](TeamList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerTournaments**
> TournamentsList GetPlayerTournaments(ctx, playerId, optional)
Retrieve all tournaments of a player

Retrieve all tournaments of a player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| The id of the player | 
 **optional** | ***PlayersApiGetPlayerTournamentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayersApiGetPlayerTournamentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**TournamentsList**](TournamentsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

