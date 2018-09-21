package stringutil

import (
    "testing"
    "fmt"
    "regexp"
    "unicode/utf8"
)

func TestIsEmpty(t *testing.T) {
    fmt.Printf("a%vb\n", IsEmpty("as d       \t"))
}

func TestContainSpace(t *testing.T) {
    fmt.Printf("contain space:%v\n", ContainSpace("ad\td"))

    reg := regexp.MustCompile("[^0-9A-Za-z_]")
    fmt.Printf("account valid:%v\n", reg.MatchString("asd_#") == false)
    reg = regexp.MustCompile("^[A-Za-z]")
    fmt.Printf("account valid2:%v\n", reg.MatchString("asd") == true)

    fmt.Printf("account len:%v\n", utf8.RuneCountInString("asd1浩"))

    fmt.Printf("account special:%v\n", ContainSpecialChar("0123"))
}