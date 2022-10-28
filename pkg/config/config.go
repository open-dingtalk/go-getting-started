package config

import (
	"github.com/jinzhu/configor"
	"github.com/pkg/errors"
)

/**
 * @Author linya.jj
 * @Date 2022/10/9 11:17 AM
 */

// Config 项目启动配置文件
type Config struct {
	InnerAppInfo   *InnerAppInfoConfig `yaml:"innerappinfo" json:"innerAppInfo" required:"true"`
	RobotInfo      *RobotInfoConfig    `yaml:"robotinfo" json:"robotInfo" required:"true"`
	CoolAppInfo001 *CoolAppInfoConfig  `yaml:"coolappinfo001" json:"coolAppInfo001" required:"true"`
	CardInfo       *CardInfoConfig     `yaml:"cardinfo" json:"cardInfo" required:"true"`
}

func (c *Config) GetInnerAppInfo() *InnerAppInfoConfig {
	if c == nil {
		return nil
	}

	return c.InnerAppInfo
}

func (c *Config) GetRobotInfo() *RobotInfoConfig {
	if c == nil {
		return nil
	}

	return c.RobotInfo
}

func (c *Config) GetCoolAppInfo() *CoolAppInfoConfig {
	if c == nil {
		return nil
	}

	return c.CoolAppInfo001
}

func (c *Config) GetCardInfo() *CardInfoConfig {
	if c == nil {
		return nil
	}

	return c.CardInfo
}

var (
	config *Config
)

func GetConfig() *Config {
	return config
}

// ActiveInitConfigWithFile 初始化配置文件
func ActiveInitConfigWithFile(filename string) error {
	config = &Config{}
	err := configor.Load(config, filename)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
