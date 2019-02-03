package schema

import (
	"strings"

	"github.com/raintank/schema"
)

type MetricData struct {
	*schema.MetricData
}

func (m *MetricData) GetTags() map[string]string {
	tags := make(map[string]string)
	for _, v := range m.Tags {
		kv := strings.Split(v, "=")
		if len(kv) == 2 {
			tags[kv[0]] = kv[1]
		}
	}
	return tags
}
