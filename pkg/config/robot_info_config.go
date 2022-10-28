package config

/**
 * @Author linya.jj
 * @Date 2022/10/19 10:10 AM
 */

// RobotInfoConfig 机器人信息
type RobotInfoConfig struct {
	Code string `yaml:"code" json:"code" env:"ROBOT_CODE" required:"true"`
}

func (c *RobotInfoConfig) GetCode() string {
	if c == nil {
		return ""
	}

	return c.Code
}
