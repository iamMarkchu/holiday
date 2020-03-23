# 判断节假日包

## 下载
```golang
go get -u github.com/iamMarkchu/holiday

```

```golang
date := "20200323"
result, err := holiday.IsHoliday(date)
if err != nil {
    fmt.Println(err)
}
if result {
    fmt.Println("当前是节假日或者周末")
} else {
    fmt.Println("当前不是节假日或者周末")
}
```
