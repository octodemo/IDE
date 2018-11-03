package methods

import (
	"reflect"

	"github.com/thestrukture/IDE/types"
)

//
func GetApp(args ...interface{}) types.App {
	apps := args[0]
	name := args[1]

	s := reflect.ValueOf(apps)
	slice := make([]types.App, s.Len())
	for i, _ := range slice {
		v := s.Index(i).Interface().(types.App)

		if v.Name == name.(string) {
			return v
		}
	}

	return types.App{}

}