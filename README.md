# dsp 技术开发规范

**线上go版本(=go1.23.0) |**
**Mysql版本(8.0.21)**

1、**文件命名规范**：

    1、同级的文件遵循同级规则
    2、自定义类遵循大驼峰命名规范
    3、结构属性命名尽量使用大驼峰

2、**代码编写规范**：

    1、等号后须有空格
    2、类、方法不能出现数字，尽量避免拼音命名

3、**注释规范**:

    1、接口、方法、结构体、必须有注释
    2、代码内部注释：
        1、方法内如果模块比较大遵循: /**xxx**/
        2、普通注释遵循双斜线规范: //xxx

4、**go mod中新增包在内部群先通知或者先沟通**.

5、**mysql操作规范**:

    1、表操作都在migrate下执行，包含表、字段、索引的新增、添加、修改、删除
    2、表字段符合释义、不能出现大写，尽量避免拼音，多字符以下划线分割，表名和字段必须加注释
    3、查询尽量避免使用原生，使用gorm查询
    4、常见方法封装在表对应的model里面，
    5、表通用字段：
        created_id : 创建者id
        created_at : 创建时间，创建时自动添加
        updated_at : 修改时间，修改时自动添加
        deleted_at : 删除时间，删除时自动添加
        status     : 表示状态值
        is_del     : 是否删除 1正常 2删除
        material_type : 素材类型 1图片 2视频
    6、数据迁移使用migrate，不使用任何sql操作
        6.1、安装migrate迁移工具：https://github.com/golang-migrate/migrate/releases
        6.2、命令行参考：https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
        6.3、创建数据迁移文件，如：migrate create -ext sql -dir ./migrations -seq create_users_table
        6.4、执行迁移：migrate -database "mysql://root:asdf1234@tcp(127.0.0.1:3306)/dsp?charset=utf8mb4&parseTime=True&loc=Local" -path ./migrations up
        6.5、执行回滚：migrate -database 'mysql://root:asdf1234@tcp(127.0.0.1:3306)/dsp?charset=utf8mb4&parseTime=True&loc=Local' -path ./migrations down
            如果回滚最后一步后面加1 ：migrate -database 'mysql://root:asdf1234@tcp(127.0.0.1:3306)/dsp?charset=utf8mb4&parseTime=True&loc=Local' -path ./migrations down 1 
        6.6、执行数据迁移后在model中创建数据表对应的结构

