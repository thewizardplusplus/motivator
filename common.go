package main

const (
	markOfShowingStart = "start showing notifications"
)

type configurableCommand struct {
	ConfigPath string `kong:"required,short='c',name='config',default='config.json',help='Config path.'"` // nolint: lll
}
