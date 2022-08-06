package main

const notificationDisplayStartMark = "start displaying notifications"

type configurableCommand struct {
	ConfigPath string `kong:"required,short='c',name='config',default='config.json',help='The path to a config file.'"` // nolint: lll
}
