# \SearchAPI

All URIs are relative to *https://open.faceit.com/data/v4*

| Method                                                      | HTTP request                  | Description              |
| ----------------------------------------------------------- | ----------------------------- | ------------------------ |
| [**SearchChampionships**](SearchAPI.md#SearchChampionships) | **Get** /search/championships | Search for championships |
| [**SearchClans**](SearchAPI.md#SearchClans)                 | **Get** /search/clans         | Search for clans         |
| [**SearchHubs**](SearchAPI.md#SearchHubs)                   | **Get** /search/hubs          | Search for hubs          |
| [**SearchOrganizers**](SearchAPI.md#SearchOrganizers)       | **Get** /search/organizers    | Search for organizers    |
| [**SearchPlayers**](SearchAPI.md#SearchPlayers)             | **Get** /search/players       | Search for players       |
| [**SearchTeams**](SearchAPI.md#SearchTeams)                 | **Get** /search/teams         | Search for teams         |
| [**SearchTournaments**](SearchAPI.md#SearchTournaments)     | **Get** /search/tournaments   | Search for tournaments   |



## SearchChampionships

> CompetitionsSearchList SearchChampionships(ctx).Name(name).Game(game).Region(region).Type_(type_).Offset(offset).Limit(limit).Execute()

Search for championships



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
	name := "name_example" // string | The name of a championship on FACEIT
	game := "game_example" // string | A game on FACEIT (optional)
	region := "region_example" // string | A region of the game (optional)
	type_ := "type__example" // string | Kind of competitions to return (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchChampionships(context.Background()).Name(name).Game(game).Region(region).Type_(type_).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchChampionships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChampionships`: CompetitionsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchChampionships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchChampionshipsRequest struct via the builder pattern


| Name       | Type       | Description                          | Notes           |
| ---------- | ---------- | ------------------------------------ | --------------- |
| **name**   | **string** | The name of a championship on FACEIT |
| **game**   | **string** | A game on FACEIT                     |
| **region** | **string** | A region of the game                 |
| **type_**  | **string** | Kind of competitions to return       |
| **offset** | **int32**  | The starting item position           | [default to 0]  |
| **limit**  | **int32**  | The number of items to return        | [default to 20] |

### Return type

[**CompetitionsSearchList**](CompetitionsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchClans

> ClansSearchList SearchClans(ctx).Name(name).Game(game).Region(region).Offset(offset).Limit(limit).Execute()

Search for clans



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
	name := "name_example" // string | The name of a clan on FACEIT
	game := "game_example" // string | A game on FACEIT (optional)
	region := "region_example" // string | A region of the game (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchClans(context.Background()).Name(name).Game(game).Region(region).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchClans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchClans`: ClansSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchClans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchClansRequest struct via the builder pattern


| Name       | Type       | Description                   | Notes           |
| ---------- | ---------- | ----------------------------- | --------------- |
| **name**   | **string** | The name of a clan on FACEIT  |
| **game**   | **string** | A game on FACEIT              |
| **region** | **string** | A region of the game          |
| **offset** | **int32**  | The starting item position    | [default to 0]  |
| **limit**  | **int32**  | The number of items to return | [default to 20] |

### Return type

[**ClansSearchList**](ClansSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchHubs

> CompetitionsSearchList SearchHubs(ctx).Name(name).Game(game).Region(region).Offset(offset).Limit(limit).Execute()

Search for hubs



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
	name := "name_example" // string | The name of a hub on FACEIT
	game := "game_example" // string | A game on FACEIT (optional)
	region := "region_example" // string | A region of the game (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchHubs(context.Background()).Name(name).Game(game).Region(region).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchHubs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchHubs`: CompetitionsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchHubs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchHubsRequest struct via the builder pattern


| Name       | Type       | Description                   | Notes           |
| ---------- | ---------- | ----------------------------- | --------------- |
| **name**   | **string** | The name of a hub on FACEIT   |
| **game**   | **string** | A game on FACEIT              |
| **region** | **string** | A region of the game          |
| **offset** | **int32**  | The starting item position    | [default to 0]  |
| **limit**  | **int32**  | The number of items to return | [default to 20] |

### Return type

[**CompetitionsSearchList**](CompetitionsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrganizers

> OrganizersSearchList SearchOrganizers(ctx).Name(name).Offset(offset).Limit(limit).Execute()

Search for organizers



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
	name := "name_example" // string | The name of a organizer on FACEIT
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchOrganizers(context.Background()).Name(name).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchOrganizers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrganizers`: OrganizersSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchOrganizers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrganizersRequest struct via the builder pattern


| Name       | Type       | Description                       | Notes           |
| ---------- | ---------- | --------------------------------- | --------------- |
| **name**   | **string** | The name of a organizer on FACEIT |
| **offset** | **int32**  | The starting item position        | [default to 0]  |
| **limit**  | **int32**  | The number of items to return     | [default to 20] |

### Return type

[**OrganizersSearchList**](OrganizersSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPlayers

> UsersSearchList SearchPlayers(ctx).Nickname(nickname).Game(game).Country(country).Offset(offset).Limit(limit).Execute()

Search for players



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
	nickname := "nickname_example" // string | The nickname of a player on FACEIT
	game := "game_example" // string | A game on FACEIT (optional)
	country := "country_example" // string | A country code (ISO 3166-1) (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchPlayers(context.Background()).Nickname(nickname).Game(game).Country(country).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPlayers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPlayers`: UsersSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPlayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPlayersRequest struct via the builder pattern


| Name         | Type       | Description                        | Notes           |
| ------------ | ---------- | ---------------------------------- | --------------- |
| **nickname** | **string** | The nickname of a player on FACEIT |
| **game**     | **string** | A game on FACEIT                   |
| **country**  | **string** | A country code (ISO 3166-1)        |
| **offset**   | **int32**  | The starting item position         | [default to 0]  |
| **limit**    | **int32**  | The number of items to return      | [default to 20] |

### Return type

[**UsersSearchList**](UsersSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTeams

> TeamsSearchList SearchTeams(ctx).Nickname(nickname).Game(game).Offset(offset).Limit(limit).Execute()

Search for teams



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
	nickname := "nickname_example" // string | The nickname of a team on FACEIT
	game := "game_example" // string | A game on FACEIT (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchTeams(context.Background()).Nickname(nickname).Game(game).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTeams`: TeamsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTeamsRequest struct via the builder pattern


| Name         | Type       | Description                      | Notes           |
| ------------ | ---------- | -------------------------------- | --------------- |
| **nickname** | **string** | The nickname of a team on FACEIT |
| **game**     | **string** | A game on FACEIT                 |
| **offset**   | **int32**  | The starting item position       | [default to 0]  |
| **limit**    | **int32**  | The number of items to return    | [default to 20] |

### Return type

[**TeamsSearchList**](TeamsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTournaments

> CompetitionsSearchList SearchTournaments(ctx).Name(name).Game(game).Region(region).Type_(type_).Offset(offset).Limit(limit).Execute()

Search for tournaments



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
	name := "name_example" // string | The name of a tournament on FACEIT
	game := "game_example" // string | A game on FACEIT (optional)
	region := "region_example" // string | A region of the game (optional)
	type_ := "type__example" // string | Kind of competitions to return (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchTournaments(context.Background()).Name(name).Game(game).Region(region).Type_(type_).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTournaments`: CompetitionsSearchList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchTournaments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTournamentsRequest struct via the builder pattern


| Name       | Type       | Description                        | Notes           |
| ---------- | ---------- | ---------------------------------- | --------------- |
| **name**   | **string** | The name of a tournament on FACEIT |
| **game**   | **string** | A game on FACEIT                   |
| **region** | **string** | A region of the game               |
| **type_**  | **string** | Kind of competitions to return     |
| **offset** | **int32**  | The starting item position         | [default to 0]  |
| **limit**  | **int32**  | The number of items to return      | [default to 20] |

### Return type

[**CompetitionsSearchList**](CompetitionsSearchList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

