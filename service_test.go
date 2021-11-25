package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_returnSingleArticle(t *testing.T) {
	testServiceCases := []struct {
		value       int
		want        int
		description string
	}{
		{5, 5, "Standard parantheses string"},
		{1000, 1000, "Long parantheses string"},
		{1, 1, "Short parantheses string"},
		{0, 0, "Emptyy parantheses string"},
	}
	for _, tc := range testServiceCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, len([]rune(GetValue(tc.value))))
		})
	}

	testServiceStatusCases := []struct {
		value       int
		want        int
		description string
	}{
		{5, 200, "Standard parantheses string"},
		{1000, 200, "Long parantheses string"},
		{1, 200, "Short parantheses string"},
		{0, 200, "Emptyy parantheses string"},
	}
	for _, tc := range testServiceStatusCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, GetStatus(tc.value))
		})
	}
}
