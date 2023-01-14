::根据指定数据库生成model
@echo off
go install gorm.io/gen/tools/gentool@latest
set /p user="Username:"
set /p pwd="Password:"
set /p server="Server name and port:"
set /p db="Database name:"
set str="%user%:%pwd%@tcp(%server%)/%db%?charset=utf8mb4&parseTime=True"
gentool -dsn %str%
