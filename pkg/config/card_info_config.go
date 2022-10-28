package config

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:14 AM
 */

// CardInfoConfig 卡片信息配置
type CardInfoConfig struct {
	MessageCardTemplateId001 string `yaml:"messagecardtemplateid001" json:"messageCardTemplateId001" env:"MESSAGE_CARD_TEMPLATE_ID" required:"true"`
	TopCardTemplateId001     string `yaml:"topcardtemplateid001" json:"topCardTemplateId001" env:"TOP_CARD_TEMPLATE_ID" required:"true"`
}

func (c *CardInfoConfig) GetMessageCardTemplateId() string {
	if c == nil {
		return ""
	}

	return c.MessageCardTemplateId001
}

func (c *CardInfoConfig) GetTopCardTemplateId() string {
	if c == nil {
		return ""
	}

	return c.TopCardTemplateId001
}
