package serde

import (
	"fmt"
	"time"
)

func init() {
	RegisterTypeWithSerde[time.Time](serializeTime, deserializeTime)
}

func serializeTime(s *Serializer, x *time.Time) error {
	data, err := x.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal time.Time: %w", err)
	}
	Serialize(s, data)
	return nil
}

func deserializeTime(d *Deserializer, x *time.Time) error {
	var b []byte
	DeserializeTo(d, &b)
	return x.UnmarshalBinary(b)
}