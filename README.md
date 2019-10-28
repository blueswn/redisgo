# redisgo

## Basic Usage

### redis strings

     package main
     
     import (
        "context"
        "fmt"
        "github.com/blueswn/redisgo/fast"
        "github.com/gomodule/redigo/redis"
     )
     
     const _CACHE_PREFIX_USER_KEY = "project:user:id:%v"
     
     func CacheCronTime(ctx context.Context, userID uint64) *fast.Strings {
        conn, err := NewPool().GetContext(ctx)
        if err != nil {
            return nil
        }
     
        return fast.NewStrings(conn, fmt.Sprintf(_CACHE_PREFIX_USER_KEY, userID))
     }
     
     func main() {
        fs := CacheCronTime(context.Background(), 1)
        err := fs.SetUint64(1234)
        if err != nil {
            fmt.Println(err)
        }
        v, err := fs.GetUint64()
        if err != nil {
            fmt.Println(err)
        }
     
        fmt.Println(v)
     }
     
     func NewPool() *redis.Pool {
        pool := &redis.Pool{
            // Other pool configuration not shown in this example.
            Dial: func () (redis.Conn, error) {
                c, err := redis.Dial("tcp", "127.0.0.1:6379")
                if err != nil {
                    return nil, err
                }
                //if _, err := c.Do("AUTH", ""); err != nil {
                //  c.Close()
                //  return nil, err
                //}
                if _, err := c.Do("SELECT", "0"); err != nil {
                    c.Close()
                    return nil, err
                }
                return c, nil
            },
        }
     
        return pool
     }