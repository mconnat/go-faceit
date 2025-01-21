# \LeaguesApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLeagueById**](LeaguesApi.md#GetLeagueById) | **Get** /leagues/{league_id} | Retrieve details of a league of a matchmaking on FACEIT
[**GetLeagueSeason**](LeaguesApi.md#GetLeagueSeason) | **Get** /leagues/{league_id}/seasons/{season_id} | Retrieve details of a season of a league on FACEIT
[**GetPlayerForLeagueSeason**](LeaguesApi.md#GetPlayerForLeagueSeason) | **Get** /leagues/{league_id}/seasons/{season_id}/players/{player_id} | Retrieve details of a player for a given league and season on FACEIT


# **GetLeagueById**
> League GetLeagueById(ctx, leagueId)
Retrieve details of a league of a matchmaking on FACEIT

Retrieve details of a league of a matchmaking on FACEIT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **leagueId** | **string**| The id of the league | 

### Return type

[**League**](League.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLeagueSeason**
> SeasonDetailed GetLeagueSeason(ctx, leagueId, seasonId)
Retrieve details of a season of a league on FACEIT

Retrieve details of a season of a league on FACEIT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **leagueId** | **string**| The id of the league | 
  **seasonId** | **int32**| The id of the season | 

### Return type

[**SeasonDetailed**](SeasonDetailed.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlayerForLeagueSeason**
> PlayerInLeague GetPlayerForLeagueSeason(ctx, leagueId, seasonId, playerId)
Retrieve details of a player for a given league and season on FACEIT

Retrieve details of a player for a given league and season on FACEIT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **leagueId** | **string**| The id of the league | 
  **seasonId** | **int32**| The id of the season | 
  **playerId** | **string**| The id of the player | 

### Return type

[**PlayerInLeague**](PlayerInLeague.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

