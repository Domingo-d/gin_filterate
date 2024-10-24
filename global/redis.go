package global

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

func HGet[T any](key string, field string) (*T, error) {
	data := new(T)
	ret, err := Redis.HGet(context.TODO(), key, field).Bytes()
	if err != nil {
		return nil, err
	}

	if err := jsoniter.Unmarshal(ret, data); nil != err {
		Logger.Error("jsoniter.Unmarshal error", zap.Error(err), zap.String("key", key), zap.String("field", field))
		return nil, err
	}

	return data, nil
}

func HSet[T any](key string, field string, data T) error {
	buf, err := jsoniter.Marshal(data)
	if nil != err {
		Logger.Error("jsoniter.Marshal error", zap.Error(err), zap.String("key", key), zap.String("field", field))
		return err
	}

	return Redis.HSet(context.TODO(), key, field, buf).Err()
}
