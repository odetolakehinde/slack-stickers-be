// Package store houses all the connections related to redis
package store

import (
	"context"
	"errors"
	"time"
)

// Store interface
//go:generate mockgen -source store.go -destination ./mock/mock_store.go -package mock Store
type Store interface {
	GetValue(ctx context.Context, key string, result interface{}) error
	GetStringValue(ctx context.Context, key string) (string, error)
	SetValue(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	DeleteValue(ctx context.Context, key string) error
	Connect() error
}

var (
	// ErrConnectionToSourceFailed if the connection to the data source cannot be established
	ErrConnectionToSourceFailed = errors.New("connection to date source cannot be established")
	// ErrFailedToRetrieveValue if there is issue retrieving the value from source
	ErrFailedToRetrieveValue = errors.New("failed to retrieve the value from source")
)
