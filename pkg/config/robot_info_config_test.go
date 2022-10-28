package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:10 AM
 */

func TestRobotInfoConfig_GetCode_IfNotNil(t *testing.T) {
	c := &RobotInfoConfig{
		Code: "Code",
	}

	assert.Equal(t, "Code", c.GetCode())
}

func TestRobotInfoConfig_GetCode_IfNil(t *testing.T) {
	var c *RobotInfoConfig

	assert.Equal(t, "", c.GetCode())
}
