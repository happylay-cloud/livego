package configure

import (
	"fmt"

	"github.com/gwuhaolin/livego/utils/uid"

	"github.com/go-redis/redis/v7"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type RoomKeysType struct {
	redisCli   *redis.Client
	localCache *cache.Cache
}

// 实例化本地缓存对象
var RoomKeys = &RoomKeysType{
	localCache: cache.New(cache.NoExpiration, 0),
}

// 启用本地缓存标志
var saveInLocal = true

func Init() {

	// 判断是否启用redis
	saveInLocal = len(Config.GetString("redis_addr")) == 0
	if saveInLocal {
		return
	}

	// 启用redis缓存
	RoomKeys.redisCli = redis.NewClient(&redis.Options{
		Addr:     Config.GetString("redis_addr"),
		Password: Config.GetString("redis_pwd"),
		DB:       0,
	})

	// 判断redis是否连接成功
	_, err := RoomKeys.redisCli.Ping().Result()
	if err != nil {
		log.Panic("Redis：", err)
	}

	log.Info("Redis成功连接")
}

// SetKey 保存一个信道（channel）的随机密钥（channelKey），set/reset接口使用
func (r *RoomKeysType) SetKey(channel string) (key string, err error) {

	// TODO 此处应该设置缓存时间，暂时不重要，警告，本地缓存本质上是用map保存，存在并发读写异常问题️
	if !saveInLocal {
		for {
			// 获取随机密钥：rfBd56ti2SMtYvSgD5xAV0YU99zampta7Z7S575KLkIZ9PYk
			key = uid.RandStringRunes(48)
			// 判断key是否存在，不存在则保存新值
			if _, err = r.redisCli.Get(key).Result(); err == redis.Nil {
				err = r.redisCli.Set(channel, key, 0).Err()
				if err != nil {
					return
				}
				err = r.redisCli.Set(key, channel, 0).Err()
				return
			} else if err != nil {
				return
			}
		}
	}

	// 查询本地缓存
	for {
		// 获取随机密钥：rfBd56ti2SMtYvSgD5xAV0YU99zampta7Z7S575KLkIZ9PYk
		key = uid.RandStringRunes(48)
		// 查询本地缓存是否存在，不存在则保存新值
		if _, found := r.localCache.Get(key); !found {
			r.localCache.SetDefault(channel, key)
			r.localCache.SetDefault(key, channel)
			break
		}
	}
	return
}

// GetKey 获取密钥
func (r *RoomKeysType) GetKey(channel string) (newKey string, err error) {

	// 从redis缓存中获取
	if !saveInLocal {
		if newKey, err = r.redisCli.Get(channel).Result(); err == redis.Nil {
			newKey, err = r.SetKey(channel)
			log.Debugf("[KEY] 新的信道（channel）[%s]：%s", channel, newKey)
			return
		}

		return
	}

	var key interface{}
	var found bool
	// 从本地缓存中获取
	if key, found = r.localCache.Get(channel); found {
		return key.(string), nil
	}

	// 获取不到，则设置新的密钥
	newKey, err = r.SetKey(channel)
	log.Debugf("[KEY] 新的信道（channel）[%s]：%s", channel, newKey)
	return
}

// GetChannel 获取信道（channel）
func (r *RoomKeysType) GetChannel(key string) (channel string, err error) {

	// 从redis中获取
	if !saveInLocal {
		return r.redisCli.Get(key).Result()
	}

	// 从本地缓存中获取
	chann, found := r.localCache.Get(key)
	if found {
		return chann.(string), nil
	} else {
		return "", fmt.Errorf("%s不存在", key)
	}
}

// DeleteChannel 删除信道（channel）
func (r *RoomKeysType) DeleteChannel(channel string) bool {

	// 删除redis信道（channel）
	if !saveInLocal {
		return r.redisCli.Del(channel).Err() != nil
	}

	// 删除本地缓存信道（channel）
	key, ok := r.localCache.Get(channel)
	if ok {
		r.localCache.Delete(channel)
		r.localCache.Delete(key.(string))
		return true
	}
	return false
}

// DeleteKey 删除密钥
func (r *RoomKeysType) DeleteKey(key string) bool {

	// 删除redis密钥
	if !saveInLocal {
		return r.redisCli.Del(key).Err() != nil
	}

	// 删除本地缓存密钥
	channel, ok := r.localCache.Get(key)
	if ok {
		r.localCache.Delete(channel.(string))
		r.localCache.Delete(key)
		return true
	}
	return false
}
