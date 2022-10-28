package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:12 AM
 */

func TestCoolAppInfoConfig_GetCode_IfNotNil(t *testing.T) {
	c := &CoolAppInfoConfig{
		Code: "Code",
	}

	assert.Equal(t, "Code", c.GetCode())
}

func TestCoolAppInfoConfig_GetCode_IfNil(t *testing.T) {
	var c *CoolAppInfoConfig
	assert.Equal(t, "", c.GetCode())
}
