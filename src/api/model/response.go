// Package model defines all the model exposed by the application to the rest of the world
package model

import (
	"github.com/gin-gonic/gin"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

// A GenericResponse is our response uniform wrapper for our rest endpoints.
// swagger:response genericResponse
// in: body
type GenericResponse struct {
	// The http response code
	//
	// Required: true
	// Example: 200
	Code int `json:"code"`
	// The http response data in cases where the request was processed successfully (optional)
	//
	// Example: {"id": "uuid", "name": "john doe"}
	Data interface{} `json:"data"`
	// Page is the pagination info
	Page interface{} `json:"page,omitempty"`
	// The success message (optional)
	//
	// Example: User has been created successfully (optional)
	Message *string `json:"message"`
	// The error message (optional)
	//
	// Example: cannot process this request at this time (optional)
	Error *string `json:"error"`
}

// PageInfo struct for api response.
type PageInfo struct {
	Page            int  `json:"page"`
	Size            int  `json:"size"`
	HasNextPage     bool `json:"hasNextPage"`
	HasPreviousPage bool `json:"hasPreviousPage"`
	TotalCount      int  `json:"totalCount"`
}

// Build is a GenericResponse constructor
func Build(code int, data, page interface{}, message, error *string) GenericResponse {
	return GenericResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Page:    page,
		Error:   error,
	}
}

// ErrorResponse template for error responses
func ErrorResponse(c *gin.Context, code int, error string) {
	c.JSON(code, Build(
		code,
		nil,
		nil,
		nil,
		helper.GetStringPointer(error)))
	c.Abort()
}

// OkResponse template for ok and successful responses
func OkResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Build(
		code,
		data,
		nil,
		helper.GetStringPointer(message),
		nil))
	c.Abort()
}

// OkPaginatedResponse template for ok and successful responses
func OkPaginatedResponse(c *gin.Context, code int, message string, data, page interface{}) {
	c.JSON(code, Build(
		code,
		data,
		page,
		helper.GetStringPointer(message),
		nil))
	c.Abort()
}
