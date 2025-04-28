package browser

import (
	"errors"
	"os"
	"os/exec"

	"github.com/pkg/browser"
)

func (b *Browser) OpenUrl(url string) error {
	err := browser.OpenURL(url)
	if err != nil && errors.Is(err, exec.ErrNotFound) {
		// github.com/pkg/browserはWSLをサポートしていない。
		// wslではwsl-openを使っている場合ここで開く
		cmds := exec.Command("wsl-open", url)
		cmds.Stdin = os.Stdin
		cmds.Stdout = os.Stdout
		cmds.Stderr = os.Stderr
		return cmds.Run()
	}
	return nil
}
