package docker

import (
	"context"
	"log"

	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/compose/v5/pkg/api"
	"github.com/docker/compose/v5/pkg/compose"
	"github.com/moby/moby/client"
)

func CreateBindVolume(ctx context.Context, cli *client.Client) error {
	_, err := cli.VolumeCreate(ctx, client.VolumeCreateOptions{
		Name:   "namedvol",
		Driver: "local",
		DriverOpts: map[string]string{
			"type":   "none",
			"o":      "bind",
			"device": "/mnt/vol01/test/",
		},
		Labels: map[string]string{
			"orca.managed": "true",
			"orca.cluster": "MyOrca",
		},
	})

	return err
}

type dockerCli struct {
	ctx context.Context
	cli *command.DockerCli
}

func (d *dockerCli) CreateVolume(name string) {
	d.cli.Client().VolumeCreate(d.ctx, client.VolumeCreateOptions{
		Name:       name,
		DriverOpts: map[string]string{},
	})
}

func (d *dockerCli) ComposeUp(dir string) {
	service, err := compose.NewComposeService(d.cli)
	if err != nil {
		log.Fatalf("Failed to create compose service: %v", err)
	}
	project, err := service.LoadProject(d.ctx, api.ProjectLoadOptions{
		WorkingDir: dir,
	})
	if err != nil {
		log.Fatalf("Failed to load project: %v", err)
	}
	err = service.Up(d.ctx, project, api.UpOptions{
		Create: api.CreateOptions{},
		Start:  api.StartOptions{},
	})
	if err != nil {
		log.Fatalf("Failed to start services: %v", err)
	}

	log.Printf("Successfully started project: %s", project.Name)
}

func main() {
	ctx := context.Background()

	dockerCLI, err := command.NewDockerCli()

	if err != nil {
		log.Fatalf("Failed to create docker CLI: %v", err)
	}
	err = dockerCLI.Initialize(&flags.ClientOptions{})
	if err != nil {
		log.Fatalf("Failed to initialize docker CLI: %v", err)
	}

	// Create a new Compose service instance
	service, err := compose.NewComposeService(dockerCLI)

	if err != nil {
		log.Fatalf("Failed to create compose service: %v", err)
	}

	// Load the Compose project from a compose file
	project, err := service.LoadProject(ctx, api.ProjectLoadOptions{
		ConfigPaths: []string{"compose.yaml"},
		ProjectName: "my-app",
	})
	if err != nil {
		log.Fatalf("Failed to load project: %v", err)
	}

	// Start the services defined in the Compose file
	err = service.Up(ctx, project, api.UpOptions{
		Create: api.CreateOptions{},
		Start:  api.StartOptions{},
	})
	if err != nil {
		log.Fatalf("Failed to start services: %v", err)
	}

	log.Printf("Successfully started project: %s", project.Name)
}
