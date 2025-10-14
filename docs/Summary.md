# Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | **int64** | Number of containers using this image. Includes both stopped and running containers.  This size is not calculated by default, and depends on which API endpoint is used. &#x60;-1&#x60; indicates that the value has not been set / calculated. | 
**Created** | **int64** | Date and time at which the image was created as a Unix timestamp (number of seconds sinds EPOCH). | 
**Id** | **string** | ID is the content-addressable ID of an image.  This identifier is a content-addressable digest calculated from the image&#39;s configuration (which includes the digests of layers used by the image).  Note that this digest differs from the &#x60;RepoDigests&#x60; below, which holds digests of image manifests that reference the image. | 
**Labels** | **map[string]string** | User-defined key/value metadata. | 
**ParentId** | **string** | ID of the parent image.  Depending on how the image was created, this field may be empty and is only set for images that were built/created locally. This field is empty if the image was pulled from an image registry. | 
**RepoDigests** | **[]string** | List of content-addressable digests of locally available image manifests that the image is referenced from. Multiple manifests can refer to the same image.  These digests are usually only available if the image was either pulled from a registry, or if the image was pushed to a registry, which is when the manifest is generated and its digest calculated. | 
**RepoTags** | **[]string** | List of image names/tags in the local image cache that reference this image.  Multiple image tags can refer to the same image, and this list may be empty if no tags reference the image, in which case the image is \&quot;untagged\&quot;, in which case it can still be referenced by its ID. | 
**SharedSize** | **int64** | Total size of image layers that are shared between this image and other images.  This size is not calculated by default. &#x60;-1&#x60; indicates that the value has not been set / calculated. | 
**Size** | **int64** | Total size of the image including all layers it is composed of. | 
**VirtualSize** | Pointer to **int64** | Total size of the image including all layers it is composed of.  Deprecated: this field is omitted in API v1.44, but kept for backward compatibility. Use Size instead. | [optional] 

## Methods

### NewSummary

`func NewSummary(containers int64, created int64, id string, labels map[string]string, parentId string, repoDigests []string, repoTags []string, sharedSize int64, size int64, ) *Summary`

NewSummary instantiates a new Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryWithDefaults

`func NewSummaryWithDefaults() *Summary`

NewSummaryWithDefaults instantiates a new Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *Summary) GetContainers() int64`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Summary) GetContainersOk() (*int64, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Summary) SetContainers(v int64)`

SetContainers sets Containers field to given value.


### GetCreated

`func (o *Summary) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Summary) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Summary) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetId

`func (o *Summary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Summary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Summary) SetId(v string)`

SetId sets Id field to given value.


### GetLabels

`func (o *Summary) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Summary) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Summary) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetParentId

`func (o *Summary) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Summary) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Summary) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetRepoDigests

`func (o *Summary) GetRepoDigests() []string`

GetRepoDigests returns the RepoDigests field if non-nil, zero value otherwise.

### GetRepoDigestsOk

`func (o *Summary) GetRepoDigestsOk() (*[]string, bool)`

GetRepoDigestsOk returns a tuple with the RepoDigests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoDigests

`func (o *Summary) SetRepoDigests(v []string)`

SetRepoDigests sets RepoDigests field to given value.


### GetRepoTags

`func (o *Summary) GetRepoTags() []string`

GetRepoTags returns the RepoTags field if non-nil, zero value otherwise.

### GetRepoTagsOk

`func (o *Summary) GetRepoTagsOk() (*[]string, bool)`

GetRepoTagsOk returns a tuple with the RepoTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoTags

`func (o *Summary) SetRepoTags(v []string)`

SetRepoTags sets RepoTags field to given value.


### GetSharedSize

`func (o *Summary) GetSharedSize() int64`

GetSharedSize returns the SharedSize field if non-nil, zero value otherwise.

### GetSharedSizeOk

`func (o *Summary) GetSharedSizeOk() (*int64, bool)`

GetSharedSizeOk returns a tuple with the SharedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSize

`func (o *Summary) SetSharedSize(v int64)`

SetSharedSize sets SharedSize field to given value.


### GetSize

`func (o *Summary) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Summary) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Summary) SetSize(v int64)`

SetSize sets Size field to given value.


### GetVirtualSize

`func (o *Summary) GetVirtualSize() int64`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *Summary) GetVirtualSizeOk() (*int64, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *Summary) SetVirtualSize(v int64)`

SetVirtualSize sets VirtualSize field to given value.

### HasVirtualSize

`func (o *Summary) HasVirtualSize() bool`

HasVirtualSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


