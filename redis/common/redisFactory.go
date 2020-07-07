package common

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

// redis链接工厂对象
type redisFactory struct {
	connectClient map[RedisDatabase]*redis.Client
	connectMutex  *sync.Mutex
}

// 初始化链接
func (this *redisFactory) init() *redisFactory {
	this.connectClient = make(map[RedisDatabase]*redis.Client)
	this.connectMutex = new(sync.Mutex)
	return this
}

// 链接redis客户端
// @param index 库下标
func (this *redisFactory) Connect(opt *redis.Options) {
	this.connectMutex.Lock()
	defer this.connectMutex.Unlock()

	client := redis.NewClient(opt)
	if ping := client.Ping().Err(); ping != nil {
		panic(errors.New(fmt.Sprintf("redis connection fail %v", ping.Error())))
	}

	this.connectClient[RedisDatabase(opt.DB)] = client
}

// 获取一个redis链接库的客户端
// @param index 库下标文件
func (this *redisFactory) Client(index RedisDatabase) *redis.Client {

	this.connectMutex.Lock()
	defer this.connectMutex.Unlock()

	connect, exists := this.connectClient[index]
	if !exists {
		panic(errors.New(fmt.Sprintf("undinfed redis connection in %v , please set redis.Client int this.SetConnect", index)))
	}
	return connect
}


