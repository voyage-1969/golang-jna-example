# golang-jna-example
go语言jna示例

java调用方见：https://github.com/voyage-1969/java-jna-golang-example

构建命令：

    go build -buildmode=c-shared -o test.dll ./run.go
    
构建完成后复制dll文件至java resource中即可
