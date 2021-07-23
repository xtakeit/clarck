package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/joho/godotenv"
)

var envMap map[string]interface{}

var once sync.Once

func LoadConfigFromFileEnv(filepath string, template interface{}) {
	once.Do(func() {
		stringEnv, err := godotenv.Read(filepath)
		if err != nil {
			panic("配置文件[" + filepath + "]未找到")
		}
		parseAsMap(stringEnv)
	})

	parse(template, envMap)
}

func parseAsMap(baseEnv map[string]string) {
	envMap = make(map[string]interface{})
	for k, v := range baseEnv {
		keys := strings.Split(k, ".")
		constructMap(keys, v, envMap)
	}
}

func constructMap(keys []string, value string, root map[string]interface{}) {
	// 循环创建容器
	beforeLast := root
	for _, key := range keys {
		_, ok := root[key]
		if !ok {
			beforeLast = root
			root[key] = make(map[string]interface{})
		}
		root = root[key].(map[string]interface{})
	}
	//给最后一个容器设置值
	beforeLast[keys[len(keys)-1]] = value
}

func parse(template interface{}, root map[string]interface{}) {
	vals := reflect.ValueOf(template)
	if vals.Kind() == reflect.Ptr {
		vals = vals.Elem()
	}
	if !vals.CanSet() {
		panic(fmt.Sprintf("对象%+v不可达，请传入指针类型", vals.Type()))
		return
	}
	parseVal(vals, root)
}

func parseVal(value reflect.Value, root map[string]interface{}) {
	// 根据类型递归设置值
	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fmt.Println(value.Type().Field(i).Name)
			key, ok := hasKey(value.Type().Field(i).Name, root)
			if !ok {
				continue
			}
			if justBaseTypeAndSet(root[key], value.Field(i)) {
				continue
			}
			parseVal(value.Field(i), root[key].(map[string]interface{}))
		}
	case reflect.Map:
		valMap := reflect.MakeMap(value.Type())
		for k, v := range root {
			// 遍历
			kind := value.Type().Elem()
			val := reflect.New(kind).Elem()
			if !justBaseTypeAndSet(v, val) {
				parseVal(val, v.(map[string]interface{}))
			}
			valMap.SetMapIndex(reflect.ValueOf(k), val)
		}
		value.Set(valMap)
		//TODO 切片类型还没有解析,基本数组
	default:
		panic("不允许直接传入普通类型构造")
	}

}

// 匹配 结构体名字
func hasKey(name string, root map[string]interface{}) (key string, ok bool) {
	// 直接匹配
	key = name
	if _, ok = root[key]; ok {
		return key, ok
	}
	//首字母小写匹配
	key = toLowerFirst(name)
	if _, ok = root[key]; ok {
		return key, ok
	}
	//驼峰转下划线
	key = camelToBaseLine(name)
	if _, ok = root[key]; ok {
		return key, ok
	}
	//驼峰转横杠
	key = camelToLine(name)
	if _, ok = root[key]; ok {
		return key, ok
	}
	return "", false
}

// 首字母小写
func toLowerFirst(str string) string {
	runes := []rune(str)
	if !unicode.IsUpper(runes[0]) {
		return str
	}
	// 转换为小写
	runes[0] = runes[0] + 32
	return string(runes)
}

// 驼峰转下划线
func camelToBaseLine(str string) string {
	baseLine := rune(95) // 下划线(_)的unicode码
	runes := []rune(str)
	res := make([]rune, 0)

	//首字母小写
	if unicode.IsUpper(runes[0]) {
		runes[0] = runes[0] + 32
	}
	// 其他大写字母，转下划线格式
	for _, one := range runes {
		if unicode.IsUpper(one) {
			res = append(res, baseLine, one+32)
		} else {
			res = append(res, one)
		}
	}
	return string(res)
}

// 驼峰转横杠
func camelToLine(str string) string {
	baseLine := rune(45) // 横杠(-)的unicode码
	runes := []rune(str)
	res := make([]rune, 0)

	//首字母小写
	if unicode.IsUpper(runes[0]) {
		runes[0] = runes[0] + 32
	}
	// 其他大写字母，转下划线格式
	for _, one := range runes {
		if unicode.IsUpper(one) {
			res = append(res, baseLine, one+32)
		} else {
			res = append(res, one)
		}
	}
	return string(res)
}

// 如果当前template是基础类型，设置值，返回true
// 如果是map，或者value是结构体，返回false
func justBaseTypeAndSet(template interface{}, value reflect.Value) bool {

	// 同时不是基础类型时才需要遍历
	if !valueIsBaseKind(value) && !interfaceIsBaseKind(template) {
		return false
		//同时是基础类型时进行设置值
	} else if valueIsBaseKind(value) && interfaceIsBaseKind(template) {
		val := getValue(template.(string), value.Kind())
		value.Set(val)
	}
	return true
}

func valueIsBaseKind(value reflect.Value) bool {
	return isBaseKind(value.Kind())
}

func interfaceIsBaseKind(template interface{}) bool {

	typ := reflect.TypeOf(template)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return isBaseKind(typ.Kind())
}

func isBaseKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		return false
	default:
		return true
	}
}

// 通过目标的kind，然后转换为对应的类型value
func getValue(str string, kind reflect.Kind) reflect.Value {

	switch kind {
	case reflect.String:
		return reflect.ValueOf(str)
	case reflect.Bool:
		val, _ := strconv.ParseBool(str)
		return reflect.ValueOf(val)
	case reflect.Int8:
		val, _ := strconv.Atoi(str)
		return reflect.ValueOf(int8(val))
	case reflect.Int16:
		val, _ := strconv.Atoi(str)
		return reflect.ValueOf(int16(val))
	case reflect.Int32:
		val, _ := strconv.Atoi(str)
		return reflect.ValueOf(int32(val))
	case reflect.Int64:
		val, _ := strconv.Atoi(str)
		return reflect.ValueOf(int64(val))
	case reflect.Int:
		val, _ := strconv.Atoi(str)
		return reflect.ValueOf(val)
	case reflect.Uint8:
		val, _ := strconv.ParseUint(str, 10, 8)
		return reflect.ValueOf(uint8(val))
	case reflect.Uint16:
		val, _ := strconv.ParseUint(str, 10, 16)
		return reflect.ValueOf(uint16(val))
	case reflect.Uint32:
		val, _ := strconv.ParseUint(str, 10, 32)
		return reflect.ValueOf(uint32(val))
	case reflect.Uint:
		val, _ := strconv.ParseUint(str, 10, 64)
		return reflect.ValueOf(uint(val))
	case reflect.Uint64:
		val, _ := strconv.ParseUint(str, 10, 64)
		return reflect.ValueOf(val)
	case reflect.Float32:
		val, _ := strconv.ParseFloat(str, 32)
		return reflect.ValueOf(float32(val))
	case reflect.Float64:
		val, _ := strconv.ParseFloat(str, 64)
		return reflect.ValueOf(val)
	default:
		panic("非基本类型，不允许直接设置")
	}
}
