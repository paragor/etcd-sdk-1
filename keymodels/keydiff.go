package keymodels

import (
	"strings"
)

type KeysDiff struct {
	Upserts   map[string]string
	Deletions []string
}

func (diff *KeysDiff) IsEmpty() bool {
	return len(diff.Upserts) == 0 && len(diff.Deletions) == 0
}

func GetKeysDiff(srcKeys map[string]KeyInfo, srcPrefix string, dstKeys map[string]KeyInfo, dstPrefix string) KeysDiff {
	diffs := KeysDiff{
		Upserts:   make(map[string]string),
		Deletions: []string{},
	}

	for key, _ := range dstKeys {
		suffix := strings.TrimPrefix(key, dstPrefix)
		if _, ok := srcKeys[srcPrefix+suffix]; !ok {
			diffs.Deletions = append(diffs.Deletions, suffix)
		}
	}

	for key, srcVal := range srcKeys {
		suffix := strings.TrimPrefix(key, srcPrefix)
		dstVal, ok := dstKeys[dstPrefix+suffix]
		if (!ok) || dstVal.Value != srcVal.Value {
			diffs.Upserts[suffix] = srcVal.Value
		}
	}

	return diffs
}
