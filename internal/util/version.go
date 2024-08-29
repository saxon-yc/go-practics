package util

import (
	"go-practics/pkg/version"
	"regexp"
)

// if cur > dst, return true; else(<=) return false
func CompareVersionStr(cur, dst string) (bool, error) {
	vCur, err := version.ParseSemantic(cur)
	if err != nil {
		return false, err
	}

	result, err := vCur.Compare(dst)
	if err != nil {
		return false, err
	}

	if result == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func Compare(cur, dst string) (int, error) {
	vCur, err := version.ParseGeneric(cur)
	if err != nil {
		return 0, err
	}

	return vCur.Compare(dst)
}

func LatestSupportLicenseVersion(v string) bool {
	if v == "" {
		return false
	}
	generic, _ := version.ParseGeneric(v)
	return generic.AtLeast(version.MustParseGeneric("v4.1.0"))
}

// StandardizeVersion 解析输入字符串并返回 VersionInfo
func StandardizeVersion(input string) string {
	// 使用正则表达式提取版本号
	re := regexp.MustCompile(`(?:(?:[a-zA-Z]+:)?v?(\d+\.\d+\.\d+)(?:-[a-zA-Z0-9]+(?:-\d+)?)?)`)
	matches := re.FindStringSubmatch(input)

	if len(matches) > 1 {
		// 返回标准版本格式
		return "v" + matches[1]
	}

	// 如果没有匹配到，返回空字符串或原始输入
	return ""
}
