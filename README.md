# gorm-gener
通过yaml配置，基于gorm generator 完成对于数据库表的领域模型的生成

## yaml 文件配置
~~~
  Model:
  - Dsn: 'postgres://username:password@yourhost:port/database'
    Table:
      yourTableName: yourTargetObjectName
  - Dsn: 'postgres://username:password@yourhost:port/database'
    Table:
      yourTableName: yourTargetObjectName
~~~

## 工具使用
~~~
go install .
gorm-gener gen -c “your config file path”
~~~
