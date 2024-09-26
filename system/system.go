package system

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// initTime is the time when the application was initialized.
var initTime = time.Now()

type SysStatus struct {
	RunPath    string // 运行路径
	SystemName string // 运行环境
	SystemArch string // 系统架构

	Uptime       string // 服务运行时间
	NumGoroutine int    // 当前 Goroutines 数量

	// General statistics.
	MemAllocated string // 当前内存使用量
	MemTotal     string // 所有被分配的内存
	MemSys       string // 内存占用量
	Lookups      uint64 // 指针查找次数
	MemMallocs   uint64 // 内存分配次数
	MemFrees     uint64 // 内存释放次数

	// Main allocation heap statistics.
	HeapAlloc    string // 当前 Heap 内存使用量
	HeapSys      string // Heap 内存占用量
	HeapIdle     string // Heap 内存空闲量
	HeapInuse    string // 正在使用的 Heap 内存
	HeapReleased string // 释放的 Heap 内存
	HeapObjects  uint64 // Heap 对象数量

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string // 启动 Stack 使用量
	StackSys    string // 被分配的 Stack 内存
	MSpanInuse  string // MSpan 结构内存使用量
	MSpanSys    string // 被分配的 MSpan 结构内存
	MCacheInuse string // MCache 结构内存使用量
	MCacheSys   string // 被分配的 MCache 结构内存
	BuckHashSys string // 被分配的剖析哈希表内存
	GCSys       string // 被分配的 GC 元数据内存
	OtherSys    string // 其它被分配的系统内存

	// Garbage collector statistics.
	NextGC       string // 下次 GC 内存回收量
	LastGC       string // 距离上次 GC 时间
	PauseTotalNs string // GC 暂停时间总量
	PauseNs      string // 上次 GC 暂停时间
	NumGC        uint32 // GC 执行次数
}

var sysStatus = SysStatus{}

func init() {
	runPath, _ := os.Getwd()
	sysStatus.RunPath = runPath

	sysStatus.SystemName = runtime.GOOS
	sysStatus.SystemArch = runtime.GOARCH
}

// GetSystemStatus 获取系统状态
func GetSystemStatus() SysStatus {
	sysStatus.Uptime = TimeSincePro(initTime)

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	sysStatus.NumGoroutine = runtime.NumGoroutine()

	sysStatus.MemAllocated = FileSize(int64(m.Alloc))
	sysStatus.MemTotal = FileSize(int64(m.TotalAlloc))
	sysStatus.MemSys = FileSize(int64(m.Sys))
	sysStatus.Lookups = m.Lookups
	sysStatus.MemMallocs = m.Mallocs
	sysStatus.MemFrees = m.Frees

	sysStatus.HeapAlloc = FileSize(int64(m.HeapAlloc))
	sysStatus.HeapSys = FileSize(int64(m.HeapSys))
	sysStatus.HeapIdle = FileSize(int64(m.HeapIdle))
	sysStatus.HeapInuse = FileSize(int64(m.HeapInuse))
	sysStatus.HeapReleased = FileSize(int64(m.HeapReleased))
	sysStatus.HeapObjects = m.HeapObjects

	sysStatus.StackInuse = FileSize(int64(m.StackInuse))
	sysStatus.StackSys = FileSize(int64(m.StackSys))
	sysStatus.MSpanInuse = FileSize(int64(m.MSpanInuse))
	sysStatus.MSpanSys = FileSize(int64(m.MSpanSys))
	sysStatus.MCacheInuse = FileSize(int64(m.MCacheInuse))
	sysStatus.MCacheSys = FileSize(int64(m.MCacheSys))
	sysStatus.BuckHashSys = FileSize(int64(m.BuckHashSys))
	sysStatus.GCSys = FileSize(int64(m.GCSys))
	sysStatus.OtherSys = FileSize(int64(m.OtherSys))

	sysStatus.NextGC = FileSize(int64(m.NextGC))
	sysStatus.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	sysStatus.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	sysStatus.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	sysStatus.NumGC = m.NumGC
	return sysStatus
}
