package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"os/exec"
	"strings"
)

type CmdAction struct {
	name    string
	command string

	prof *profile.Profile
}

func NewCmdAction(name string, command string, prof *profile.Profile) *CmdAction {
	return &CmdAction{name: name, command: command, prof: prof}
}

func (ca *CmdAction) Name() string {
	return ca.name
}

func (ca *CmdAction) Do(_ *agouti.Page) error {
	if ca.IsActual() {
		if cmd, err := parseVariables(ca.command, ca.prof); err != nil {
			return err
		} else {
			return exec.Command(cmd).Run()
		}
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
