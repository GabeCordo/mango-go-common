package main

import (
	"fmt"
	"os"

	"github.com/GabeCordo/keitt-common/clusters"
	"github.com/GabeCordo/keitt/processor"
	"github.com/GabeCordo/keitt/processor/components/cluster"
)

func main() {

	if len(os.Args) != 2 {
		panic("need to pass a yaml config path as the first parameter")
	}

	cfg := &processor.Config{}
	processor.ConfigFromYAML(cfg, os.Args[1])

	fmt.Println(cfg)

	processor, err := processor.New(cfg)
	if err != nil {
		panic(err)
	}

	mod := processor.Module("common")
	mod.Version = 1.0

	vcfg := cluster.GenericConfig("vec")
	mod.AddCluster("vec", cluster.Batch, &clusters.V, vcfg)

	kcfg := cluster.GenericConfig("key")
	mod.AddCluster("key", cluster.Batch, &clusters.K, kcfg)

	processor.Run()
}
