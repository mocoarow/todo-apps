package handler_test

import (
	"bytes"
	"io"
	"math/rand"
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func randomUserID() int {
	return rand.Intn(1000000) + 1 //nolint:gosec
}

func readBytes(t *testing.T, b *bytes.Buffer) []byte {
	t.Helper()
	respBytes, err := io.ReadAll(b)
	require.NoError(t, err)
	return respBytes
}

func parseJSON(t *testing.T, bytes []byte) interface{} {
	t.Helper()
	obj, err := oj.Parse(bytes)
	require.NoError(t, err)
	return obj
}

func parseExpr(t *testing.T, v string) jp.Expr {
	t.Helper()
	expr, err := jp.ParseString(v)
	require.NoError(t, err)
	return expr
}

func validateErrorResponse(t *testing.T, respBytes []byte, expectedErrorCode string, expectedErrorMessage string) {
	t.Helper()

	jsonObj := parseJSON(t, respBytes)

	// - error code
	errorCodeExpr := parseExpr(t, "$.code")
	errorCode := errorCodeExpr.Get(jsonObj)
	require.Len(t, errorCode, 1, "response should have one code: %+v", jsonObj)
	assert.Equal(t, expectedErrorCode, errorCode[0])

	// - error message
	errorMessageExpr := parseExpr(t, "$.message")
	errorMessage := errorMessageExpr.Get(jsonObj)
	require.Len(t, errorMessage, 1, "response should have one message: %+v", jsonObj)
	assert.Equal(t, expectedErrorMessage, errorMessage[0])
}
