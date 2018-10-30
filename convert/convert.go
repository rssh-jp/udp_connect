package convert
import(
    "fmt"
    "strings"
)

func MapInterface2string(src interface{})string{
    switch val := src.(type){
    case map[string]interface{}:
        return Map2string(val)
    }
    return ""
}
func Map2string(src map[string]interface{})string{
    list := make([]string, 0, 10)
    for key, val := range src{
        list = append(list, fmt.Sprintf(`"%s":%s`, key, interface2string(val)))
    }
    return fmt.Sprintf("{%s}", strings.Join(list, ","))
}

func array2string(src []interface{})string{
    list := make([]string, 0, 10)
    for _, val := range src{
        list = append(list, fmt.Sprintf(`%s`, interface2string(val)))
    }
    return fmt.Sprintf("[%s]", strings.Join(list, ","))
}

func interface2string(src interface{})string{
    var ret string
    switch v:= src.(type){
    case map[string]interface{}:
        ret = fmt.Sprintf("%s", Map2string(v))
    case []interface{}:
        ret = fmt.Sprintf("%s", array2string(v))
    case string:
        ret = fmt.Sprintf(`"%s"`, v)
    case float64:
        ret = fmt.Sprintf("%v", v)
    }
    return ret
}

