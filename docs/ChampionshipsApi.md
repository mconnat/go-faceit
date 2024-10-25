# {{classname}}

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChampionship**](ChampionshipsApi.md#GetChampionship) | **Get** /championships/{championship_id} | Retrieve championship details
[**GetChampionshipMatches**](ChampionshipsApi.md#GetChampionshipMatches) | **Get** /championships/{championship_id}/matches | Retrieve all matches of a championship
[**GetChampionshipResults**](ChampionshipsApi.md#GetChampionshipResults) | **Get** /championships/{championship_id}/results | Retrieve all results of a championship
[**GetChampionshipSubscriptions**](ChampionshipsApi.md#GetChampionshipSubscriptions) | **Get** /championships/{championship_id}/subscriptions | Retrieve all subscriptions of a championship
[**GetChampionships**](ChampionshipsApi.md#GetChampionships) | **Get** /championships | Retrieve all championships of a game

# **GetChampionship**
> Championship GetChampionship(ctx, championshipId, optional)
Retrieve championship details

Retrieve championship details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **championshipId** | **string**| The id of the championship | 
 **optional** | ***ChampionshipsApiGetChampionshipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChampionshipsApiGetChampionshipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expanded** | [**optional.Interface of []string**](string.md)| List of entity names to expand in request | 

### Return type

[**Championship**](Championship.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChampionshipMatches**
> MatchList GetChampionshipMatches(ctx, championshipId, optional)
Retrieve all matches of a championship

Retrieve all matches of a championship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **championshipId** | **string**| The id of the championship | 
 **optional** | ***ChampionshipsApiGetChampionshipMatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChampionshipsApiGetChampionshipMatchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Kind of matches to return. Can be all(default), upcoming, ongoing or past | 
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

# **GetChampionshipResults**
> ChampionshipResultList GetChampionshipResults(ctx, championshipId, optional)
Retrieve all results of a championship

Retrieve all results of a championship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **championshipId** | **string**| The id of the championship | 
 **optional** | ***ChampionshipsApiGetChampionshipResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChampionshipsApiGetChampionshipResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**ChampionshipResultList**](ChampionshipResultList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChampionshipSubscriptions**
> ChampionshipSubscriptionsList GetChampionshipSubscriptions(ctx, championshipId, optional)
Retrieve all subscriptions of a championship

Retrieve all subscriptions of a championship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **championshipId** | **string**| The id of the championship | 
 **optional** | ***ChampionshipsApiGetChampionshipSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChampionshipsApiGetChampionshipSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 10]

### Return type

[**ChampionshipSubscriptionsList**](ChampionshipSubscriptionsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChampionships**
> ChampionshipsList GetChampionships(ctx, game, optional)
Retrieve all championships of a game

Retrieve all championships of a game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **game** | **string**| The id of the game | 
 **optional** | ***ChampionshipsApiGetChampionshipsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChampionshipsApiGetChampionshipsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Kind of matches to return. Can be all(default), upcoming, ongoing or past | 
 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 10]

### Return type

[**ChampionshipsList**](ChampionshipsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

