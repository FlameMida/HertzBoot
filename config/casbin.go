package config

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"` // 存放casbin模型的相对路径
	ApiLevel  bool   `mapstructure:"api-level" json:"apiLevel" yaml:"api-level"`    // api粒度的鉴权
}
