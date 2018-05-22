package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	agent "github.com/go-ignite/ignite-agent"
	"github.com/go-ignite/ignite-agent/service"
	"github.com/olekukonko/tablewriter"

	"github.com/briandowns/spinner"
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
	if c.Bool("verbose") {
		output = os.Stdout
	}
	services := service.GetServices()
	wg := new(sync.WaitGroup)
	wg.Add(len(services))

	for _, s := range services {
		func(s *service.Service) {
			defer wg.Done()
			sp := spinner.New(spinner.CharSets[0], 100*time.Millisecond)
			sp.Color("green")
			sp.Start()
			fmt.Printf("init service %s", color.MagentaString(s.Name))

			go func() {
				err := agent.BuildImage(s.Dockerfile, output)
				sp.Stop()
				if err != nil {
					fmt.Printf("\t%s", IconBad)
				} else {
					fmt.Printf("\t%s", IconGood)
				}
				fmt.Println()
			}()
			time.Sleep(1 * time.Second)
		}(s)
	}
	wg.Wait()
	return nil
}

func list(c *cli.Context) error {
	services := service.GetServices()
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
	table.SetHeader([]string{"Name", "Image", "Version", "Latest"})
	showCaption := false
	for _, service := range services {
		serviceImage := ""
		localVersion := IconBad
		latestVersion := IconGood
		if service.ImageExist {
			serviceImage = service.Image
			localVersion = color.GreenString(service.ImageVersion)
			if service.Version != service.ImageVersion {
				showCaption = true
				latestVersion = color.RedString(service.Version)
			}
		} else {
			showCaption = true
			latestVersion = color.RedString(service.Version)
		}
		table.Append([]string{color.MagentaString(service.Name), serviceImage, localVersion, latestVersion})
	}
	table.Render()
	if showCaption {
		fmt.Printf("\n%s use `%s` command to init or upgrade services.\n", IconInfo, color.GreenString("agent-cli init"))
	}
	return nil
}
