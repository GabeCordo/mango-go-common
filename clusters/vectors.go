package clusters

import (
	"fmt"

	"github.com/GabeCordo/processor-framework/processor/components/cluster"
)

var V = VectorCluster{}

type Vec2D struct {
	x int
	y int
}

type VectorCluster struct {
}

func (cluster VectorCluster) ExtractFunc(h cluster.H, m cluster.M, out cluster.Out) {

	v := Vec2D{1, 5} // simulate pulling data from a source
	for i := 0; i < 100; i++ {
		out.Push(v) // send data to the TransformFunc
	}
}

func (cluster VectorCluster) TransformFunc(h cluster.H, m cluster.M, in any) (out any, success bool) {

	v := (in).(Vec2D)

	v.x *= 5
	v.y += 5

	return v, true
}

func (cluster VectorCluster) LoadFunc(h cluster.H, m cluster.M, in any) {

	v := (in).(Vec2D)
	fmt.Printf("Vec(x: %d, y: %d)\n", v.x, v.y)
}
