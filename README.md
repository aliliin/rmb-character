# 将阿拉伯数字金额转换为汉字的形式 #


Convert Arabic numeric amounts to Chinese character form. 

#### 安装使用

##### Golang 版本大于等于1.16

```go

go get -u github.com/aliliin/rmb-character

import (
    "github.com/aliliin/rmb-character"
)

```

#### 注意注意

> 目前只能输入整数金额且输入的金额会除以 100 进行计算 如：输入 15809 得到的结果为：壹佰伍拾捌圆零角玖分

##### 用法示例

```go

fmt.Println(rmb_character.RmbCapitalCharacters(1234340)) // 壹万贰仟叁佰肆拾叁圆肆角零分

fmt.Println(rmb_character.RmbCapitalCharacters(567843)) // 伍仟陆佰柒拾捌圆肆角叁分
```


##### 错误处理

> 开发中...


#### 参考项目

* [chuoke/rmb-capital](https://github.com/chuoke/rmb-capital)