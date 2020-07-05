package main

import "strings"

func simplifyPath(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	pathSlice := strings.Split(path, "/")
	result := []string{}
	for _, item := range pathSlice {
		if item == ".." {
			if len(result) == 0 {
				continue
			}
			result = result[:len(result)-1]
		} else if item == "." || item == "" {
			continue
		} else {
			result = append(result, item)
		}
	}

	return "/" + strings.Join(result, "/")
}
