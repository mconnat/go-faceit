# {{classname}}

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMatch**](MatchesApi.md#GetMatch) | **Get** /matches/{match_id} | Retrieve match details
[**GetMatchStats**](MatchesApi.md#GetMatchStats) | **Get** /matches/{match_id}/stats | Retrieve statistics of a match

# **GetMatch**
> Match GetMatch(ctx, matchId)
Retrieve match details

Retrieve match details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchId** | **string**| The id of the match | 

### Return type

[**Match**](Match.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchStats**
> MatchStats GetMatchStats(ctx, matchId)
Retrieve statistics of a match

Retrieve statistics of a match

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **matchId** | **string**| The id of the match | 

### Return type

[**MatchStats**](MatchStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

