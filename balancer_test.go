package blncdparantheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalancer(t *testing.T) {
	testBalancerCases := []struct {
		value       string
		want        bool
		description string
	}{
		{"[(){}]", true, "Balanced standart parantheses"},
		{"[(dafdsf)&&{fsgdf}bnbhvmhm]", true, "Balanced parantheses with symbols"},
		{"[}(])", false, "Unbalanced standart parantheses"},
		{"[({dfhdjmm]]]]", false, "Unbalanced parantheses with symbols"},
		{" ", true, "Empty string"},
		{"Hello, World!", true, "Simple text string"},
	}
	for _, tc := range testBalancerCases {
		assert.Equal(t, tc.want, balancer(tc.value), tc.description)
	}

}
