# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**CountryBlacklist** | Pointer to **[]string** |  | [optional] 
**CountryWhitelist** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** | ExternalId refers to the ID of the item in a third-party shop | [optional] 
**Games** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Media** | Pointer to [**ItemMedia**](ItemMedia.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**New** | Pointer to **bool** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Sizes** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Unlimited** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewItem

`func NewItem() *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Item) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Item) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Item) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Item) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCountryBlacklist

`func (o *Item) GetCountryBlacklist() []string`

GetCountryBlacklist returns the CountryBlacklist field if non-nil, zero value otherwise.

### GetCountryBlacklistOk

`func (o *Item) GetCountryBlacklistOk() (*[]string, bool)`

GetCountryBlacklistOk returns a tuple with the CountryBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryBlacklist

`func (o *Item) SetCountryBlacklist(v []string)`

SetCountryBlacklist sets CountryBlacklist field to given value.

### HasCountryBlacklist

`func (o *Item) HasCountryBlacklist() bool`

HasCountryBlacklist returns a boolean if a field has been set.

### GetCountryWhitelist

`func (o *Item) GetCountryWhitelist() []string`

GetCountryWhitelist returns the CountryWhitelist field if non-nil, zero value otherwise.

### GetCountryWhitelistOk

`func (o *Item) GetCountryWhitelistOk() (*[]string, bool)`

GetCountryWhitelistOk returns a tuple with the CountryWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryWhitelist

`func (o *Item) SetCountryWhitelist(v []string)`

SetCountryWhitelist sets CountryWhitelist field to given value.

### HasCountryWhitelist

`func (o *Item) HasCountryWhitelist() bool`

HasCountryWhitelist returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Item) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Item) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Item) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Item) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Item) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Item) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Item) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Item) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExternalId

`func (o *Item) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Item) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Item) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Item) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGames

`func (o *Item) GetGames() []string`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *Item) GetGamesOk() (*[]string, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *Item) SetGames(v []string)`

SetGames sets Games field to given value.

### HasGames

`func (o *Item) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Item) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMedia

`func (o *Item) GetMedia() ItemMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *Item) GetMediaOk() (*ItemMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *Item) SetMedia(v ItemMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *Item) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Item) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNew

`func (o *Item) GetNew() bool`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *Item) GetNewOk() (*bool, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *Item) SetNew(v bool)`

SetNew sets New field to given value.

### HasNew

`func (o *Item) HasNew() bool`

HasNew returns a boolean if a field has been set.

### GetOrderType

`func (o *Item) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *Item) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *Item) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *Item) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetPlatforms

`func (o *Item) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *Item) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *Item) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *Item) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetPrice

`func (o *Item) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Item) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Item) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Item) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPublished

`func (o *Item) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *Item) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *Item) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *Item) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetQuantity

`func (o *Item) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Item) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Item) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Item) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSizes

`func (o *Item) GetSizes() []string`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *Item) GetSizesOk() (*[]string, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *Item) SetSizes(v []string)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *Item) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### GetTags

`func (o *Item) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Item) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Item) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Item) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnlimited

`func (o *Item) GetUnlimited() bool`

GetUnlimited returns the Unlimited field if non-nil, zero value otherwise.

### GetUnlimitedOk

`func (o *Item) GetUnlimitedOk() (*bool, bool)`

GetUnlimitedOk returns a tuple with the Unlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimited

`func (o *Item) SetUnlimited(v bool)`

SetUnlimited sets Unlimited field to given value.

### HasUnlimited

`func (o *Item) HasUnlimited() bool`

HasUnlimited returns a boolean if a field has been set.

### GetUrl

`func (o *Item) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Item) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Item) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Item) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


