# {{classname}}

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTournament**](TournamentsApi.md#GetTournament) | **Get** /tournaments/{tournament_id} | Retrieve tournament details
[**GetTournamentBrackets**](TournamentsApi.md#GetTournamentBrackets) | **Get** /tournaments/{tournament_id}/brackets | Retrieve brackets of a tournament
[**GetTournamentMatches**](TournamentsApi.md#GetTournamentMatches) | **Get** /tournaments/{tournament_id}/matches | Retrieve all matches of a tournament
[**GetTournamentTeams**](TournamentsApi.md#GetTournamentTeams) | **Get** /tournaments/{tournament_id}/teams | Retrieve all teams of a tournament
[**GetTournamentsList**](TournamentsApi.md#GetTournamentsList) | **Get** /tournaments | Retrieve tournaments v1 (no longer used)

# **GetTournament**
> Tournament GetTournament(ctx, tournamentId, optional)
Retrieve tournament details

Retrieve tournament details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tournamentId** | **string**| The id of the tournament | 
 **optional** | ***TournamentsApiGetTournamentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TournamentsApiGetTournamentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expanded** | [**optional.Interface of []string**](string.md)| List of entity names to expand in request | 

### Return type

[**Tournament**](Tournament.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTournamentBrackets**
> Brackets GetTournamentBrackets(ctx, tournamentId)
Retrieve brackets of a tournament

Retrieve brackets of a tournament

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tournamentId** | **string**| The id of the tournament | 

### Return type

[**Brackets**](Brackets.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTournamentMatches**
> MatchList GetTournamentMatches(ctx, tournamentId, optional)
Retrieve all matches of a tournament

Retrieve all matches of a tournament

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tournamentId** | **string**| The id of the tournament | 
 **optional** | ***TournamentsApiGetTournamentMatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TournamentsApiGetTournamentMatchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**MatchList**](MatchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTournamentTeams**
> TournamentTeams GetTournamentTeams(ctx, tournamentId, optional)
Retrieve all teams of a tournament

Retrieve all teams of a tournament

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tournamentId** | **string**| The id of the tournament | 
 **optional** | ***TournamentsApiGetTournamentTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TournamentsApiGetTournamentTeamsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**TournamentTeams**](TournamentTeams.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTournamentsList**
> TournamentsList GetTournamentsList(ctx, optional)
Retrieve tournaments v1 (no longer used)

Retrieve tournaments v1 (no longer used). Please refer to the Championships controller to retrieve tournaments v2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TournamentsApiGetTournamentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TournamentsApiGetTournamentsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **game** | **optional.String**| A game on FACEIT | 
 **region** | **optional.String**| A region of the game | 
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

