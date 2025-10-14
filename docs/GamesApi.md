# \GamesAPI

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllGames**](GamesAPI.md#GetAllGames) | **Get** /games | Retrieve details of all games on FACEIT
[**GetGame**](GamesAPI.md#GetGame) | **Get** /games/{game_id} | Retrieve game details
[**GetGameMatchmakings**](GamesAPI.md#GetGameMatchmakings) | **Get** /games/{gameId}/matchmakings | Retrieve details of all matchmakings of a game on FACEIT
[**GetParentGame**](GamesAPI.md#GetParentGame) | **Get** /games/{game_id}/parent | Retrieve the details of the parent game, if the game is region-specific
[**GetQueueBans**](GamesAPI.md#GetQueueBans) | **Get** /games/{game_id}/queues/{queue_id}/bans | Retrieve queue bans on FACEIT
[**GetQueueById**](GamesAPI.md#GetQueueById) | **Get** /games/{game_id}/queues/{queue_id} | Retrieve details of a queue on FACEIT
[**GetQueuesByEntityFilters**](GamesAPI.md#GetQueuesByEntityFilters) | **Get** /games/{game_id}/queues | Retrieve queues by filters on FACEIT
[**GetQueuesByRegion**](GamesAPI.md#GetQueuesByRegion) | **Get** /games/{game_id}/regions/{region_id}/queues | Retrieve queues by region on FACEIT



## GetAllGames

> GamesList GetAllGames(ctx).Offset(offset).Limit(limit).Execute()

Retrieve details of all games on FACEIT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetAllGames(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetAllGames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllGames`: GamesList
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetAllGames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**GamesList**](GamesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGame

> Game GetGame(ctx, gameId).Execute()

Retrieve game details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gameId := "gameId_example" // string | The id of the game

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetGame(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGame`: Game
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The id of the game | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Game**](Game.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGameMatchmakings

> MatchmakingList GetGameMatchmakings(ctx, gameId).Region(region).Offset(offset).Limit(limit).Execute()

Retrieve details of all matchmakings of a game on FACEIT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gameId := "gameId_example" // string | The id of the game
	region := "region_example" // string | The region of the matchmakings (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetGameMatchmakings(context.Background(), gameId).Region(region).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetGameMatchmakings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGameMatchmakings`: MatchmakingList
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetGameMatchmakings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The id of the game | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGameMatchmakingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | The region of the matchmakings | 
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**MatchmakingList**](MatchmakingList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentGame

> Game GetParentGame(ctx, gameId).Execute()

Retrieve the details of the parent game, if the game is region-specific



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gameId := "gameId_example" // string | The id of the game

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetParentGame(context.Background(), gameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetParentGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParentGame`: Game
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetParentGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The id of the game | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParentGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Game**](Game.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueBans

> QueueBansList GetQueueBans(ctx, gameId, queueId).Offset(offset).Limit(limit).Execute()

Retrieve queue bans on FACEIT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gameId := "gameId_example" // string | The id of the game
	queueId := "queueId_example" // string | The id of the queue
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetQueueBans(context.Background(), gameId, queueId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetQueueBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueueBans`: QueueBansList
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetQueueBans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The id of the game | 
**queueId** | **string** | The id of the queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**QueueBansList**](QueueBansList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueById

> Queue GetQueueById(ctx, gameId, queueId).Execute()

Retrieve details of a queue on FACEIT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gameId := "gameId_example" // string | The id of the game
	queueId := "queueId_example" // string | The id of the queue

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetQueueById(context.Background(), gameId, queueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetQueueById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueueById`: Queue
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetQueueById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The id of the game | 
**queueId** | **string** | The id of the queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Queue**](Queue.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueuesByEntityFilters

> QueuesList GetQueuesByEntityFilters(ctx, gameId).EntityType(entityType).EntityId(entityId).Offset(offset).Limit(limit).Execute()

Retrieve queues by filters on FACEIT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gameId := "gameId_example" // string | The id of the game
	entityType := "entityType_example" // string | The type of the entity
	entityId := "entityId_example" // string | The id of the entity
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetQueuesByEntityFilters(context.Background(), gameId).EntityType(entityType).EntityId(entityId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetQueuesByEntityFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueuesByEntityFilters`: QueuesList
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetQueuesByEntityFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The id of the game | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueuesByEntityFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityType** | **string** | The type of the entity | 
 **entityId** | **string** | The id of the entity | 
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**QueuesList**](QueuesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueuesByRegion

> QueuesList GetQueuesByRegion(ctx, gameId, regionId).Offset(offset).Limit(limit).Execute()

Retrieve queues by region on FACEIT



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	gameId := "gameId_example" // string | The id of the game
	regionId := "regionId_example" // string | The id of the region
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GamesAPI.GetQueuesByRegion(context.Background(), gameId, regionId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GamesAPI.GetQueuesByRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueuesByRegion`: QueuesList
	fmt.Fprintf(os.Stdout, "Response from `GamesAPI.GetQueuesByRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gameId** | **string** | The id of the game | 
**regionId** | **string** | The id of the region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueuesByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**QueuesList**](QueuesList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

