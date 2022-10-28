package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:10 AM
 */

func TestInnerAppInfoConfig_GetAppKey_IfNotNil(t *testing.T) {
	c := &InnerAppInfoConfig{
		AppKey: "AppKey",
	}
	assert.Equal(t, "AppKey", c.GetAppKey())
}

func TestInnerAppInfoConfig_GetAppKey_IfNil(t *testing.T) {
	var c *InnerAppInfoConfig
	assert.Equal(t, "", c.GetAppKey())
}

func TestInnerAppInfoConfig_GetAppSecret_IfNotNil(t *testing.T) {
	c := &InnerAppInfoConfig{
		AppSecret: "AppSecret",
	}
	assert.Equal(t, "AppSecret", c.GetAppSecret())
}

func TestInnerAppInfoConfig_GetAppSecret_IfNil(t *testing.T) {
	var c *InnerAppInfoConfig
	assert.Equal(t, "", c.GetAppSecret())
}

func TestInnerAppInfoConfig_GetAgentId_IfNotNil(t *testing.T) {
	c := &InnerAppInfoConfig{
		AgentId: 1,
	}
	assert.Equal(t, int64(1), c.GetAgentId())
}

func TestInnerAppInfoConfig_GetAgentId_IfNil(t *testing.T) {
	var c *InnerAppInfoConfig
	assert.Equal(t, int64(0), c.GetAgentId())
}
