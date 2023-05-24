package server

import (
	"fmt"
	"sync"
)
var ErrOffsetNotFound = fmt.Errorf("offset not found")

type Record struct {
	Value  []byte `json:"value,omitempty"`
	Offset uint64 `json:"offset,omitempty"`
}

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

// 길이는 int64가 낫지 않나?
// 왜 l이 아니고 c일까?
// mutex는 RWMutex가 낫지 않으려나?
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records)) // 추가할 위치 인덱스 == 오프셋
	c.records = append(c.records, record)
	return record.Offset, nil
}

//
func (c *Log) Read(offset uint64) (Record, error) {
	
	c.mu.Lock()
	defer c.mu.Unlock()

	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}

	return c.records[offset], nil
}
