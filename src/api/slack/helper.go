package slack

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/odetolakehinde/slack-stickers-be/src/pkg/helper"
)

// validateSlackSignature validate slack signature
// https://api.slack.com/authentication/verifying-requests-from-slack
func validateSlackSignature(c *gin.Context, log zerolog.Logger, signingSecret string) ([]byte, int, error) {
	log = log.With().Str(helper.LogStrKeyMethod, "validateSlackSignature").Logger()

	const (
		// versionPrefix defines the version identifier used in Slack's signature verification scheme.
		versionPrefix = "v0"

		// maxAllowedTimeDrift specifies the maximum duration of time drift allowed between
		// an incoming request's timestamp and the server's current time. Requests with
		// timestamps outside this window are considered invalid to protect against replay attacks.
		maxAllowedTimeDrift = 2 * time.Minute
	)

	slackSignature := c.GetHeader("X-Slack-Signature")
	slackRequesTimestamp := c.GetHeader("X-Slack-Request-Timestamp")

	if slackRequesTimestamp == "" || slackSignature == "" {
		err := errors.New("missing Slack signature or timestamp")
		log.Error().Err(err).Msg("missing headers")
		return nil, http.StatusBadRequest, err
	}

	ts, err := strconv.ParseInt(slackRequesTimestamp, 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("invalid timestamp format")
		return nil, http.StatusBadRequest, err
	}

	if abs(time.Now().Unix()-ts) > int64(maxAllowedTimeDrift.Seconds()) {
		err := errors.New("request timestamp too old")
		log.Error().Err(err).Msg("old timestamp")
		return nil, http.StatusUnauthorized, err
	}

	// Read and keep the body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Err(err).Msg("failed to read body")
		return nil, http.StatusInternalServerError, err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // reset body

	sigBase := fmt.Sprintf("%s:%s:%s", versionPrefix, slackRequesTimestamp, bodyBytes)

	// HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(signingSecret))
	mac.Write([]byte(sigBase))
	expectedSig := versionPrefix + "=" + hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expectedSig), []byte(slackSignature)) {
		err = errors.New("invalid Slack signature")
		log.Err(err).
			Str("received-signature", slackSignature).
			Str("expected-signature", expectedSig).
			Msg("signature mismatch")
		return nil, http.StatusUnauthorized, err
	}

	return bodyBytes, http.StatusOK, nil
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
