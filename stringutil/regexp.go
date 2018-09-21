package stringutil

import "regexp"

// 匹配空字符串
var RegSpace *regexp.Regexp = regexp.MustCompile("\\s")

// 匹配特殊字符
var RegSpecial *regexp.Regexp = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]|]`)

// 邮箱
var RegMail = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

/**
校验是否包含空白字符
 */
func ContainSpace(str string) bool {
    return RegSpace.MatchString(str)
}

/**
校验是否包含特殊字符
 */
func ContainSpecialChar(str string) bool {
    return RegSpecial.MatchString(str)
}

/**
校验是否是有效的邮箱格式
 */
func ValidMail(mail string) bool {
    return RegMail.MatchString(mail)
}