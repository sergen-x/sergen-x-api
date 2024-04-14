package podman

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/specgen"
	"github.com/sergen-x/sergen-x-api/pkg/utils"
)

var conn context.Context

type Container struct {
	Name string // UUID corresponding to our database
	Id   string // Containers unique ID
}

func init() {
	var err error
	conn, err = bindings.NewConnection(context.Background(), "unix:///run/podman/podman.sock")
	if err != nil {
		log.Fatalf("Failed to bind to podman: %v", err)
	}
}

func Start(containerID string) (bool, error) {
	if err := containers.Start(conn, containerID, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return true, nil
}

func Create(image string) (Container, error) {
	// generate a UUID taking the place of a container name
	containerName, err := utils.GenerateUUID(16)
	if err != nil {
		return Container{}, err
	}

	s := specgen.NewSpecGenerator(image, false)
	s.Name = containerName
	response, err := containers.CreateWithSpec(conn, s, nil)
	if err != nil {
		return Container{}, err
	}

	return Container{
		Name: containerName,
		Id:   response.ID,
	}, nil
}

func Prune() (uint64, error) {
	report, err := containers.Prune(conn, nil)
	if err != nil {
		return 0, err
	}
	var totalPruned uint64
	for _, container := range report {
		totalPruned += container.Size
	}

	return totalPruned, nil
}

func Stop(container Container) (bool, error) {
	if err := containers.Stop(conn, container.Name, nil); err != nil {
		return false, err
	}
	return true, nil
}