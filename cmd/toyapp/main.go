package main

import (
	"fmt"
	"os"

	"github.com/moby/buildkit/client"
	"github.com/moby/buildkit/client/llb"
)

func main() {
	fmt.Println("hello")

	goAlpine := llb.Image("docker.io/library/golang:1.13-alpine")

	foo := goAlpine.
		AddEnv("MYENV", "MYVAL").
		Run(llb.Shlex("echo hello > /file")).Root()

	//var out llb.State
	//out = foo.Run(llb.Shlex("ls -l /"))

	dt, err := foo.Marshal(llb.LinuxAmd64)
	if err != nil {
		panic(err)
	}
	llb.WriteTo(dt, os.Stdout)

	//bkClient, err := client.New(context.Todo(), address, opts...)

}
