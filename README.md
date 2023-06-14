# 一个借鉴了jhipster 与 ruby on rails 的代码生成器

* jhipster的jdl很棒，可以很方便的描述模型并且快速生成可用的CRUD，但是java的分层不适合小规模团队，且生成的代码有一点学习成本。
* ruby on rails生成的代码简单，复合直觉，容易修改。 但是生态逐渐衰落，而且ruby性能也是个问题。

JDL的介绍:
  *  https://start.jhipster.tech/jdl-studio/
  *  https://www.jhipster.tech/jdl/intro

于是有了本项目!

## 特点
* 基于Golang的Revel Framework，性能高，省内存
* 可以根据JDL生成CRUD的代码Generator
* 生成的代码非常简单和复合直觉，不需要学习，看着上下文就会改。类似Ruby on rails风格
* 低下限： 默认的View为bootstrap的纯HTML,  意味着一个人可以快速写完一切。
* 高上限： 内建Vite与前端项目的基础代码，可以针对某个页面使用Vue增强，甚至全站前后端分离。 

## 使用方法

### 目前在开发阶段, 临时安装方式

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

### 使用方式

创建新项目可以使用:
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
