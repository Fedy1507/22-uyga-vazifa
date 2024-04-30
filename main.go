package main

import (
	"NT_Homeworks/22.Working_with_git_and_remote_repositories/22-uyga-vazifa/git"
	"os"
	"os/exec"
)

const PATH = "/home/projects/go/src/NT_Homeworks/22.Working_with_git_and_remote_repositories/22-uyga-vazifa/file.txt"

func PrintToFileDAta(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	file.Write([]byte(data))

	cmd := exec.Command("git", "add", ".")
	cmd.Output()
	cmd := exec.Command("git", "commit", "-m", "appendUsers")
	cmd.Output()
	return nil
}

func main() {
	strData, err := git.GetUserName()
	if err != nil {
		panic(err)
	}
	PrintToFileDAta(PATH, strData)

	strDataemail, err := git.GetUserEmail()
	if err != nil {
		panic(err)
	}
	PrintToFileDAta(PATH, strDataemail)
}
