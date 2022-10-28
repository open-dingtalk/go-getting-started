package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:14 AM
 */

func TestCardInfoConfig_GetMessageCardTemplateId_IfNotNil(t *testing.T) {
	c := &CardInfoConfig{
		MessageCardTemplateId001: "MessageCardTemplateId001",
	}

	assert.Equal(t, "MessageCardTemplateId001", c.GetMessageCardTemplateId())
}

func TestCardInfoConfig_GetMessageCardTemplateId_IfNil(t *testing.T) {
	var c *CardInfoConfig
	assert.Equal(t, "", c.GetMessageCardTemplateId())
}

func TestCardInfoConfig_GetTopCardTemplateId_IfNotNil(t *testing.T) {
	c := &CardInfoConfig{
		TopCardTemplateId001: "TopCardTemplateId001",
	}

	assert.Equal(t, "TopCardTemplateId001", c.GetTopCardTemplateId())
}

func TestCardInfoConfig_GetTopCardTemplateId_IfNil(t *testing.T) {
	var c *CardInfoConfig
	assert.Equal(t, "", c.GetTopCardTemplateId())
}
