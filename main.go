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
	// add to ipfs
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
	if err != nil {
		os.Stderr.WriteString(err1.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))

}
