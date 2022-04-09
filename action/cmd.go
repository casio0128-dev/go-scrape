package action

import (
	"github.com/sclevine/agouti"
	"os/exec"
	"strings"
)

type CmdAction struct {
	name    string
	command string
}

func NewCmdAction(name string, command string) *CmdAction {
	return &CmdAction{name: name, command: command}
}

func (ca *CmdAction) Name() string {
	return ca.name
}

func (ca *CmdAction) Do(_ *agouti.Page) error {
	if ca.IsActual() {
		return exec.Command(ca.command).Run()
	}
	return NotActualFormat(ca.name)
}

func (ca *CmdAction) IsActual() bool {
	if !strings.EqualFold(ca.name, "cmd") {
		return false
	}
	if strings.EqualFold(ca.command, "") {
		return false
	}
	return true
}
