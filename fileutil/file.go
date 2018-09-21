package fileutil

import (
    "os"
    "path/filepath"
    "strings"
)

/**
判断对应路径是否存在
 */
func PathExists(path string) (bool, error)  {
    // 获取path的信息
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return false, err
}

/**
读取当前文件夹
 */
func GetCurrentDirectory() (string, error) {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        return "", err
    }
    return strings.Replace(dir, "\\", "/", -1), nil
}
