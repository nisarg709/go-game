package service

import (
	"net/http"
	"github.com/go-redis/redis"
	"github.com/globalsign/mgo"
)

type ServiceHandler struct {
	redis  *redis.Client
	config *Configuration
	buffer *Buffer
}

func NewService(config *Configuration) *ServiceHandler {

	s := new(ServiceHandler)
	s.config = config

	s.redis = redis.NewClient(&redis.Options{
		Addr:     s.config.RedisEndpoint,
		Password: s.config.RedisPassword,
		DB:       s.config.RedisDatabase,
	})

	s.buffer = new(Buffer)

	return s
}

func (h *ServiceHandler) Handle(w http.ResponseWriter, r *http.Request) error {

	var data []byte
	var etag = r.Header.Get("If-None-Match")

	if h.buffer.IsBufferedDataFresh() {

		if h.buffer.isEtagLatest(etag) {
			w.WriteHeader(http.StatusNotModified)
			return nil
		}

		h.PrintData(w, h.buffer.GetBufferedData())
		return nil
	}

	if h.isEtagLatest(etag) {
		h.buffer.ExtendExpiration(etag)
		w.WriteHeader(http.StatusNotModified)
		return nil
	}

	data, err := h.getCachedJson()

	if err != nil {
		data, err = h.GetFreshData()

		if err != nil {
			return err
		}
	}

	h.buffer.SetBufferedData(data)
	h.PrintData(w, data)

	return nil
}

func (h *ServiceHandler) isEtagLatest(etag string) bool {
	if etag == "" {
		return false
	}

	latest, err := h.redis.Get("etag").Result()

	if err != nil || etag != latest {
		return false
	}

	return true
}

func (h *ServiceHandler) getCachedJson() ([]byte, error) {

	current, err := h.redis.Get("data").Result()

	if err != nil {
		return nil, err
	}

	return []byte(current), nil
}

func (h *ServiceHandler) PrintData(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Etag", Etag(data))
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func (h *ServiceHandler) GetFreshData() ([]byte, error) {

	session, err := mgo.Dial(h.config.MongoEndpoint)
	if err != nil {
		return nil, err
	}

	repository := Repository{
		session.DB(h.config.MongoDatabase),
	}

	data, err := repository.GetFreshData()

	if err != nil {
		return nil, err
	}

	h.redis.Set("data", data, 0)
	h.redis.Set("etag", Etag(data), 0)

	return data, nil
}
