package clusters

import (
	"fmt"

	"github.com/GabeCordo/processor-framework/processor/components/cluster"
)

var K = KeyCluster{}

type KeyCluster struct {
}

func (cluster *KeyCluster) ExtractFunc(h cluster.H, m cluster.M, out cluster.Out) {
	msg := m.GetKey("message")
	out.Push(msg)
}

func (cluster *KeyCluster) TransformFunc(h cluster.H, m cluster.M, in any) (out any, success bool) {
	msg := (in).(string)
	return msg + ", World!", true
}

func (cluster *KeyCluster) LoadFunc(h cluster.H, m cluster.M, in any) {
	msg := (in).(string)
	fmt.Println(msg)
}
