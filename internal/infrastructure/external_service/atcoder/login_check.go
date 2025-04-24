package atcoder

import (
	"fmt"
	"net/http"
)

func (a *Atcoder) LoginCheck() (bool, error) {
	path := "/contests/abc001/submit"
	resp, err := a.get(path)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusFound {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if resp.Request.URL.Path != path {
		return false, nil
	}

	if resp.StatusCode == http.StatusFound {
		return false, nil
	}

	return true, nil
}
