package main

import (
	"os"

	"github.com/rancher/system-tools/logs"
	"github.com/rancher/system-tools/remove"
	"github.com/rancher/system-tools/stats"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	DefaultCattleNamespace = "cattle-system"
)

var VERSION = "dev"

func main() {
	commonFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "kubeconfig,c",
			EnvVar: "KUBECONFIG",
			Usage:  "kubeconfig absolute path",
		},
		cli.StringFlag{
			Name:  "namespace,n",
			Value: DefaultCattleNamespace,
			Usage: "rancher 2.x deployment namespace. default is `cattle-system`",
		},
	}
	app := cli.NewApp()
	app.Name = "system-tools"
	app.Version = VERSION
	app.Usage = "Rancher 2.x operations tool kit"
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "remove",
			Usage:  "safely remove rancher 2.x management plane",
			Action: remove.DoRemoveRancher,
			Flags:  append(commonFlags, remove.ForceFlag),
		},

		cli.Command{
			Name:   "logs",
			Usage:  "inspect logs for rancher 2.x managed clusters",
			Action: logs.DoLogs,
			Flags:  logs.LogFlags,
		},
		cli.Command{
			Name:   "stats",
			Usage:  "show live system stats from cluster nodes",
			Action: stats.DoStats,
			Flags:  stats.StatsFlags,
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
