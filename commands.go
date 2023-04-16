package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Command struct {
	command   string
	arguments []string
}

func NewCommand(command, arguments string) *Command {
	return &Command{
		command:   command,
		arguments: strings.Split(arguments, " "),
	}
}

func (c *Command) ExecCommand(command *Command) (string, error) {
	cmd := exec.Command(command.command, command.arguments...)
	execCommand, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error while executing the command : ", err)
		return "", err
	}
	return string(execCommand), nil
}

func SplitVideoToframes(file string) (string, error) {
	arguments := fmt.Sprintf("-i %s -vf fps=1 %s%%04d.png", file, file)
	log.Println("arguments", arguments)
	command := NewCommand("ffmpeg", arguments)
	output, err := command.ExecCommand(command)
	if err != nil {
		return "", err
	}
	return output, nil
}
