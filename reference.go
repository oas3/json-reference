package ref

import (
    "github.com/oas3/json-pointer"
    "net/url"
    "regexp"
    "strings"
)

func New(reference string) (JSONReference, error) {
    refURL, err := url.Parse(reference)
    if err != nil {
        return JSONReference{}, err
    }
    normalize(refURL)

    pointer, _ := ptr.New(refURL.Fragment)
    return JSONReference{
        URL:     refURL,
        Pointer: pointer,
    }, nil
}

type JSONReference struct {
    URL     *url.URL
    Pointer ptr.JSONPointer
}

func normalize(u *url.URL) {
    u.Scheme = strings.ToLower(u.Scheme)
    u.Host = strings.ToLower(u.Host)

    // remove default port
    // http://host:80 -> http://host
    u.Host = regexp.MustCompile(`(:\d+)/?$`).ReplaceAllStringFunc(
        u.Host,
        func(val string) string {
            if (u.Scheme == "http" && val == ":80") ||
                (u.Scheme == "https" && val == ":80") {
                return ""
            }
            return val
        },
    )

    // remove duplicate slashes
    // http://host/path//a///b -> http://host/path/a/b
    u.Path = regexp.MustCompile(`/{2,}`).ReplaceAllString(u.Path, "/")
}
