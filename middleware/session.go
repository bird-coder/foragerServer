package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SessionMgrType string

const (
	SessionCookieName                 = "GOSESSID"
	SessionContextName                = "session"
	File               SessionMgrType = "file"
	Memcached          SessionMgrType = "memcached"
	Redis              SessionMgrType = "redis"
)

type Session interface {
	//获取session对象的id
	ID() string
	//加载cache数据到session data
	Load() error
	//获取key对应的value
	Get(string) (interface{}, error)
	//设置key对应的value
	Set(string, interface{})
	//删除key对应的value
	Del(string)
	//数据存储到cache
	Save()
	//设置过期时间
	SetExpired(int)
}

type SessionMgr interface {
	//初始化
	Init(addr string, options ...string) error
	//通过sessionid获取session对象
	GetSession(string) (Session, error)
	//创建session对象
	CreateSession() Session
	//根据sessionid删除session对象
	Clear(string)
}

type Options struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

func CreateSessionMgr(name SessionMgrType, addr string, options ...string) (sm SessionMgr, err error) {
	switch name {
	case File:
	case Memcached:
	case Redis:
	default:
		err := fmt.Sprintf("unsupported type %s\n", name)
		fmt.Print(err)
		break
	}
	return
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
