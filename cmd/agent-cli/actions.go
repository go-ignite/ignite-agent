package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/docker/docker/api/types"
	"github.com/go-ignite/ignite-agent"
	"github.com/olekukonko/tablewriter"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var (
	IconGood = color.GreenString("✔")
	IconBad  = color.RedString("✗")
	IconInfo = color.GreenString("▸")
)

func initImages(c *cli.Context) error {
	output := ioutil.Discard
	showVerbose := c.Bool("verbose")
	if showVerbose {
		output = os.Stdout
	}

	services := GetServices()
	wg := new(sync.WaitGroup)

	for _, service := range services {
		serviceName := color.MagentaString(service.Name)
		reader, err := agent.PullImage(service.Image)
		if err != nil {
			fmt.Printf("%s init service %s error: %v\n", IconBad, serviceName, err)
			continue
		}
		fmt.Printf("%s init service %s\n", IconGood, serviceName)

		wg.Add(1)
		go func() {
			defer wg.Done()
			io.Copy(output, reader)
		}()
	}

	s := spinner.New(spinner.CharSets[0], 100*time.Millisecond)
	s.Suffix = " download..."
	if !showVerbose {
		s.Start()
	}
	wg.Wait()
	if !showVerbose {
		s.Stop()
	}
	return nil
}

func list(c *cli.Context) error {
	services := GetServices()
	images, err := agent.ListImages()
	if err != nil {
		fmt.Printf("%s List images error: %v\n", IconBad, err)
		os.Exit(1)
	}
	imageMap := map[string]types.ImageSummary{}
	for _, image := range images {
		for _, name := range image.RepoTags {
			imageMap[name] = image
		}
	}
	for _, service := range services {
		if image, ok := imageMap[service.Image]; ok {
			service.ImageExist = true
			service.ImageVersion = image.Labels["version"]
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader([]string{"Name", "Image", "Version"})
	showCaption := false
	for _, service := range services {
		serviceImage := ""
		localVersion := IconBad
		if service.ImageExist {
			serviceImage = service.Image
			localVersion = color.GreenString(service.ImageVersion)
		} else {
			showCaption = true
		}
		table.Append([]string{color.MagentaString(service.Name), serviceImage, localVersion})
	}
	table.Render()
	if showCaption {
		fmt.Printf("\n%s use `%s` command to init services.\n", IconInfo, color.GreenString("agent-cli init"))
	}
	return nil
}
