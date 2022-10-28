package config

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:10 AM
 */

// InnerAppInfoConfig 自建应用的凭证信息
type InnerAppInfoConfig struct {
	AgentId   int64  `yaml:"agentid" json:"agentId" env:"AGENT_ID" required:"false"`
	AppKey    string `yaml:"appkey" json:"appKey" env:"APP_KEY" required:"true"`
	AppSecret string `yaml:"appsecret" json:"appSecret" env:"APP_SECRET" required:"true"`
}

func (c *InnerAppInfoConfig) GetAppKey() string {
	if c == nil {
		return ""
	}

	return c.AppKey
}

func (c *InnerAppInfoConfig) GetAppSecret() string {
	if c == nil {
		return ""
	}

	return c.AppSecret
}

func (c *InnerAppInfoConfig) GetAgentId() int64 {
	if c == nil {
		return 0
	}

	return c.AgentId
}
