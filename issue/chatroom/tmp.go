
type RocketGroup struct {
	ID         bson.ObjectId
	App        string
	Area       string
	Name       string
	CreateTime int64
	UpdateTime int64
	State      RocketStat
	Cycle      int64
	Rockets    []int64
}

type RocketGroupStat string

const (
	RocketStateOnline     RocketGroupStat = "online"
	RocketStateOffline    RocketGroupStat = "offline"
	RocketStateWaitOnline RocketGroupStat = "wait_online"
)

