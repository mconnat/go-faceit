# \TeamsApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTeam**](TeamsApi.md#GetTeam) | **Get** /teams/{team_id} | Retrieve team details
[**GetTeamStats**](TeamsApi.md#GetTeamStats) | **Get** /teams/{team_id}/stats/{game_id} | Retrieve statistics of a team
[**GetTeamTournaments**](TeamsApi.md#GetTeamTournaments) | **Get** /teams/{team_id}/tournaments | Retrieve tournaments of a team


# **GetTeam**
> Team GetTeam(ctx, teamId)
Retrieve team details

Retrieve team details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**| The id of the team | 

### Return type

[**Team**](Team.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamStats**
> TeamStats GetTeamStats(ctx, teamId, gameId)
Retrieve statistics of a team

Retrieve statistics of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**| The id of the team | 
  **gameId** | **string**| A game on FACEIT | 

### Return type

[**TeamStats**](TeamStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamTournaments**
> TournamentsList GetTeamTournaments(ctx, teamId, optional)
Retrieve tournaments of a team

Retrieve tournaments of a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**| The id of the team | 
 **optional** | ***TeamsApiGetTeamTournamentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiGetTeamTournamentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**TournamentsList**](TournamentsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

