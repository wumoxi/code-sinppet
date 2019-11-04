package main

import (
	"encoding/json"
	"fmt"
)

// 定义手机屏幕
type Screen struct {
	Size       float32
	ResX, ResY int
}

// 定义电池
type Battery struct {
	Capacity int // 容量
}

// 生成 JSON 数据
func generateJSONData() []byte {
	// 完整的数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			Capacity: 10000,
		},
		HasTouchID: true,
	}

	bytes, _ := json.Marshal(raw)

	return bytes
}

func main() {
	// 生成一段JSON 数据
	data := generateJSONData()

	fmt.Println(string(data))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	json.Unmarshal(data, &screenAndTouch)
	fmt.Printf("screenAndTouch: %+v\n", screenAndTouch)

	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	json.Unmarshal(data, &batteryAndTouch)
	fmt.Printf("batteryAndTouch: %+v\n", batteryAndTouch)
}
