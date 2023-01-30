package sonyflake

import (
	"github.com/prometheus/common/log"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func NextID() (uint64, error) {
	id, err := sf.NextID()
	if err != nil {
		log.Errorf("sonyflake get next id err : %v", err)
		return 0, err
	}
	return id, nil
}
