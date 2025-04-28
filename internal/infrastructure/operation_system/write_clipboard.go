package operation_system

import "golang.design/x/clipboard"

func (o *OperationSystem) WriteClipboard(data []byte) error {
	if err := clipboard.Init(); err != nil {
		return err
	}
	clipboard.Write(clipboard.FmtText, data)
	return nil
}
