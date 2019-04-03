package cron

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/golog"
	"hoper/client/controller/credis"
	"hoper/initialize"
	"hoper/model"
	"strconv"
	"sync"
	"time"
)

//这样设计的话，锁的设计就多余了
var StartTime = struct {
	Time int64
	L    sync.RWMutex
}{Time: time.Now().Unix()}

func UserRedisToDB() error {
	conn := initialize.RedisPool.Get()
	defer conn.Close()

	StartTime.L.Lock()
	StartTime.Time = time.Now().Unix() - 3600
	conn.Send("SELECT", credis.CronIndex)
	ids, err := redis.Int64s(conn.Do("ZRANGEBYSCORE", model.LoginUser+"ActiveTime", "-inf", StartTime.Time))
	if err != nil {
		return err
	}
	conn.Do("ZREMRANGEBYSCORE", model.LoginUser+"ActiveTime", "-inf", StartTime.Time)
	StartTime.L.Unlock()

	for id := range ids {
		conn.Send("SELECT")
		err := initialize.DB.Exec("INSERT INTO " + approve.Kind + "_approve_user VALUES (" +
			strconv.FormatUint(approve.RefID, 10) + "," + strconv.FormatUint(userID, 10) + ")").Error
	}

	return nil
}
