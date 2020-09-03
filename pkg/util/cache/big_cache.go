package cache

//var config = bigcache.Config{
//	Shards:             8,                // shard map 分片大小 2的整数倍
//	LifeWindow:         10 * time.Minute, // 存活时间，超过之后视为无效条目，但不能删除
//	CleanWindow:        5 * time.Minute,  // 删除过期对象之间的间隔，bigcache分辨率为1秒
//	MaxEntriesInWindow: 1000 * 10 * 60,   // 初始窗口大小 仅用于初始内存分配
//	MaxEntrySize:       500,              // 单个对象最大大小 字节单位，仅用于初始内存分配
//	Verbose:            true,             // 是否打印额外分配内存信息
//	HardMaxCacheSize:   8192,             // 总缓存上限，超过将覆盖旧对象 单位MB 0表示不限制
//	OnRemove:           nil,              // 删除旧对象触发回调
//	OnRemoveWithReason: nil,              // 删除旧对象触发回调-附带原因
//}
//
//func init() {
//	cache, err := bigcache.NewBigCache(config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	_ = cache.Set("a", []byte("aaa"))
//	_, _ = cache.Get("a")
//}
