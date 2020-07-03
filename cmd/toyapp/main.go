package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/moby/buildkit/client"
	_ "github.com/moby/buildkit/client/connhelper/dockercontainer" // Load "docker-container://" helper.
	"github.com/moby/buildkit/client/llb"
)

func prompt(question string) string {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	fmt.Println("running!")

	goAlpine := llb.Image("docker.io/library/golang:1.13-alpine")

	data := prompt("enter a string: ")

	foo := goAlpine.
		AddEnv("MYENV", "MYVAL").
		File(llb.Mkfile("/file", 0, []byte(data))).
		Run(llb.Shlex("apk add --no-cache git"))

	dt, err := foo.Marshal(llb.LinuxAmd64)
	if err != nil {
		panic(err)
	}

	ContainerName := "earthly-buildkitd"
	address := fmt.Sprintf("docker-container://%s", ContainerName)
	//address = "172.17.0.2"

	bkClient, err := client.New(context.TODO(), address)
	if err != nil {
		panic(err)
	}
	defer bkClient.Close()

	pipeR, pipeW := io.Pipe()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		cmd := exec.CommandContext(context.TODO(), "docker", "load")
		cmd.Stdin = pipeR
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("done loading into docker\n")
	}()

	ch := make(chan *client.SolveStatus)
	go func() {
		defer wg.Done()
		for {
			status := <-ch
			if status == nil {
				break
			}
			fmt.Printf("status is %v\n", status)
			for _, x := range status.Vertexes {
				fmt.Printf(" vertex: %v\n", x)
			}
			for _, x := range status.Statuses {
				fmt.Printf(" status: %v\n", x)
			}
			for _, x := range status.Logs {
				fmt.Printf(" log: %v\n", x)
			}
		}
	}()

	dockerTag := "alextest:latest"

	solveOpt := client.SolveOpt{
		Exports: []client.ExportEntry{
			{
				Type: client.ExporterDocker,
				Attrs: map[string]string{
					"name":                  dockerTag,
					"containerimage.config": "{}",
				},
				Output: func(_ map[string]string) (io.WriteCloser, error) {
					return pipeW, nil
				},
			},
		},
		//Session:             s.attachables,
		//AllowedEntitlements: s.enttlmnts,
		//LocalDirs:           localDirs,
	}
	_, err = bkClient.Solve(context.TODO(), dt, solveOpt, ch)
	if err != nil {
		panic(err)
	}

	wg.Wait()
	fmt.Printf("finished!\n")
}
