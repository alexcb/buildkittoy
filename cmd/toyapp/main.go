package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/moby/buildkit/client"
	_ "github.com/moby/buildkit/client/connhelper/dockercontainer" // Load "docker-container://" helper.
	"github.com/moby/buildkit/client/llb"
)

func main() {
	fmt.Println("running!")

	goAlpine := llb.Image("docker.io/library/golang:1.13-alpine")

	foo := goAlpine.
		AddEnv("MYENV", "MYVAL").
		Run(llb.Shlex("touch /file")).Root()

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

	go func() {
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
		for {
			status := <-ch
			fmt.Printf("status is %v\n", status)
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

	//eventually it gets pipped to "docker load" (see loadDockerTar)
	_ = pipeR

	fmt.Printf("here i am\n")
}
