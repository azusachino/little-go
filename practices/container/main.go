package main

import (
	"context"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"log"
	"syscall"
	"time"
)

func main() {
	if err := runRedis(); err != nil {
		log.Fatal(err)
	}
}

func runRedis() error {
	// create a new client connected to the default socket path
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}

	defer client.Close()

	// create new context with "ficus" namespace
	ctx := namespaces.WithNamespace(context.Background(), "ficus")

	img, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)

	if err != nil {
		return err
	}

	container, err := client.NewContainer(
		ctx,
		"redis-server",
		containerd.WithImage(img),
		containerd.WithNewSnapshot("redis-server-snapshot", img),
		containerd.WithNewSpec(oci.WithImageConfig(img)))

	if err != nil {
		return err
	}

	defer container.Delete(ctx, containerd.WithSnapshotCleanup)

	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return err
	}

	defer task.Delete(ctx)

	exitStatusCh, err := task.Wait(ctx)
	if err != nil {
		log.Println(err)
	}

	if err := task.Start(ctx); err != nil {
		return err
	}

	time.Sleep(3 * time.Second)

	if err := task.Kill(ctx, syscall.SIGTERM); err != nil {
		return err
	}

	status := <-exitStatusCh
	code, _, err := status.Result()

	if err != nil {
		return err
	}

	log.Printf("redis-server exited with status: %d\n", code)
	return nil
}
