package store

import (
	"context"
	"crypto/tls"
	"os"
	"strings"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/model/env"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

const packageNameRedis = "store.redis"

type (
	// Redis store object
	Redis struct {
		env             *environment.Env
		logger          zerolog.Logger
		connectionError error
		client          *redis.Client
		ConnectionInfo  ConnectionInfo
	}
	// ConnectionInfo connection info
	ConnectionInfo struct {
		Address  string
		Password string
		Username string
	}
)

// NewRedis creates a new Redis object as a KeyValue instance
func NewRedis(e *environment.Env, z zerolog.Logger, c ConnectionInfo) Store {
	log := z.With().Str(helper.LogStrPackageLevel, packageNameRedis).Logger()
	r := &Redis{
		env:            e,
		logger:         log,
		ConnectionInfo: c,
	}
	// connect to the storage
	r.connectionError = r.Connect()
	return Store(r)
}

// Connect handles connection to data source server implementation
func (r *Redis) Connect() error {
	ctx := context.Background()

	var tlsConfig *tls.Config
	if !strings.EqualFold(os.Getenv(env.RedisTLSEnabled), "false") {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:      r.ConnectionInfo.Address,
		Password:  r.ConnectionInfo.Password,
		Username:  r.ConnectionInfo.Username,
		DB:        0, // use default DB,
		TLSConfig: tlsConfig,
	})

	st := rdb.Ping(ctx)
	if err := st.Err(); err != nil {
		r.connectionError = err
		r.logger.Fatal().Err(err).Msg("connection to redis server failed")
		return err
	}
	r.logger.Info().Msg("[success] connected to redis server")
	r.connectionError = nil // connection did NOT fail.
	r.client = rdb          // set the client
	return nil
}

// GetValue retrieves the value of a key from inside redis
func (r *Redis) GetValue(ctx context.Context, key string, result interface{}) error {
	if r.connectionError != nil {
		// attempt to reconnect
		err := r.Connect()
		if err != nil {
			return ErrConnectionToSourceFailed
		}
	}
	// hopefully the connection to the store is okay
	err := r.client.Get(ctx, key).Scan(result)
	if err != nil {
		// connecting issue or not able to retrieve from server
		r.logger.Err(err).Str("key", key).Msg(ErrFailedToRetrieveValue.Error())
		return ErrFailedToRetrieveValue
	}
	return nil
}

// GetStringValue retrieves the value of a key from inside redis as string
func (r *Redis) GetStringValue(ctx context.Context, key string) (string, error) {
	if r.connectionError != nil {
		// attempt to reconnect
		err := r.Connect()
		if err != nil {
			return "", ErrConnectionToSourceFailed
		}
	}

	// hopefully the connection to the store is okay
	return r.client.Get(ctx, key).String(), nil
}

// SetValue [Not Implemented] sets and writes value into redis
func (r *Redis) SetValue(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	res := r.client.Set(ctx, key, value, ttl)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

// DeleteValue will delete redis key
func (r *Redis) DeleteValue(ctx context.Context, key string) error {
	res := r.client.Del(ctx, key)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
