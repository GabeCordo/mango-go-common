package clusters

import (
	"fmt"

	"github.com/GabeCordo/keitt/processor/components/channel"
	"github.com/GabeCordo/keitt/processor/components/cluster"
	"github.com/GabeCordo/keitt/processor/components/helper"
)

type KeyCluster struct {
	helper helper.Helper
}

func (cluster *KeyCluster) SetHelper(helper helper.Helper) {
	cluster.helper = helper
}

func (cluster *KeyCluster) ExtractFunc(m cluster.M, c channel.OneWay) {
	msg := m.GetKey("message")
	c.Push(msg)
}

func (cluster *KeyCluster) TransformFunc(m cluster.M, in any) (out any, success bool) {
	msg := (in).(string)
	return msg + ", World!", true
}

func (cluster *KeyCluster) LoadFunc(m cluster.M, in any) {
	msg := (in).(string)
	fmt.Println(msg)
}
