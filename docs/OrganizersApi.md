# \OrganizersApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizer**](OrganizersApi.md#GetOrganizer) | **Get** /organizers/{organizer_id} | Retrieve organizer details
[**GetOrganizerByName**](OrganizersApi.md#GetOrganizerByName) | **Get** /organizers | Retrieve organizer details from name
[**GetOrganizerChampionships**](OrganizersApi.md#GetOrganizerChampionships) | **Get** /organizers/{organizer_id}/championships | Retrieve all championships of an organizer
[**GetOrganizerGames**](OrganizersApi.md#GetOrganizerGames) | **Get** /organizers/{organizer_id}/games | Retrieve all games an organizer is involved with
[**GetOrganizerHubs**](OrganizersApi.md#GetOrganizerHubs) | **Get** /organizers/{organizer_id}/hubs | Retrieve all hubs of an organizer
[**GetOrganizerTournaments**](OrganizersApi.md#GetOrganizerTournaments) | **Get** /organizers/{organizer_id}/tournaments | Retrieve all tournaments of an organizer


# **GetOrganizer**
> Organizer GetOrganizer(ctx, organizerId)
Retrieve organizer details

Retrieve organizer details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizerId** | **string**| The id of the organizer | 

### Return type

[**Organizer**](Organizer.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizerByName**
> Organizer GetOrganizerByName(ctx, name)
Retrieve organizer details from name

Retrieve organizer details from name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the organizer | 

### Return type

[**Organizer**](Organizer.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizerChampionships**
> ChampionshipsList GetOrganizerChampionships(ctx, organizerId, optional)
Retrieve all championships of an organizer

Retrieve all championships of an organizer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizerId** | **string**| The id of the organizer | 
 **optional** | ***OrganizersApiGetOrganizerChampionshipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizersApiGetOrganizerChampionshipsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**ChampionshipsList**](ChampionshipsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizerGames**
> GamesList GetOrganizerGames(ctx, organizerId)
Retrieve all games an organizer is involved with

Retrieve all games an organizer is involved with

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizerId** | **string**| The id of the organizer | 

### Return type

[**GamesList**](GamesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizerHubs**
> HubsList GetOrganizerHubs(ctx, organizerId, optional)
Retrieve all hubs of an organizer

Retrieve all hubs of an organizer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizerId** | **string**| The id of the organizer | 
 **optional** | ***OrganizersApiGetOrganizerHubsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizersApiGetOrganizerHubsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 50]

### Return type

[**HubsList**](HubsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizerTournaments**
> TournamentsList GetOrganizerTournaments(ctx, organizerId, optional)
Retrieve all tournaments of an organizer

Retrieve all tournaments of an organizer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizerId** | **string**| The id of the organizer | 
 **optional** | ***OrganizersApiGetOrganizerTournamentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizersApiGetOrganizerTournamentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Kind of tournament. Can be upcoming(default) or past | 
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

