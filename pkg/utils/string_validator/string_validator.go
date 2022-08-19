package stringvalidator

import "regexp"

func IsKebab(s string) bool {
	return regexp.MustCompile(`^[a-z0-9-]*$`).MatchString(s)
}

func IsSemVer(s string) bool {
	return regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`).MatchString(s)
}

func IsEmail(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_.]+[@]{1}[a-z0-9]+[\.][a-z]+$`).MatchString(s)
}
