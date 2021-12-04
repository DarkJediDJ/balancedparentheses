package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancer(t *testing.T) {
	testBalancerCases := []struct {
		value       string
		want        bool
		description string
	}{
		{"[(){}]", true, "Balanced standard parantheses"},
		{"[(dafdsf)&&{fsgdf}bnbhvmhm]", true, "Balanced parantheses with symbols"},
		{"[}(])", false, "Unbalanced standart parantheses"},
		{"[({dfhdjmm]]]]", false, "Unbalanced parantheses with symbols"},
		{" ", true, "Empty string"},
		{"Hello, World!", true, "Simple text string"},
		{"[{)()]", false, "Open parantheses string"},
		{"}{)()]", false, "Close bracket string"},
	}
	for _, tc := range testBalancerCases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, Balancer(tc.value))
		})
	}

}
