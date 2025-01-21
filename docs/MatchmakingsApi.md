# \MatchmakingsApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMatchmaking**](MatchmakingsApi.md#GetMatchmaking) | **Get** /matchmakings/{matchmaking_id} | Retrieve details of a matchmaking of a game on FACEIT


# **GetMatchmaking**
> Matchmaking GetMatchmaking(ctx, matchmakingId)
Retrieve details of a matchmaking of a game on FACEIT

Retrieve details of a matchmaking of a game on FACEIT

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchmakingId** | **string**| The id of the matchmaking | 

### Return type

[**Matchmaking**](Matchmaking.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

