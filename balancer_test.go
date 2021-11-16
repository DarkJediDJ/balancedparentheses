package blncdparantheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalancer(t *testing.T) {
	assert.Equal(t, true, balancer("[(){}]"), "Balanced standart parantheses")
	assert.Equal(t, true, balancer("[(dafdsf)&&{fsgdf}bnbhvmhm]"), "Balanced parantheses with symbols")
	assert.Equal(t, false, balancer("[}(])"), "Unbalanced standart parantheses")
	assert.Equal(t, false, balancer("[({dfhdjmm]]]]"), "Unbalanced parantheses with symbols")
	assert.Equal(t, false, balancer(" "), "Empty string")
	assert.NotEqual(t, true, balancer("Hello, World!"), "Simple text string")
}
