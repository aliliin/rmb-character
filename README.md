# RMB-Character Convert RMB amount to Chinese uppercase.  #


将阿拉伯数字金额转换为汉字的形式


#### 安装使用

##### Golang 版本大于等于1.16

```go
// 使用 github 库
go get -u github.com/golang-module/carbon/v2

import (
    "github.com/aliliin/rmb-character"
)

```

#### 注意注意

> 目前只能输入整数金额且输入的金额会除以 100 进行计算 如：输入 15809 得到的结果为：壹佰伍拾捌圆零角玖分

##### 用法示例

```go
// 今天此刻
fmt.Sprintf("%s", RmbCapitalCharacters(123456)) // 2020-08-05 13:14:15

```


##### 错误处理

> 开发中...


#### 参考项目

* [chuoke/rmb-capital](https://github.com/chuoke/rmb-capital)