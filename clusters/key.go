package clusters

import (
	"fmt"

	"github.com/GabeCordo/keitt/processor/components/channel"
	"github.com/GabeCordo/keitt/processor/components/cluster"
)

var K = KeyCluster{}

type KeyCluster struct {
}

func (cluster *KeyCluster) ExtractFunc(h cluster.H, m cluster.M, c channel.OneWay) {
	msg := m.GetKey("message")
	c.Push(msg)
}

func (cluster *KeyCluster) TransformFunc(h cluster.H, m cluster.M, in any) (out any, success bool) {
	msg := (in).(string)
	return msg + ", World!", true
}

func (cluster *KeyCluster) LoadFunc(h cluster.H, m cluster.M, in any) {
	msg := (in).(string)
	fmt.Println(msg)
}
