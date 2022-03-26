package main

type stopCommand struct{}

func (command stopCommand) Run() error {
	return killBackgroundProcess("motivator")
}
