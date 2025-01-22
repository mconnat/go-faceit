/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

// OK response to ContainerWait operation
type WaitResponse struct {
	Error_ *WaitExitError `json:"Error,omitempty"`
	// Exit code of the container
	StatusCode int64 `json:"StatusCode"`
}
