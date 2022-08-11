// Package model defines all the model exposed by the application to the rest of the world
package model

import "github.com/google/uuid"

type (
	// IDArrays contains an array of uuids as string
	IDArrays []string

	// Page is a struct for pagination purpose.
	Page struct {
		Number            *int    `json:"number"`
		Size              *int    `json:"size"`
		SortBy            *string `json:"sort_by"`
		SortDirectionDesc *bool   `json:"sort_direction_desc"`
	}
)

// Parse validates IDArrays to ensure that all the strings are valid uuids.
func (i IDArrays) Parse() []uuid.UUID {
	var uuidArray []uuid.UUID
	for _, id := range i {
		uid, err := uuid.Parse(id)
		if err == nil {
			uuidArray = append(uuidArray, uid)
		}
	}
	return uuidArray
}
