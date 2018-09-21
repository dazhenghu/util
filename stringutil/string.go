package stringutil

import "strings"

/**
判断是否为空字符串
 */
func IsEmpty(str string) bool {
    res := strings.TrimSpace(str)
    if len(res) == 0 {
        return true
    }
    return false
}
