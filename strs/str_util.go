package strings

import (
	"strconv"
	"strings"
)

func StrToInt(str string) (int, error) {
	int, err := strconv.Atoi(str)
	return int, err
}

func StrToInt64(str string) (int64, error) {
	int64, err := strconv.ParseInt(str, 10, 64)
	return int64, err
}

func IntToStr(int int) string {
	string := strconv.Itoa(int)
	return string
}

func Int64ToStr(int64 int64) string {
	string := strconv.FormatInt(int64, 10)
	return string
}

func StrTofloat32(str string) (float32, error) {
	float64, err := strconv.ParseFloat(str, 32)
	return float32(float64), err
}

func StrTofloat64(str string) (float64, error) {
	float64, err := strconv.ParseFloat(str, 64)
	return float64, err
}

func GetDigitFromStr(str string) string {
	if str == "" {
		return ""
	}
	strs := []string{}
	for _, char := range []byte(str) {
		if char >= 48 && char <= 57 {
			strs = append(strs, string(char))
		}
	}
	return strings.Join(strs, "")
}
