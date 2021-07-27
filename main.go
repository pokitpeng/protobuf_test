package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"protobuf_test/p3optional"
	"protobuf_test/wrapper"
)

func main() {
	// TestOptionJsonToStruct()
	// TestOptionStructToJson()

	TestWrapperJsonToStruct()
}

func TestOptionJsonToStruct() {
	s := `{
    "network": "tcp",
    "addr": "127.0.0.1",
	"db": 0
	}`
	var redis p3optional.Redis
	err := unmarshalJSON([]byte(s), &redis)
	if err != nil {
		fmt.Printf("Option json --> struct err: %s", err.Error())
		return
	}
	fmt.Printf("redis: %+v\n", redis)
	fmt.Println("redis.Db:", *redis.Db)
	// 	如果没有写option关键字，没有传的参数值，反序列化后为零值
	// 	如果写了option关键字，没有传的参数值，反序列化后为nil
	// 	如果写了option关键字，传的参数值，反序列化后都为指针，需要解引用才能得到值
}

func TestOptionStructToJson() {
	var db int32 = 0
	var redis = p3optional.Redis{
		Network:      "",
		Addr:         "127.0.0.1",
		Password:     "",
		Db:           &db,
		DialTimeout:  nil,
		ReadTimeout:  nil,
		WriteTimeout: nil,
	}
	s, err := marshalJSON(&redis)
	if err != nil {
		fmt.Printf("Option struct --> json err: %s", err.Error())
		return
	}
	fmt.Printf("redis: %s\n", string(s))
	// 如果没有写option关键字，没有传的参数值，序列化后key直接去除
	// 如果写了option关键字，没有传的参数值，序列化后key直接去除
	// 如果写了option关键字，传的参数值，必须为指针类型
}

func TestWrapperJsonToStruct() {
	s := `{
    "network": "tcp",
    "addr": "127.0.0.1",
	"db": 0,
	"dial_timeout": "1s"
	}`
	var redis wrapper.Redis
	err := unmarshalJSON([]byte(s), &redis)
	// err := json.Unmarshal([]byte(s), &redis)
	if err != nil {
		fmt.Printf("Wrapper json --> struct err: %s", err.Error())
		return
	}
	fmt.Printf("redis: %+v\n", redis)
	fmt.Println("redis.Db:", redis.Db)
	fmt.Println("redis.Password:", redis.Password)
	// 没有wrap的字段，没有传参数的值，反序列化后是零值
	// wrap的字段，没有传参数的值，反序列化后是nil
	// wrap的字段，传参数的零值，反序列化后是零值
}

func marshalJSON(v interface{}) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(m)
	}
	return json.Marshal(v)
}

func unmarshalJSON(data []byte, v interface{}) error {
	if m, ok := v.(proto.Message); ok {
		return protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(data, m)
	}
	return json.Unmarshal(data, v)
}
