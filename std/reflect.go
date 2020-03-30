package main

import (
    "reflect"
    "fmt"
)

type myInt int64

func reflectType(x interface{}) {
    t := reflect.TypeOf(x)
    fmt.Printf("type:%v, kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
    v := reflect.ValueOf(x)
    k := v.Kind()
    switch k {
    case reflect.Int64:
        // v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
        fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
    case reflect.Float32:
        // v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
        fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
    case reflect.Float64:
        // v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
        fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
    }
}

// 通过反射设置变量的值
func reflectSetValue(x interface{}) {
    v := reflect.ValueOf(x)
    // 反射中使用 Elem()方法获取指针对应的值
    if v.Elem().Kind() == reflect.Int64 {
        v.Elem().SetInt(200)
    }
}

func testSet() {
    var a int64 = 100
    reflectSetValue(&a)
    fmt.Println(a)
}

type student struct {
    Name  string `json:"name"`
    Score int    `json:"score"`
}

func (s student) Study() string {
    msg := "好好学习，天天向上。"
    fmt.Println(msg)
    return msg
}

func (s student) Sleep() string {
    msg := "好好睡觉，快快长大。"
    fmt.Println(msg)
    return msg
}



func testStudent() {
    stu1 := student{
        Name:  "小王子",
        Score: 90,
    }
    
    t := reflect.TypeOf(stu1)
    fmt.Println(t.Name(), t.Kind()) // student struct
    // 通过for循环遍历结构体的所有字段信息
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
    }
    
    // 通过字段名获取指定结构体字段信息
    if scoreField, ok := t.FieldByName("Score"); ok {
        fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
    }
}

func printMethod(x interface{}) {
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)
    
    fmt.Println(t.NumMethod())
    for i := 0; i < v.NumMethod(); i++ {
        methodType := v.Method(i).Type()
        fmt.Printf("method name:%s\n", t.Method(i).Name)
        fmt.Printf("method:%s\n", methodType)
        // 通过反射调用方法传递的参数必须是 []reflect.Value 类型
        var args = []reflect.Value{}
        v.Method(i).Call(args)
    }
}

func main() {
    var a float32 = 1.234
    reflectType(a)
    
    var b int8 = 10
    reflectType(b)
    
}
