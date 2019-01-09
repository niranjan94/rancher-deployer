package utils

import (
	"github.com/spf13/viper"
	"reflect"
)

//noinspection GoUnusedExportedFunction
func GetSliceInterfaceAsMaps(args interface{}) []map[string]string {
	s := reflect.ValueOf(args)
	output := make([]map[string]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		mapValue := s.Index(i).Interface().(map[interface{}]interface{})
		stringMap := make(map[string]string)
		for key, value := range mapValue {
			switch key := key.(type) {
			case string:
				switch value := value.(type) {
				case string:
					stringMap[key] = value
				}
			}
		}
		output[i] = stringMap
	}
	return output
}

func GetSliceInterfaceAsSubs(args interface{}) []*viper.Viper {
	s := reflect.ValueOf(args)
	output := make([]*viper.Viper, s.Len())
	for i := 0; i < s.Len(); i++ {
		mapValue := s.Index(i).Interface()
		subViper := viper.New()
		subViper.SetDefault("intermediate", mapValue)
		output[i] = subViper.Sub("intermediate")
	}
	return output
}
