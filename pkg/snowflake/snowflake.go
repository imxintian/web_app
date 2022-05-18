package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

// Snowflake is a distributed unique ID generator.
// It's a combination of a 64-bit timestamp and a 64-bit sequence number.
// The timestamp is the number of milliseconds since the Unix epoch.

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02 15:04:05", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return

}

func GenID() (id int64) {
	id = node.Generate().Int64()
	return
}
