目前在开发阶段, 临时用法如下:

```
git clone git@github.com:flaboy/generator-gorilla.git
```

npm link一下， 让本机知道 generator-gorilla 就在当前目录.
```
cd generator-gorilla
npm install
npm link
```

安装一下yoeman,  -g 表示全局.
```
npm install -g yo
```

随便找个空目录，执行yo
```
yo
```
应该能看到 gorilla 的选项。

之后，创建新项目可以使用:
```
yo gorilla
```

使用 https://start.jhipster.tech/jdl-studio/ 编写jdl文件. 
下载下来后, 可以自动生成所有CRUD的Model、Controller、View、Route、ApiController
```
yo gorilla:jdl jhipster-jdl.jdl
```

TODO:
1. Go禁止循环依赖， 生成的Models有互相引用的情况需要手工修改一下
