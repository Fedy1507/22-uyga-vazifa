package git

import "os/exec"

func GetUserName() (string, error) {
	cmd := exec.Command("git", "config", "user.email")
	str, err = cmd.Output()
	if err != nil {
		return nil, err
	}
	return string(str),err
}
