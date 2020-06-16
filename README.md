#### 克隆项目
```
# 进入到$GOPATH/src, 注意不要使用go get 项目地址
cd $GOPATH/src

git clone 项目地址 go-xb-song
 
cd go-xb-song
```

#### 使用govendor管理包

```
# 安装govendor
go get -u github.com/kardianos/govendor
```


#### todo
```
   1、制作makefile，生成bin文件到指定目录
   2、实现reload
```