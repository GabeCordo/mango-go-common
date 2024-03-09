## processor-framework template

This template is an example of how the [processor framework](https://github.com/GabeCordo/processor-framework) 
can be extended to include your own custom etl functions. 

How the processor communicates with the the cluster.tools core or what additional
functionality it supports can be found [here](https://cluster.tools/).

## creating a new processor
The processor is the entrypoint of the framework that allows you to specify
runtime behaviour and hook custom functions that are callable by the cluster.tools
core.
```go
config := &processor.Config{
	Name: "test",
	Debug: "true",
}

processor, err := processor.New(config)
if err != nil {
	panic(err)
}

// a blocking process that spins up the HTTP endpoint
// and accompanying threads
proecessor.Run()
```

## creating a module
The module encapsulates functions that share common functionality. A processor can support zero to many modules that will be
registered to the cluster.tools core. When a user attempts to call a function inside a module, the cluster.tools core will
look for processors that include support it. 

```go
newModule := processor.Module("common")
```

## hooking a cluster to a module
A cluster is a collection of functions that implement the ETL pattern. Hooking a cluster to a module
signifies that it belongs to a logical separation the module supports, and is callable from the cluster.tools
core under that logical module.

```go
defaultConfig := cluster.GenericConfig("test")

// do not worry about clusters.Batch for now, simply know that this means the function can be invoked on demand
if _, err := mod.AddCluster("test", string(cluster.Batch), &clusterImplementation, defaultConfig); err != nil {
	panic(err)
}
```

### Disclosure

This repository is not related to the contributing members (of the repository) to the organizations they currently belong, the work they have, currently, or will perform at such organizations. All work completed within this repository pre-dates these organizations. All work completed withon this repository shall not be through company resources. Where "company resources" includes but is not limited to working hours, intellectual property, and electronic devices.
