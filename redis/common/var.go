package common

// redis链接库
var RFStruct = new(redisFactory).init()

// redis库下标
type RedisDatabase int

//项目redis库下标
const RedisYMZY RedisDatabase = 0
