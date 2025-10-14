# \ChampionshipsAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                                               | HTTP request                                           | Description                                  |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------ | -------------------------------------------- |
| [**GetChampionship**](ChampionshipsAPI.md#GetChampionship)                           | **Get** /championships/{championship_id}               | Retrieve championship details                |
| [**GetChampionshipMatches**](ChampionshipsAPI.md#GetChampionshipMatches)             | **Get** /championships/{championship_id}/matches       | Retrieve all matches of a championship       |
| [**GetChampionshipResults**](ChampionshipsAPI.md#GetChampionshipResults)             | **Get** /championships/{championship_id}/results       | Retrieve all results of a championship       |
| [**GetChampionshipSubscriptions**](ChampionshipsAPI.md#GetChampionshipSubscriptions) | **Get** /championships/{championship_id}/subscriptions | Retrieve all subscriptions of a championship |
| [**GetChampionships**](ChampionshipsAPI.md#GetChampionships)                         | **Get** /championships                                 | Retrieve all championships of a game         |



## GetChampionship

> Championship GetChampionship(ctx, championshipId).Expanded(expanded).Execute()

Retrieve championship details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mconnat/go-faceit"
)

func main() {
	championshipId := "championshipId_example" // string | The id of the championship
	expanded := []string{"Expanded_example"} // []string | List of entity names to expand in request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChampionshipsAPI.GetChampionship(context.Background(), championshipId).Expanded(expanded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChampionshipsAPI.GetChampionship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChampionship`: Championship
	fmt.Fprintf(os.Stdout, "Response from `ChampionshipsAPI.GetChampionship`: %v\n", resp)
}
```

### Path Parameters


| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **championshipId** | **string**          | The id of the championship                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetChampionshipRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **expanded** | **[]string** | List of entity names to expand in request | 

### Return type

[**Championship**](Championship.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChampionshipMatches

> MatchList GetChampionshipMatches(ctx, championshipId).Type_(type_).Offset(offset).Limit(limit).Execute()

Retrieve all matches of a championship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mconnat/go-faceit"
)

func main() {
	championshipId := "championshipId_example" // string | The id of the championship
	type_ := "type__example" // string | Kind of matches to return. Can be all(default), upcoming, ongoing or past (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChampionshipsAPI.GetChampionshipMatches(context.Background(), championshipId).Type_(type_).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChampionshipsAPI.GetChampionshipMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChampionshipMatches`: MatchList
	fmt.Fprintf(os.Stdout, "Response from `ChampionshipsAPI.GetChampionshipMatches`: %v\n", resp)
}
```

### Path Parameters


| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **championshipId** | **string**          | The id of the championship                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetChampionshipMatchesRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **type_** | **string** | Kind of matches to return. Can be all(default), upcoming, ongoing or past | 
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**MatchList**](MatchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChampionshipResults

> ChampionshipResultList GetChampionshipResults(ctx, championshipId).Offset(offset).Limit(limit).Execute()

Retrieve all results of a championship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mconnat/go-faceit"
)

func main() {
	championshipId := "championshipId_example" // string | The id of the championship
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChampionshipsAPI.GetChampionshipResults(context.Background(), championshipId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChampionshipsAPI.GetChampionshipResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChampionshipResults`: ChampionshipResultList
	fmt.Fprintf(os.Stdout, "Response from `ChampionshipsAPI.GetChampionshipResults`: %v\n", resp)
}
```

### Path Parameters


| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **championshipId** | **string**          | The id of the championship                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetChampionshipResultsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**ChampionshipResultList**](ChampionshipResultList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChampionshipSubscriptions

> ChampionshipSubscriptionsList GetChampionshipSubscriptions(ctx, championshipId).Offset(offset).Limit(limit).Execute()

Retrieve all subscriptions of a championship



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mconnat/go-faceit"
)

func main() {
	championshipId := "championshipId_example" // string | The id of the championship
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChampionshipsAPI.GetChampionshipSubscriptions(context.Background(), championshipId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChampionshipsAPI.GetChampionshipSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChampionshipSubscriptions`: ChampionshipSubscriptionsList
	fmt.Fprintf(os.Stdout, "Response from `ChampionshipsAPI.GetChampionshipSubscriptions`: %v\n", resp)
}
```

### Path Parameters


| Name               | Type                | Description                                                                 | Notes |
| ------------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**            | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **championshipId** | **string**          | The id of the championship                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetChampionshipSubscriptionsRequest struct via the builder pattern


| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 10]

### Return type

[**ChampionshipSubscriptionsList**](ChampionshipSubscriptionsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChampionships

> ChampionshipsList GetChampionships(ctx).Game(game).Type_(type_).Offset(offset).Limit(limit).Execute()

Retrieve all championships of a game



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mconnat/go-faceit"
)

func main() {
	game := "game_example" // string | The id of the game
	type_ := "type__example" // string | Kind of matches to return. Can be all(default), upcoming, ongoing or past (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChampionshipsAPI.GetChampionships(context.Background()).Game(game).Type_(type_).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChampionshipsAPI.GetChampionships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChampionships`: ChampionshipsList
	fmt.Fprintf(os.Stdout, "Response from `ChampionshipsAPI.GetChampionships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChampionshipsRequest struct via the builder pattern


| Name       | Type       | Description                                                               | Notes           |
| ---------- | ---------- | ------------------------------------------------------------------------- | --------------- |
| **game**   | **string** | The id of the game                                                        |
| **type_**  | **string** | Kind of matches to return. Can be all(default), upcoming, ongoing or past |
| **offset** | **int32**  | The starting item position                                                | [default to 0]  |
| **limit**  | **int32**  | The number of items to return                                             | [default to 10] |

### Return type

[**ChampionshipsList**](ChampionshipsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

