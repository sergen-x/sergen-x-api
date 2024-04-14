package podman

import (
	"context"
	"log"

	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/containers/podman/v5/pkg/specgen"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/sergen-x/sergen-x-api/pkg/utils"
)

var conn context.Context

type Container struct {
	Name string // UUID corresponding to our database
	Id   string // Containers unique ID
}

type Resouces struct {
	Memory uint8 // Ram in GiB
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
		return false, err
	}
	return true, nil
}

func Create(image string, resources Resouces) (Container, error) {
	// generate a UUID taking the place of a container name
	containerName, err := utils.GenerateUUID(16)
	if err != nil {
		return Container{}, err
	}

	spec := specgen.NewSpecGenerator(image, false)
	ram := utils.GibiBytesToBytes(resources.Memory)
	spec.ResourceLimits = &specs.LinuxResources{
		CPU: &specs.LinuxCPU{},
		Memory: &specs.LinuxMemory{
			Limit: &ram,
		},
	}

	response, err := containers.CreateWithSpec(conn, spec, nil)
	if err != nil {
		return Container{}, err
	}

	return Container{
		Name: containerName,
		Id:   response.ID,
	}, nil
}

func Prune() (uint64, error) {
	var totalPruned uint64
	report, err := containers.Prune(conn, nil)
	if err != nil {
		return totalPruned, err
	}

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

func Update(container Container, newResources Resouces) (bool, error) {
	ram := utils.GibiBytesToBytes(newResources.Memory)
	spec := specgen.NewSpecGenerator("", false)
	spec.ResourceLimits = &specs.LinuxResources{
		CPU: &specs.LinuxCPU{},
		Memory: &specs.LinuxMemory{
			Limit: &ram,
		},
	}

	options := types.ContainerUpdateOptions{
		NameOrID: container.Name,
		Specgen:  spec,
	}

	_, err := containers.Update(conn, &options)
	if err != nil {
		return false, err
	}
	return true, nil
}
