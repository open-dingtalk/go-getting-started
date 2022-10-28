package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/23 1:53 PM
 */

func TestNewSuccessApiResponseVO(t *testing.T) {
	assert.NotNil(t, NewSuccessApiResponseVO(""))
}
