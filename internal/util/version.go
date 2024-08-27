package util

import "go-practics/pkg/version"

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
