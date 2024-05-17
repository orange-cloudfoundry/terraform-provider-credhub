package credhub

import "fmt"

func validateFromMap(mapValid map[string]bool, keyType string) func(elem interface{}, index string) ([]string, []error) {
	return func(elem interface{}, index string) ([]string, []error) {
		if _, ok := mapValid[elem.(string)]; !ok {
			return make([]string, 0), []error{fmt.Errorf("the provided %s is not supported. Valid values include %s", keyType, validateMapToString(mapValid))}
		}
		return make([]string, 0), []error{}
	}
}
