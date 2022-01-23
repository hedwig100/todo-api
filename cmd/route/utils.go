package route

import (
	"errors"
	"net/url"
	"strings"
)

func isCorrectURL(prefix string, url *url.URL) (trailing string, err error) {
	if strings.HasPrefix(url.Path, prefix) {
		trailing = strings.TrimPrefix(url.Path, prefix)
	} else {
		err = errors.New("uri is not valid")
	}
	return
}
