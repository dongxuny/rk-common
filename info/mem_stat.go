package rk_info

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"runtime"
	"time"
)

type MemStats struct {
	MemAllocMB      uint64 `json:"mem_alloc_byte"`
	SysAllocMB      uint64 `json:"sys_alloc_byte"`
	LastGCTimestamp string `json:"last_gc_timestamp"`
	GCCount         uint32 `json:"gc_count"`
	ForceGCCount    uint32 `json:"force_gc_count"`
}

func MemStatsToStruct() *MemStats {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return &MemStats{
		MemAllocMB:      stats.Alloc,
		SysAllocMB:      stats.Sys,
		LastGCTimestamp: time.Unix(int64(stats.LastGC), 0).Format(time.RFC3339),
		GCCount:         stats.NumGC,
		ForceGCCount:    stats.NumForcedGC,
	}
}

func MemStatsToJSON() string {
	return structToJSON(MemStatsToStruct())
}

func MemStatsToJSONPretty() string {
	return structToJSONPretty(MemStatsToStruct())
}

func MemStatsToBytes() []byte {
	return structToBytes(MemStatsToStruct())
}

func MemStatsToMap() map[string]interface{} {
	return structToMap(MemStatsToStruct())
}

func MemStatsToFields() []zap.Field {
	return structToFields(MemStatsToStruct())
}

func structToJSON(src interface{}) string {
	return string(structToBytes(src))
}

func structToJSONPretty(src interface{}) string {
	mid := structToBytes(src)
	dest := &bytes.Buffer{}
	json.Indent(dest, mid, "", "  ")

	return dest.String()
}

func structToBytes(src interface{}) []byte {
	bytes, _ := json.Marshal(src)
	return bytes
}

func structToMap(src interface{}) map[string]interface{} {
	bytes := structToBytes(src)
	res := make(map[string]interface{})

	json.Unmarshal(bytes, res)

	return res
}

func structToFields(src interface{}) []zap.Field {
	mid := structToMap(src)
	fields := make([]zap.Field, 0)

	for k, v := range mid {
		fields = append(fields, zap.Any(k, v))
	}

	return fields
}
