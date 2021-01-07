package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}
