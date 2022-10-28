package config

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:12 AM
 */

// CoolAppInfoConfig 酷应用信息
type CoolAppInfoConfig struct {
	Code string `yaml:"code" json:"code" env:"COOL_APP_CODE" required:"true"`
}

func (c *CoolAppInfoConfig) GetCode() string {
	if c == nil {
		return ""
	}

	return c.Code
}
