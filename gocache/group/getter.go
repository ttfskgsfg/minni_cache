package group

// 当缓存中找不到数时让用户去数据库找数据
type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

// 类似 http.Handler
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
