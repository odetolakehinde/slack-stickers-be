// Package helper defines helper constants and functions used across other packages in the application
package helper

import (
	"bytes"
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	// SortOrder struct
	SortOrder string
	// Key is a middleware key sting value
	Key string
)

const (
	// ZeroUUID default empty or non set UUID value
	ZeroUUID = "00000000-0000-0000-0000-000000000000"
	// LogStrKeyModule log service name value
	LogStrKeyModule = "ser_name"
	// LogStrKeyLevel log service level value
	LogStrKeyLevel = "lev_name"
	// LogStrPartnerLevel log partner name value
	LogStrPartnerLevel = "partner_name"
	// LogStrKeyMethod log method name value
	LogStrKeyMethod = "method_name"
	// LogStrKeyEndpointName log endpoint name value
	LogStrKeyEndpointName = "endpoint_name"
	// LogEndpointLevel log endpoint value
	LogEndpointLevel = "endpoint"
	// LogStrRequestIDLevel log request id value
	LogStrRequestIDLevel = "request-id"
	// LogStrPayloadLevel log payload value
	LogStrPayloadLevel = "payload"
	// LogStrPackageLevel log package name value
	LogStrPackageLevel = "package_name"
	// LogStrResponseLevel log response level
	LogStrResponseLevel = "response"
	// SortOrderASC for ascending sorting
	SortOrderASC SortOrder = "ASC"
	// SortOrderDESC for descending sorting
	SortOrderDESC SortOrder = "DESC"
	// GinContextKey constant that holds the Gin context key
	GinContextKey Key = "SlACK_STICKERS_GinContextKey"
)

// GinContextFromContext gets a gin context from a context.Context
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey)
	if ginContext == nil {
		return nil, ErrGinContextRetrieveFailed
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		return nil, ErrGinContextWrongType
	}

	return gc, nil
}

// GetStringPointer returns a string pointer
func GetStringPointer(val string) *string {
	return &val
}

// GetTimePointer returns a time pointer
func GetTimePointer(time time.Time) *time.Time {
	return &time
}

// GetStringVal return string from pointer
func GetStringVal(strVal *string) string {
	var val string
	if strVal != nil {
		return *strVal
	}
	return val
}

// GetFloat64Pointer returns a float64 pointer
func GetFloat64Pointer(val float64) *float64 {
	return &val
}

// GetFloatVal returns valid float val from pointer
func GetFloatVal(floatVal *float64) float64 {
	var val float64
	if floatVal != nil {
		return *floatVal
	}
	return val
}

// GetBoolPointer returns a bool pointer
func GetBoolPointer(val bool) *bool {
	return &val
}

// GetIntPointer returns an int pointer
func GetIntPointer(val int) *int {
	return &val
}

// GetUUIDPointer returns a UUID pointer
func GetUUIDPointer(val uuid.UUID) *uuid.UUID {
	return &val
}

// GetUUIDVal returns a uuid
func GetUUIDVal(uuidVal *uuid.UUID) uuid.UUID {
	var val uuid.UUID
	if uuidVal != nil {
		return *uuidVal
	}
	return val
}

// StreamToByte converts an io Stream to a slice of byte
func StreamToByte(stream io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(stream)
	//if err == nil {
	//	return []byte{}, err
	//}
	return buf.Bytes(), nil
}

// SortMap helper to sort a Map
func SortMap(m map[string]interface{}) map[string]interface{} {
	ret := map[string]interface{}{}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		ret[k] = m[k]
	}

	return ret
}

// StringToTime system default helper to convert a possible dateTime string to time.Time
func StringToTime(s string) (time.Time, error) {
	return time.Parse("02/01/2006", s)
}

// StringToDatetime converts date string with format "mm/dd/yyyy hh:mm"  to time.Time
func StringToDatetime(s string) (time.Time, error) {
	return time.Parse("01/02/2006 3:04 PM", s)
}

// DatePartFromString trims off the extra "T00:00:00Z" at the end of the DueDate string
func DatePartFromString(date *string) *string {
	if date != nil && len(*date) >= 10 {
		dateString := *date
		date := dateString[:10]
		return &date
	}
	return nil
}

// GetTimeVal converts time to string
func GetTimeVal(timeVal *time.Time) time.Time {
	var t time.Time
	if timeVal != nil {
		return *timeVal
	}
	return t
}

// RandomNumbers generates random numerics with length specified
func RandomNumbers(max int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}

// TimeAfterNowBy allows you to +/- a designated duration to the current time.Time
func TimeAfterNowBy(by time.Duration, arith int) time.Time {
	return time.Now().Add(by * time.Duration(arith))
}

// ToSQLString NullString converter
func ToSQLString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// ToSQLInt32 NullInt32 converter
func ToSQLInt32(i int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: true,
	}
}

// ToSQLFloat64 NullFloat64 converter
func ToSQLFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: f,
		Valid:   true,
	}
}

// GenerateKey generate random key
func GenerateKey(max int) (string, error) {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	b := make([]byte, max)
	_, err := io.ReadAtLeast(rand.Reader, b, max)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

// GenerateRandomDigits generate random digits
func GenerateRandomDigits(max int) (string, error) {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	b := make([]byte, max)
	_, err := io.ReadAtLeast(rand.Reader, b, max)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

// GetDateString converts date to string
func GetDateString(t time.Time) string {
	switch {
	case t.Month() <= 9 && t.Day() <= 9:
		return fmt.Sprintf("%d-0%d-0%d", t.Year(), t.Month(), t.Day())
	case t.Month() <= 9:
		return fmt.Sprintf("%d-0%d-%d", t.Year(), t.Month(), t.Day())
	case t.Day() <= 9:
		return fmt.Sprintf("%d-%d-0%d", t.Year(), t.Month(), t.Day())
	default:
		return fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	}
}

// ReformatDateString converts string date to a specified format
func ReformatDateString(date, initialSeparator, resultSeparator string) string {
	val := strings.Split(date, initialSeparator)
	return fmt.Sprintf("%s%s%s%s%s", val[2], resultSeparator, val[1], resultSeparator, val[0])
}

// StringToInt helps to convert string to int
func StringToInt(value string) int {
	var newInt int
	if i, err := strconv.Atoi(value); err == nil {
		newInt = i
	}
	return newInt
}

// StringToIntPointer helps to convert string to int pointer
func StringToIntPointer(value string) *int {
	var newValue int
	if i, err := strconv.Atoi(value); err == nil {
		newValue = i
	}
	return &newValue
}

// StringToBool helps to convert string to bool
func StringToBool(value string) bool {
	var newValue bool
	if i, err := strconv.ParseBool(value); err == nil {
		newValue = i
	}
	return newValue
}

// StringToBoolPointer helps to convert string to bool pointer
func StringToBoolPointer(value string) *bool {
	var newValue bool
	if i, err := strconv.ParseBool(value); err == nil {
		newValue = i
	}
	return &newValue
}
