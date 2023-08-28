package main

import (
	"github.com/GabeCordo/keitt-common/clusters"
	"github.com/GabeCordo/keitt/processor"
	"github.com/GabeCordo/keitt/processor/components/cluster"
)

func main() {

	pcfg := &processor.Config{
		Name:               "vectors",
		Debug:              true,
		MaxWaitForResponse: 2,
		Core:               "",
		Net: struct {
			Host string "yaml:\"host\""
			Port int    "yaml:\"port\""
		}{Host: "localhost", Port: 5023},
		StandaloneMode: true,
	}

	processor, err := processor.New(pcfg)
	if err != nil {
		panic(err)
	}

	mod := processor.Module("common")

	v := clusters.VectorCluster{}
	ccfg := &cluster.Config{
		Identifier:                  "vec",
		OnLoad:                      cluster.CompleteAndPush,
		OnCrash:                     cluster.DoNothing,
		StartWithNTransformClusters: 1,
		StartWithNLoadClusters:      1,
		ETChannelThreshold:          1,
		ETChannelGrowthFactor:       2,
		TLChannelThreshold:          1,
		TLChannelGrowthFactor:       2,
	}
	mod.AddCluster("vec", cluster.Batch, v, ccfg)

	k := clusters.KeyCluster{}
	kcfg := &cluster.Config{
		Identifier:                  "key",
		OnLoad:                      cluster.CompleteAndPush,
		OnCrash:                     cluster.DoNothing,
		StartWithNTransformClusters: 1,
		StartWithNLoadClusters:      1,
		ETChannelThreshold:          1,
		ETChannelGrowthFactor:       2,
		TLChannelThreshold:          1,
		TLChannelGrowthFactor:       2,
	}
	mod.AddCluster("key", cluster.Batch, &k, kcfg)

	processor.Run()
}
