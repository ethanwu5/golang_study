 # 引入对应的依赖
```
go get gopkg.in/yaml.v3


```
# 读取文件内容
```
dataBytes, err := os.ReadFile("./conf.yaml")
```

# 定义对应struct 结构体
```
type T struct {
        F int `yaml:"a,omitempty"`
        B int
}

```

# 解析对应结构体
```
err = yaml.Unmarshal(dataBytes, &config)
```
