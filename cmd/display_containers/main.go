package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"os"
)

func main() {

	os.Setenv("DOCKER_HOST", "unix:///home/eunice99x/.docker/desktop/docker.sock")
	os.Setenv("DOCKER_CONTEXT", "desktop-linux")
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	// List all containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}

	// Separate active and inactive containers
	var activeContainers []types.Container
	var inactiveContainers []types.Container

	for _, container := range containers {
		if container.State == "running" {
			activeContainers = append(activeContainers, container)
		} else {
			inactiveContainers = append(inactiveContainers, container)
		}
	}

	// Display metadata for active containers
	fmt.Println("Active Containers:")
	for _, container := range activeContainers {
		fmt.Printf("ID: %s\n", container.ID)
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Names: %v\n", container.Names)
		fmt.Printf("State: %s\n", container.State)
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Println("-------")
	}

	// Display metadata for inactive containers
	fmt.Println("Inactive Containers:")
	for _, container := range inactiveContainers {
		fmt.Printf("ID: %s\n", container.ID)
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Names: %v\n", container.Names)
		fmt.Printf("State: %s\n", container.State)
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Println("-------")
	}
}
