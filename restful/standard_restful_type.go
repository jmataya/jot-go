package restful

import "regexp"

type StandardRestfulType struct {
	basePath string
}

func (s *StandardRestfulType) SetBasePath(path string) {
	s.basePath = path
}

func (s StandardRestfulType) GetBasePath() string {
	return s.basePath
}

func (StandardRestfulType) getUrlParamRegex() string {
	return "[a-zA-Z0-9_-]+"
}

func (s StandardRestfulType) getCollectionPath() string {
	/*
	 * This essentially works by looking at the base path and stripping off an
	 * ID parameter if it exists on the end of the string.
	 */
	lastIdRegexStr := "{[a-zA-Z0-9_-]+}/?$"
	lastIdRegexMatcher := regexp.MustCompile(lastIdRegexStr)
	return lastIdRegexMatcher.ReplaceAllString(s.basePath, "")
}

func (StandardRestfulType) pathIsMatch(base string, actual string) bool {
	// First, interpolate the placeholders that are in for the strings.
	keyMatcher := regexp.MustCompile("{[a-zA-Z0-9_-]+}")
	interpolatedStr := keyMatcher.ReplaceAllString(base, "[a-zA-Z0-9_-]+")

	// Second, clean up the string:
	// 1. Make sure that the last "/" is optional
	// 2. Make sure that nothing can come after this string.
	slashRegexMatcher := regexp.MustCompile("/$")
	interpolatedStr = slashRegexMatcher.ReplaceAllString(interpolatedStr, "")
	interpolatedStr += "/?$"

	// Finally, do the actual match
	valueMatcher := regexp.MustCompile(interpolatedStr)
	return valueMatcher.MatchString(actual)
}

func (s StandardRestfulType) isCollectionMatch(path string) bool {
	collectionPath := s.getCollectionPath()
	return s.pathIsMatch(collectionPath, path)
}

func (s StandardRestfulType) isMemberMatch(path string) bool {
	return s.pathIsMatch(s.basePath, path)
}

func (s StandardRestfulType) ActionMatch(path string, method string) (bool, string) {
	if s.isCollectionMatch(path) {
		if method == GET {
			return true, LIST
		} else if method == POST {
			return true, CREATE
		} else {
			return false, ""
		}
	} else if s.isMemberMatch(path) {
		if method == GET {
			return true, SHOW
		} else if method == PUT {
			return true, UPDATE
		} else if method == DELETE {
			return true, DESTROY
		} else {
			return false, ""
		}
	} else {
		return false, ""
	}
}
