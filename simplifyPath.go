package main

import "strings"

func simplifyPath(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	pathSlice := strings.Split(path, "/")
	result := []string{}
	for _, val := range pathSlice {
		if val == "." || val == "" {
			continue
		}
		if val == ".." {
			if len(result) <= 0 {
				continue
			}
			result = result[:len(result)-1]
			continue
		}

		result = append(result, val)
	}
	return "/" + strings.Join(result, "/")
}
