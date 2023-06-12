目前在开发阶段, 临时用法如下:

```
git clone git@github.com:flaboy/generator-gorilla.git
```

npm link一下， 让本机知道 generator-gorilla 就在当前目录.
```
cd generator-gorilla
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

使用 https://start.jhipster.tech/jdl-studio/ 编写jdl文件. 下载下来后, 可以:
```
yo gorilla:jdl jhipster-jdl.jdl
···