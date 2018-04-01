package utilities

import (
    "reflect"
    "errors"
)

// GetStructTags returns struct attribute name as key and tag as value
func GetStructTags(v interface{}, tagKey string, includeChildTags bool) (result map[string]string, err error) {
    sValue := reflect.ValueOf(v)
    sType := reflect.TypeOf(v)
    if sValue.Kind() == reflect.Ptr && !sValue.IsNil() {
        return GetStructTags(sValue.Elem().Interface(), tagKey, includeChildTags)
    }
    if sValue.Kind() != reflect.Struct && sValue.Kind() != reflect.Interface {
        err = errors.New("Must be a Struct Interface")
        return
    }

    result = map[string]string{}
    numField := sValue.NumField()
    for i := 0; i < numField; i++ {
        name := sType.Field(i).Name
        tagValue := sType.Field(i).Tag.Get(tagKey)
        if sValue.Field(i).Kind() == reflect.Struct && includeChildTags {
            values := map[string]string{}
            values, err = GetStructTags(sValue.Field(i).Interface(), tagKey, includeChildTags)
            if err != nil {
                return
            }
            for k, v := range values {
                result[k] = v
                result[name + "." + k] = v
            }
        } else {
            result[name] = tagValue
        }
    }
    return
}
