package cache

import "time"

type MemoryCache struct {
	Cache
	Buffer   []Message
	Capacity int
}

func (c *MemoryCache) Read() (msg []string, lasttime int64) {
	msg, lasttime = c.ReadFrom(0)
	return
}

func (c *MemoryCache) ReadFrom(startTime int64) (msg []string, lasttime int64) {
	result := make([]string, 0)
	endTime := time.Now().Unix()
	for _, v := range c.Buffer {
		if v.time > startTime {
			result = append(result, v.msg)
		}
	}
	return result, endTime
}

func (c *MemoryCache) Write(b []byte) (n int, err error) {
	if len(c.Buffer) == c.Capacity {
		c.Buffer = c.Buffer[1:]
	}
	c.Buffer = append(c.Buffer, Message{msg: string(b), time: time.Now().Unix()})
	n = len(b)
	return
}
