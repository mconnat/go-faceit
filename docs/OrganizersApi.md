# \OrganizersAPI

All URIs are relative to *https://open.faceit.com/data/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizer**](OrganizersAPI.md#GetOrganizer) | **Get** /organizers/{organizer_id} | Retrieve organizer details
[**GetOrganizerByName**](OrganizersAPI.md#GetOrganizerByName) | **Get** /organizers | Retrieve organizer details from name
[**GetOrganizerChampionships**](OrganizersAPI.md#GetOrganizerChampionships) | **Get** /organizers/{organizer_id}/championships | Retrieve all championships of an organizer
[**GetOrganizerGames**](OrganizersAPI.md#GetOrganizerGames) | **Get** /organizers/{organizer_id}/games | Retrieve all games an organizer is involved with
[**GetOrganizerHubs**](OrganizersAPI.md#GetOrganizerHubs) | **Get** /organizers/{organizer_id}/hubs | Retrieve all hubs of an organizer
[**GetOrganizerTournaments**](OrganizersAPI.md#GetOrganizerTournaments) | **Get** /organizers/{organizer_id}/tournaments | Retrieve all tournaments of an organizer



## GetOrganizer

> Organizer GetOrganizer(ctx, organizerId).Execute()

Retrieve organizer details



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
	organizerId := "organizerId_example" // string | The id of the organizer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizersAPI.GetOrganizer(context.Background(), organizerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizersAPI.GetOrganizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizer`: Organizer
	fmt.Fprintf(os.Stdout, "Response from `OrganizersAPI.GetOrganizer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizerId** | **string** | The id of the organizer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organizer**](Organizer.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizerByName

> Organizer GetOrganizerByName(ctx).Name(name).Execute()

Retrieve organizer details from name



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
	name := "name_example" // string | The name of the organizer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizersAPI.GetOrganizerByName(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizersAPI.GetOrganizerByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizerByName`: Organizer
	fmt.Fprintf(os.Stdout, "Response from `OrganizersAPI.GetOrganizerByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizerByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the organizer | 

### Return type

[**Organizer**](Organizer.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizerChampionships

> ChampionshipsList GetOrganizerChampionships(ctx, organizerId).Offset(offset).Limit(limit).Execute()

Retrieve all championships of an organizer



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
	organizerId := "organizerId_example" // string | The id of the organizer
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizersAPI.GetOrganizerChampionships(context.Background(), organizerId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizersAPI.GetOrganizerChampionships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizerChampionships`: ChampionshipsList
	fmt.Fprintf(os.Stdout, "Response from `OrganizersAPI.GetOrganizerChampionships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizerId** | **string** | The id of the organizer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizerChampionshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

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


## GetOrganizerGames

> GamesList GetOrganizerGames(ctx, organizerId).Execute()

Retrieve all games an organizer is involved with



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
	organizerId := "organizerId_example" // string | The id of the organizer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizersAPI.GetOrganizerGames(context.Background(), organizerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizersAPI.GetOrganizerGames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizerGames`: GamesList
	fmt.Fprintf(os.Stdout, "Response from `OrganizersAPI.GetOrganizerGames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizerId** | **string** | The id of the organizer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizerGamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetOrganizerHubs

> HubsList GetOrganizerHubs(ctx, organizerId).Offset(offset).Limit(limit).Execute()

Retrieve all hubs of an organizer



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
	organizerId := "organizerId_example" // string | The id of the organizer
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizersAPI.GetOrganizerHubs(context.Background(), organizerId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizersAPI.GetOrganizerHubs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizerHubs`: HubsList
	fmt.Fprintf(os.Stdout, "Response from `OrganizersAPI.GetOrganizerHubs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizerId** | **string** | The id of the organizer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizerHubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 50]

### Return type

[**HubsList**](HubsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizerTournaments

> TournamentsList GetOrganizerTournaments(ctx, organizerId).Type_(type_).Offset(offset).Limit(limit).Execute()

Retrieve all tournaments of an organizer



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
	organizerId := "organizerId_example" // string | The id of the organizer
	type_ := "type__example" // string | Kind of tournament. Can be upcoming(default) or past (optional)
	offset := int32(56) // int32 | The starting item position (optional) (default to 0)
	limit := int32(56) // int32 | The number of items to return (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizersAPI.GetOrganizerTournaments(context.Background(), organizerId).Type_(type_).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizersAPI.GetOrganizerTournaments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizerTournaments`: TournamentsList
	fmt.Fprintf(os.Stdout, "Response from `OrganizersAPI.GetOrganizerTournaments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizerId** | **string** | The id of the organizer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizerTournamentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Kind of tournament. Can be upcoming(default) or past | 
 **offset** | **int32** | The starting item position | [default to 0]
 **limit** | **int32** | The number of items to return | [default to 20]

### Return type

[**TournamentsList**](TournamentsList.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

