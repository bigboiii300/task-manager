package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
