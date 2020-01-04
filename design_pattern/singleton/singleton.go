package singleton

import "sync"

var Single1 *Singleton
var Single2 = &Singleton{} //进程启动时初始化,先于Main.main之前
var Single3 *Singleton

type Singleton struct {
	mu  sync.RWMutex
	one sync.Once
}

//协程安全：双重锁机制（真正需要时才进行初始化）
func (this *Singleton) Singleton1() *Singleton {
	if Single1 == nil {
		this.mu.Lock()
		if Single1 == nil {
			Single1 = &Singleton{}
		}
		this.mu.Unlock()
	}
	return Single1
}

//协程安全：进程启动时，提前初始化对象，在需要时返回。
func (this *Singleton) Singleton2() *Singleton {
	return Single2
}

//协程安全：通过go语言的once.do特性，完成实例的创建。（真正需要时才进行初始化）
//once.do的底层实现逻辑：也是以atomic.uint32变量作为标志，加锁的方式对该变量进行读取
func (this *Singleton) Singleton3() *Singleton {
	this.one.Do(func() {
		Single3 = &Singleton{}
	})
	return Single3
}
