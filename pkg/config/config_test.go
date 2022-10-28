package config

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 11:17 AM
 */

func TestConfig_GetInnerAppInfo_IfNotNil(t *testing.T) {
	innerAppConfig := &InnerAppInfoConfig{}
	c := &Config{
		InnerAppInfo: innerAppConfig,
	}
	assert.Equal(t, innerAppConfig, c.GetInnerAppInfo())
}

func TestConfig_GetInnerAppInfo_IfNil(t *testing.T) {
	var c *Config
	assert.Nil(t, c.GetInnerAppInfo())
}

func TestConfig_GetRobotInfo_IfNotNil(t *testing.T) {
	robotInfoConfig := &RobotInfoConfig{}
	c := &Config{
		RobotInfo: robotInfoConfig,
	}
	assert.Equal(t, robotInfoConfig, c.GetRobotInfo())
}

func TestConfig_GetRobotInfo_IfNil(t *testing.T) {
	var c *Config
	assert.Nil(t, c.GetRobotInfo())
}

func TestConfig_GetCoolAppInfo_IfNotNil(t *testing.T) {
	coolAppConfig := &CoolAppInfoConfig{}
	c := &Config{
		CoolAppInfo001: coolAppConfig,
	}
	assert.Equal(t, coolAppConfig, c.GetCoolAppInfo())
}

func TestConfig_GetCoolAppInfo_IfNil(t *testing.T) {
	var c *Config
	assert.Nil(t, c.GetCoolAppInfo())
}

func TestConfig_GetCardInfo_IfNotNil(t *testing.T) {
	cardInfoConfig := &CardInfoConfig{}
	c := &Config{
		CardInfo: cardInfoConfig,
	}
	assert.Equal(t, cardInfoConfig, c.GetCardInfo())
}

func TestConfig_GetCardInfo_IfNil(t *testing.T) {
	var c *Config
	assert.Nil(t, c.GetCardInfo())
}

func TestActiveInitConfigWithFile_IfSuccess(t *testing.T) {
	config = nil
	assert.Nil(t, GetConfig())

	filePath := "config_test.yaml"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	data := "innerappinfo:\n  appkey: your-app-key\n  appsecret: your-app-secret\nrobotinfo:\n  code: your-robot-code\ncoolappinfo001:\n  code: your-cool-app-code\ncardinfo:\n  messagecardtemplateid001: your-message-card-template-id\n  topcardtemplateid001: your-top-card-template-id"

	write := bufio.NewWriter(file)
	write.WriteString(data)
	write.Flush()
	file.Close()

	err = ActiveInitConfigWithFile(filePath)

	os.Remove(filePath)

	assert.Nil(t, err)
}

func TestActiveInitConfigWithFile_IfFail(t *testing.T) {
	filePath := "config_test.yaml"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	data := "innerappinfo:\n  appSecret: your-app-secret\nrobotinfo:\n  code: your-robot-code\ncoolappinfo001:\n  code: your-cool-app-code\ncardinfo:\n  messagecardtemplateid001: your-message-card-template-id\n  topcardtemplateid001: your-top-card-template-id"

	write := bufio.NewWriter(file)
	write.WriteString(data)
	write.Flush()
	file.Close()

	err = ActiveInitConfigWithFile(filePath)

	os.Remove(filePath)

	assert.NotNil(t, err)
}
