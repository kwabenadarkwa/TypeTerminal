package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRandomQuoteAPI(t *testing.T) {
	result, err := GetRandomQuote()
	require.NoError(t, err)
	require.NotNil(t, result)
}
