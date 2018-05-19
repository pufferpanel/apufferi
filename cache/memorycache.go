package cache

import "time"

type MemoryCache struct {
	Cache
	Buffer   []Message
	Capacity int
}

func (c *MemoryCache) Read() (msg []string, lastTime int64) {
	msg, lastTime = c.ReadFrom(0)
	return
}

func (c *MemoryCache) ReadFrom(startTime int64) (msg []string, lastTime int64) {
	result := make([]string, 0)

	var endTime int64 = 0

	for _, v := range c.Buffer {
		if v.time > startTime {
			result = append(result, v.msg)
			endTime = v.time
		}
	}

	if endTime == 0 {
		endTime = time.Now().Unix()
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
