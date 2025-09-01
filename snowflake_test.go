package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowflake(t *testing.T) {
	sf, err := NewSnowflake(1, 1)
	if err != nil {
		t.Fatal(err)
	}

	id := sf.NextID()
	if id == 0 {
		t.Fatal("ID should not be 0")
	}

	fmt.Println(id)

	parts := ParseID(id)
	if parts["datacenter"] != 1 || parts["node"] != 1 {
		t.Fatal("Datacenter or Node mismatch")
	}
}
