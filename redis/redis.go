package redis

import (
	"DataStructureGolang/redis/common"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func ConnectRedis(address string, port int, password string) {

	if address == "" {
		panic(errors.New("指定redis连接地址"))
	}

	if password == "" {
		panic(errors.New("指定redis连接密码"))
	}

	if port < 1 {
		panic(errors.New("请指定redis连接端口"))
	}

	common.RFStruct.Connect(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", address, port),
		Password: password,
		DB: int(common.RedisYMZY),
		PoolSize: 500,
		MaxConnAge: time.Minute,
	})
}
