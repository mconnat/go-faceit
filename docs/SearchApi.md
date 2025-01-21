# \SearchApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchChampionships**](SearchApi.md#SearchChampionships) | **Get** /search/championships | Search for championships
[**SearchClans**](SearchApi.md#SearchClans) | **Get** /search/clans | Search for clans
[**SearchHubs**](SearchApi.md#SearchHubs) | **Get** /search/hubs | Search for hubs
[**SearchOrganizers**](SearchApi.md#SearchOrganizers) | **Get** /search/organizers | Search for organizers
[**SearchPlayers**](SearchApi.md#SearchPlayers) | **Get** /search/players | Search for players
[**SearchTeams**](SearchApi.md#SearchTeams) | **Get** /search/teams | Search for teams
[**SearchTournaments**](SearchApi.md#SearchTournaments) | **Get** /search/tournaments | Search for tournaments


# **SearchChampionships**
> CompetitionsSearchList SearchChampionships(ctx, name, optional)
Search for championships

Search for championships

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of a championship on FACEIT | 
 **optional** | ***SearchApiSearchChampionshipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchChampionshipsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **game** | **optional.String**| A game on FACEIT | 
 **region** | **optional.String**| A region of the game | 
 **type_** | **optional.String**| Kind of competitions to return | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**CompetitionsSearchList**](CompetitionsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchClans**
> ClansSearchList SearchClans(ctx, name, optional)
Search for clans

Search for clans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of a clan on FACEIT | 
 **optional** | ***SearchApiSearchClansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchClansOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **game** | **optional.String**| A game on FACEIT | 
 **region** | **optional.String**| A region of the game | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**ClansSearchList**](ClansSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchHubs**
> CompetitionsSearchList SearchHubs(ctx, name, optional)
Search for hubs

Search for hubs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of a hub on FACEIT | 
 **optional** | ***SearchApiSearchHubsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchHubsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **game** | **optional.String**| A game on FACEIT | 
 **region** | **optional.String**| A region of the game | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**CompetitionsSearchList**](CompetitionsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrganizers**
> OrganizersSearchList SearchOrganizers(ctx, name, optional)
Search for organizers

Search for organizers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of a organizer on FACEIT | 
 **optional** | ***SearchApiSearchOrganizersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchOrganizersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**OrganizersSearchList**](OrganizersSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchPlayers**
> UsersSearchList SearchPlayers(ctx, nickname, optional)
Search for players

Search for players

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nickname** | **string**| The nickname of a player on FACEIT | 
 **optional** | ***SearchApiSearchPlayersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchPlayersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **game** | **optional.String**| A game on FACEIT | 
 **country** | **optional.String**| A country code (ISO 3166-1) | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**UsersSearchList**](UsersSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTeams**
> TeamsSearchList SearchTeams(ctx, nickname, optional)
Search for teams

Search for teams

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nickname** | **string**| The nickname of a team on FACEIT | 
 **optional** | ***SearchApiSearchTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchTeamsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **game** | **optional.String**| A game on FACEIT | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**TeamsSearchList**](TeamsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTournaments**
> CompetitionsSearchList SearchTournaments(ctx, name, optional)
Search for tournaments

Search for tournaments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of a tournament on FACEIT | 
 **optional** | ***SearchApiSearchTournamentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchTournamentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **game** | **optional.String**| A game on FACEIT | 
 **region** | **optional.String**| A region of the game | 
 **type_** | **optional.String**| Kind of competitions to return | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**CompetitionsSearchList**](CompetitionsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

