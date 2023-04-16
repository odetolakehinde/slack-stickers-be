package store

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/environment"
	"github.com/odetolakehinde/slack-stickers-be/src/store/mock"
)

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

type Suite struct {
	suite.Suite
	redisStore Store
}

func (s *Suite) SetupSuite() {
	env := &environment.Env{}
	env.HelperForMocking(map[string]string{
		"REDIS_SERVER_ADDRESS": "dummy-redis-server-address",
	})
	dummyLog := zerolog.Nop()

	s.redisStore = NewRedis(env, dummyLog, ConnectionInfo{})
}

func (s *Suite) AfterTest(_, _ string) {
}

func (s *Suite) Test_GetValue() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	foundKey := "FOUND_KEY"
	errorKey := "ERROR_KEY"
	m := mock.NewMockStore(ctrl)
	m.EXPECT().GetValue(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, key string, result *string) error {
		if strings.EqualFold(key, foundKey) {
			*result = foundKey
			return nil
		} else if strings.EqualFold(key, errorKey) {
			result = nil
			return ErrFailedToRetrieveValue
		}

		return nil // dont error but return nil for error and value retrieved
	}).AnyTimes()

	kv := Store(m)
	var fVal string
	fErr := kv.GetValue(context.Background(), foundKey, &fVal)
	require.NoError(s.T(), fErr)
	require.Equal(s.T(), fVal, foundKey)
	var eVal *string
	eErr := kv.GetValue(context.Background(), errorKey, eVal)
	require.Error(s.T(), eErr)
	require.Nil(s.T(), eVal)
	require.EqualError(s.T(), eErr, ErrFailedToRetrieveValue.Error())
}
