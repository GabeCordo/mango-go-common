package clusters

import (
	"fmt"

	"github.com/GabeCordo/mango/components/channel"
	"github.com/GabeCordo/mango/components/cluster"
	"github.com/GabeCordo/mango/utils"
)

type KeyCluster struct {
	helper utils.Helper
}

func (cluster *KeyCluster) SetHelper(helper utils.Helper) {
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
