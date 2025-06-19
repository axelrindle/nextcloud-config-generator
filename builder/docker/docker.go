package docker

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"embed"
	"io"
	"io/fs"

	"github.com/axelrindle/nc-cfg-gen/nextcloud"
	t "github.com/axelrindle/nc-cfg-gen/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

//go:embed rootfs
var rootfs embed.FS

type Docker struct {
	client    *client.Client
	ctx       context.Context
	container string

	result *nextcloud.ConfigSecrets
}

type ErrorFunc func() *t.Error

func makeBuildContext() (io.Reader, error) {
	b := bytes.Buffer{}
	w := bufio.NewWriter(&b)
	t := tar.NewWriter(w)

	sub, err := fs.Sub(rootfs, "rootfs")
	if err != nil {
		return nil, err
	}

	if err = t.AddFS(sub); err != nil {
		return nil, err
	}

	if err = w.Flush(); err != nil {
		return nil, err
	}

	if err = t.Close(); err != nil {
		return nil, err
	}

	return bufio.NewReader(&b), nil
}

func makeDockerClient() (*client.Client, *t.Error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, &t.Error{Err: err, Context: "Create Docker API Client"}
	}

	return client, nil
}

func (b *Docker) buildImage() *t.Error {
	context, err := makeBuildContext()
	if err != nil {
		return &t.Error{Err: err, Context: "Build Context"}
	}

	resp, err := b.client.ImageBuild(b.ctx, context, types.ImageBuildOptions{
		Tags:    []string{"nc-cfg-gen"},
		NoCache: false,
	})
	if err != nil {
		return &t.Error{Err: err, Context: "Build Image"}
	}

	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return &t.Error{Err: err, Context: "Build Image"}
	}

	return nil
}

func (b *Docker) runGenerator() *t.Error {
	resp, err := b.client.ContainerCreate(b.ctx,
		&container.Config{
			Image: "nc-cfg-gen",
			Cmd:   []string{"/custom-entrypoint.sh"},
			Tty:   true,
		},
		nil,
		nil,
		nil,
		"nc-cfg-gen",
	)
	if err != nil {
		return &t.Error{Err: err, Context: "Create Container"}
	}
	b.container = resp.ID

	if err := b.client.ContainerStart(b.ctx, resp.ID, container.StartOptions{}); err != nil {
		return &t.Error{Err: err, Context: "Start Container"}
	}

	rc, err := b.client.ContainerLogs(b.ctx, resp.ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: false,
		Follow:     true,
	})
	if err != nil {
		return &t.Error{Err: err, Context: "Stream Container Logs"}
	}
	defer rc.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, rc)
	if err != nil {
		return &t.Error{Err: err, Context: "Read Container Logs"}
	}

	config := &nextcloud.ConfigSecrets{}
	err = config.Parse(buf.Bytes())
	if err != nil {
		return &t.Error{Err: err, Context: "Parse Nextcloud Config"}
	}

	b.result = config

	return nil
}

func (b *Docker) deleteContainer() *t.Error {
	err := b.client.ContainerRemove(b.ctx, b.container, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		return &t.Error{Err: err, Context: "Remove Container"}
	}

	return nil
}

func GenerateConfig() (*nextcloud.ConfigSecrets, *t.Error) {
	client, err := makeDockerClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	builder := &Docker{
		client: client,
		ctx:    context.Background(),
	}

	tasks := []ErrorFunc{builder.buildImage, builder.runGenerator, builder.deleteContainer}

	for _, task := range tasks {
		err = task()
		if err != nil {
			builder.deleteContainer()
			return nil, err
		}
	}

	return builder.result, nil
}
