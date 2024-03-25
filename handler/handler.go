package handler

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(urlsToPaths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := urlsToPaths[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urlsPath []UrlsPath
	err := yaml.Unmarshal(yml, &urlsPath)
	if err != nil {
		return nil, err
	}

	pathsUrlMap := make(map[string]string)
	for _, pu := range(urlsPath) {
		pathsUrlMap[pu.Path] = pu.Url
	}

	return MapHandler(pathsUrlMap, fallback), nil
}

type UrlsPath struct {
	Path string `yaml: "path"`
	Url string `yaml:"url"`
}
