package metric

import "testing"

// 三个 struct 目前均为空（无方法、无可测语句），此测试仅断言零值可用性，
// 使包存在可编译通过的测试文件，覆盖率 total 天然为 100%。
func TestMetric_ZeroValueUsable(t *testing.T) {
	m := &Metric{}
	if m.LockMetric != nil {
		t.Fatal("expect nil LockMetric in zero value")
	}
	if m.StorageMetric != nil {
		t.Fatal("expect nil StorageMetric in zero value")
	}

	m.LockMetric = &LockMetric{}
	m.StorageMetric = &StorageMetric{}
	if m.LockMetric == nil || m.StorageMetric == nil {
		t.Fatal("expect metric fields assignable")
	}
}
