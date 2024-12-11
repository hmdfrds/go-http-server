package utils

import (
	"regexp"
)

var URLUnescapeMap = map[string]string{
	"%20": " ",
	"%21": "!",
	"%22": "\"",
	"%23": "#",
	"%24": "$",
	"%25": "%",
	"%26": "&",
	"%27": "'",
	"%28": "(",
	"%29": ")",
	"%2A": "*",
	"%2B": "+",
	"%2C": ",",
	"%2D": "-",
	"%2E": ".",
	"%2F": "/",
	"%3A": ":",
	"%3B": ";",
	"%3C": "<",
	"%3D": "=",
	"%3E": ">",
	"%3F": "?",
	"%40": "@",
	"%5B": "[",
	"%5C": "\\",
	"%5D": "]",
	"%5E": "^",
	"%5F": "_",
	"%60": "`",
	"%7B": "{",
	"%7C": "|",
	"%7D": "}",
	"%7E": "~",
}

// can use url.QueryUnescape()
func UnescapeString(query string) string {
	re := regexp.MustCompile(`(%\w\w)`)

	return re.ReplaceAllStringFunc(query, func(match string) string {
		if unescaped, ok := URLUnescapeMap[match]; ok {
			return unescaped
		}
		return match
	})
}
