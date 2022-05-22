package tinygo

import "path"

func joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}
	finalPath := path.Join(absolutePath, relativePath)
	if lastChar(relativePath) == '/' && lastChar(finalPath) != '/' {
		finalPath += "/"
	}
	return finalPath
}

func lastChar(str string) byte {
	return str[len(str)-1]
}
