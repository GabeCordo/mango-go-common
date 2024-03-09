package main

import (
	"os"

	"github.com/GabeCordo/processor-framework/processor"
	"github.com/GabeCordo/processor-framework/processor/components/cluster"
	"github.com/GabeCordo/processor-template/clusters"
)

func main() {

	if len(os.Args) != 2 {
		panic("you need to pass a yaml config path as the first parameter")
	}

	cfg := &processor.Config{}
	if err := processor.ConfigFromYAML(cfg, os.Args[1]); err != nil {
		panic(err)
	}

	p, err := processor.New(cfg)
	if err != nil {
		panic(err)
	}

	mod := p.Module("common")
	mod.Version = 1.0

	vecRunConfig := cluster.GenericConfig("vec")
	if _, err := mod.AddCluster("vec", string(cluster.Batch), &clusters.V, vecRunConfig); err != nil {
		panic(err)
	}

	keyRunConfig := cluster.GenericConfig("key")
	if _, err := mod.AddCluster("key", string(cluster.Batch), &clusters.K, keyRunConfig); err != nil {
		panic(err)
	}

	p.Run()
}
