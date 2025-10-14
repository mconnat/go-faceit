# FilesystemChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **int32** | Can be one of:  &#x60;0&#x60;: Modified (\&quot;C\&quot;) &#x60;1&#x60;: Added (\&quot;A\&quot;) &#x60;2&#x60;: Deleted (\&quot;D\&quot;) | 
**Path** | **string** | Path to file or directory that has changed. | 

## Methods

### NewFilesystemChange

`func NewFilesystemChange(kind int32, path string, ) *FilesystemChange`

NewFilesystemChange instantiates a new FilesystemChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemChangeWithDefaults

`func NewFilesystemChangeWithDefaults() *FilesystemChange`

NewFilesystemChangeWithDefaults instantiates a new FilesystemChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *FilesystemChange) GetKind() int32`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FilesystemChange) GetKindOk() (*int32, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FilesystemChange) SetKind(v int32)`

SetKind sets Kind field to given value.


### GetPath

`func (o *FilesystemChange) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesystemChange) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesystemChange) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


