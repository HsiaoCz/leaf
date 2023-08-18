package snowflak

import (
	"github/HsiaoCz/leaf/etc"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Init() (err error) {
	app := etc.Conf.App
	var st time.Time
	st, err = time.Parse("2006-01-02", app.StartTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(int64(app.MachineID))
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
