## Guideline

> 这是来自字节跳动团队cloudwego的hertz的demo，本demo主要学习了gorm，即go的orm框架

### Step1 使用hertz搭建一个简单的http框架

这里可以通过hz脚手架进行快速的搭建，首先我们需要编写一个idl文件，定义接口和数据模型

可见如下文件：
[idl](./gorm/idl/api.thrift)

这里采用的序列化协议是thrift，常见的还有google开源的proto

然后我们可以通过hz命令(这里需要事先安装好)

```bash
hz new --model_dir biz/hertz_gen -mod github.com/ggstudy11/learn_go_from_cloudwego/gorm -idl idl/api.thrift
```

model_dir 指的是生成代码的存放位置， mod指的是你想要的mod名

然后这样其实就给我们生成了大部分的代码，这里我们需要实现的是dal层，model层、handler层

### Step2 实现具体业务逻辑

model层即数据模型层，对应的是数据库中的数据模型

dal层即数据访问层，我们需要初始化mysql连接，同时实现一些数据库操作

其次就是handler，handler实现了业务逻辑的处理，通过接收参数，然后进行相应的业务逻辑实现

### Other

作为Java转go的选手，那么显然就会思考二者框架之间的差异和共同点，首先这里的hertz框架一般采用thrift序列化协议，同时提供了自动生成代码的脚手架

如果说要在hertz中找到和spring boot相似的，那么就是model层，dal层和handler层，显然这里handler层可以看作controller + service层的融合

这里少了很多dto vo等设计，如果需要数据模型的转换一般放到了pack中做一个函数式的转换
