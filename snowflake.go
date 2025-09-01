package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	epochMilli     int64 = 1612224000000 // 2021-02-02 00:00:00 UTC
	datacenterBits uint8 = 5
	nodeBits       uint8 = 5
	sequenceBits   uint8 = 12

	maxDatacenter int64 = -1 ^ (-1 << datacenterBits)
	maxNode       int64 = -1 ^ (-1 << nodeBits)
	maxSequence   int64 = -1 ^ (-1 << sequenceBits)

	nodeShift       = sequenceBits
	datacenterShift = sequenceBits + nodeBits
	timestampShift  = sequenceBits + nodeBits + datacenterBits
)

// Snowflake struct
type Snowflake struct {
	mu          sync.Mutex
	datacenter  int64
	node        int64
	sequence    int64
	lastStampMs int64
}

// NewSnowflake creates a new instance
func NewSnowflake(datacenter, node int64) (*Snowflake, error) {
	if datacenter < 0 || datacenter > maxDatacenter {
		return nil, errors.New("invalid datacenter id")
	}
	if node < 0 || node > maxNode {
		return nil, errors.New("invalid node id")
	}

	return &Snowflake{
		datacenter: datacenter,
		node:       node,
		sequence:   0,
	}, nil
}

// timestampMs returns current timestamp in ms
func (s *Snowflake) timestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// NextID generates next unique ID
func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	ts := s.timestampMs()
	if ts == s.lastStampMs {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// wait next millisecond
			for ts <= s.lastStampMs {
				ts = s.timestampMs()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastStampMs = ts

	id := ((ts - epochMilli) << timestampShift) |
		(s.datacenter << datacenterShift) |
		(s.node << nodeShift) |
		s.sequence

	return id
}

// ParseID parses a Snowflake ID into its components
func ParseID(id int64) map[string]int64 {
	sequence := id & maxSequence
	node := (id >> nodeShift) & maxNode
	datacenter := (id >> datacenterShift) & maxDatacenter
	timestamp := (id >> timestampShift) + epochMilli

	return map[string]int64{
		"timestamp":  timestamp,
		"datacenter": datacenter,
		"node":       node,
		"sequence":   sequence,
	}
}
