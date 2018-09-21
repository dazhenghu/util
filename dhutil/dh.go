package dhutil

import "time"

const (
    TIME_FORMAT_MIDDLE_SPLIT int8 = 1 + iota // 格式如：2006-01-02 15:04:05 或 2006-01-02
    TIME_FORMAT_NO_SPLIT // 格式如：20060102150405 或 20060102
    TIME_FORMAT_UNDER_SPLIT // 格式如：2006_01_02_15_04_05 或 2006_01_02
)

/**
时间
 */
func CurrTimeFormat(format int8) string {
    if format == TIME_FORMAT_NO_SPLIT {
        return time.Now().Format("20060102150405")
    } else if format == TIME_FORMAT_UNDER_SPLIT {
        return time.Now().Format("2006_01_02_15_04_05")
    }
    return time.Now().Format("2006-01-02 15:04:05")
}

/**
时间
 */
func TimeFormat(datetime time.Time, format int8) string {
    if format == TIME_FORMAT_NO_SPLIT {
        return datetime.Format("20060102150405")
    } else if format == TIME_FORMAT_UNDER_SPLIT {
        return datetime.Format("2006_01_02_15_04_05")
    }
    return datetime.Format("2006-01-02 15:04:05")
}

/**
默认时间格式化
 */
func TimeFormatDefault(datetime time.Time) string {
    return datetime.Format("2006-01-02 15:04:05")
}

/**
日期
 */
func CurrDateFormat(format int8) string {
    if format == TIME_FORMAT_NO_SPLIT {
        return time.Now().Format("20060102")
    } else if format == TIME_FORMAT_UNDER_SPLIT {
        return time.Now().Format("2006_01_02")
    }
    return time.Now().Format("2006-01-02")
}
