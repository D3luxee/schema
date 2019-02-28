package schema

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/raintank/schema"
)

//Format will be as below:
// $Prefix.$hostname.$metric
//

type MetricData struct {
	schema.MetricData
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

//Get the lookup id used by cyclone / eye
func (m *MetricData) LookupID() string {
	h := sha256.New()
	h.Write([]byte(m.Hostname()))
	h.Write([]byte(m.MetricName()))
	return hex.EncodeToString(h.Sum(nil))
}

//Get the name of the Metric without the prefix and hostname
func (m *MetricData) MetricName() string {
	name := strings.SplitAfterN(m.Name, ".", 3)
	if len(name) == 3 {
		return name[2]
	}
	return ""
}

//Get the Hostname associated with this metric
func (m *MetricData) Hostname() string {
	name := strings.Split(m.Name, ".")
	if len(name) > 1 {
		return name[1]
	}
	return ""
}
