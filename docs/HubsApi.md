# \HubsApi

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHub**](HubsApi.md#GetHub) | **Get** /hubs/{hub_id} | Retrieve hub details
[**GetHubMatches**](HubsApi.md#GetHubMatches) | **Get** /hubs/{hub_id}/matches | Retrieve all matches of a hub
[**GetHubMembers**](HubsApi.md#GetHubMembers) | **Get** /hubs/{hub_id}/members | Retrieve all members of a hub
[**GetHubRoles**](HubsApi.md#GetHubRoles) | **Get** /hubs/{hub_id}/roles | Retrieve all roles members can have in a hub
[**GetHubRules**](HubsApi.md#GetHubRules) | **Get** /hubs/{hub_id}/rules | Retrieve rules of a hub
[**GetHubStats**](HubsApi.md#GetHubStats) | **Get** /hubs/{hub_id}/stats | Retrieve statistics of a hub


# **GetHub**
> Hub GetHub(ctx, hubId, optional)
Retrieve hub details

Retrieve hub details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
 **optional** | ***HubsApiGetHubOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HubsApiGetHubOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expanded** | [**optional.Interface of []string**](string.md)| List of entity names to expand in request | 

### Return type

[**Hub**](Hub.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubMatches**
> MatchList GetHubMatches(ctx, hubId, optional)
Retrieve all matches of a hub

Retrieve all matches of a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
 **optional** | ***HubsApiGetHubMatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HubsApiGetHubMatchesOpts struct

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

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubMembers**
> HubMembers GetHubMembers(ctx, hubId, optional)
Retrieve all members of a hub

Retrieve all members of a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
 **optional** | ***HubsApiGetHubMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HubsApiGetHubMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 50]

### Return type

[**HubMembers**](HubMembers.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubRoles**
> RolesList GetHubRoles(ctx, hubId, optional)
Retrieve all roles members can have in a hub

Retrieve all roles members can have in a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
 **optional** | ***HubsApiGetHubRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HubsApiGetHubRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 50]

### Return type

[**RolesList**](RolesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubRules**
> Rules GetHubRules(ctx, hubId)
Retrieve rules of a hub

Retrieve rules of a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 

### Return type

[**Rules**](Rules.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHubStats**
> HubStats GetHubStats(ctx, hubId, optional)
Retrieve statistics of a hub

Retrieve statistics of a hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hubId** | **string**| The id of the hub | 
 **optional** | ***HubsApiGetHubStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HubsApiGetHubStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The starting item position | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 20]

### Return type

[**HubStats**](HubStats.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

