package maputil

func Merge(left, right map[string]interface{}) map[string]interface{}  {
    for key, val := range right {
        left[key] = val
    }
    return left
}