package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/salihkemaloglu/go-ipfs-tests/db"
	"github.com/salihkemaloglu/go-ipfs-tests/repos"
)

func main() {
	// execute ipfs daemon
	go func() {
		cmdex := exec.Command("ipfs", "daemon")
		cmdOutputex := &bytes.Buffer{}
		cmdex.Stdout = cmdOutputex
		errex := cmdex.Run()
		if errex != nil {
			os.Stderr.WriteString(errex.Error())
		}
		fmt.Print(string(cmdOutputex.Bytes()))
	}()

	//add to ipfs
	repo := db.Item{Name: "hello world!"}
	var e repos.BaseRepository = repo
	var items, _ = e.Insert()
	fmt.Println(items)

	// unpin from ipfs(pin/rm)
	sh := shell.NewShell("localhost:5001")
	err := sh.Unpin(items)
	if err != nil {
		fmt.Println(err)
	}
	// execute garbage collection
	cmd := exec.Command("ipfs", "repo", "gc")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err1 := cmd.Run()
	if err1 != nil {
		os.Stderr.WriteString(err1.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))

}
