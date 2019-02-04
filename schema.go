package schema

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/raintank/schema"
)

type MetricData struct {
	*schema.MetricData
}

func (m *MetricData) GetTagMap() map[string]string {
	tags := make(map[string]string)
	for _, v := range m.Tags {
		kv := strings.Split(v, "=")
		if len(kv) == 2 {
			tags[kv[0]] = kv[1]
		}
	}
	return tags
}

func (m *MetricData) GetTag(key string) string {
	for _, v := range m.Tags {
		if strings.HasPrefix(v, key) {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				return kv[1]
			}
		}

	}
	return ""
}

func (m *MetricData) AddTag(key, value string) {
	m.Tags = append(m.Tags, key+"="+value)
}

func (m *MetricData) LookupID() string {
	h := sha256.New()
	h.Write([]byte(m.GetTag("hostname")))
	h.Write([]byte(m.Name))
	return hex.EncodeToString(h.Sum(nil))
}
