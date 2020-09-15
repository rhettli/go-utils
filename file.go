package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		panic("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}

func CopyDir(src string, dest string) (bool, error) {
	src = FormatPath(src)
	dest = FormatPath(dest)
	log.Println(src)
	log.Println(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	default:
		panic("not support os")
	}

	outPut, e := cmd.Output()
	if e != nil {
		return false, e
	}
	fmt.Println(outPut)
	return true, nil
}
