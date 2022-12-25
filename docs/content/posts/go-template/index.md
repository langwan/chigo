---
title: "go模板"
date: 2022-12-25T10:53:49+08:00
lastmod: 2022-12-25T10:54:35+08:00
---

## html/template

go 语言实现的数据驱动模板，用于生成可防止代码注入的 HTML 输出。它的功能与 text/template 基础功能相同，html/template 侧重于安全。如果与 http 服务结合可以输出 H5 页面。

### 基础

```go {github="https://github.com/langwan/chigo/blob/main/Template/Basic/main.go"}

```

```output
<h1>hello chihuo</h1>
```

`template.New()` 方法创建一个新的模板，不同的模板之间可以互相套用。`Parse()` 方法解析的是模板的内容。`Execute()` 绑定数据 `data` 并输出 HTML 内容。

### 模板文件

```html {github="https://github.com/langwan/chigo/blob/main/Template/TemplateFile/index.tmpl"}

```

```go {github="https://github.com/langwan/chigo/blob/main/Template/TemplateFile/main.go"}

```

```output
<h1>chihuo<h1>
```

index.tmpl 就是独立的模板文件

### 布局

布局 layout 实际上就是模板之间的嵌套调用

```{title="目录结构"}
.
├── main.go
├── pages
│   ├── about.html
│   └── home.html
└── tpl
    ├── footer.html
    ├── header.html
    └── layout.html
```

```go {github="https://github.com/langwan/chigo/blob/main/Template/Layout/main.go"}

```

我们用一个 gin 的 web 服务器来返回页面，访问 home 的时候返回主页，访问 about 的时候返回关于。

```html {github="https://github.com/langwan/chigo/blob/main/Template/Layout/tpl/layout.html"}

```

```html {github="https://github.com/langwan/chigo/blob/main/Template/Layout/tpl/header.html"}

```

```html {github="https://github.com/langwan/chigo/blob/main/Template/Layout/tpl/footer.html"}

```

以上是 layout 的内容

```html {github="https://github.com/langwan/chigo/blob/main/Template/Layout/pages/home.html"}

```

```html {github="https://github.com/langwan/chigo/blob/main/Template/Layout/pages/about.html"}

```

home 和 about 页的内容

```output {title="http://127.0.0.1:8100/home"}
<html>
    <head>
        <title>home</title>
    </head>
    <body>

<header>
    logo | header
</header>

        hello chihuo, this is home page

<footer>
    2022 by chihuo
</footer>

    </body>
</html>
```

```output {title="http://127.0.0.1:8100/about"}
<html>
    <head>
        <title>about</title>
    </head>
    <body>

<header>
    logo | header
</header>

        my name is chihuo

<footer>
    2022 by chihuo
</footer>

    </body>
</html>
```

### 变量

在模板内可以使用变量来完成一些复杂的操作。

```html {github="https://github.com/langwan/chigo/blob/main/Template/Variables/main.html"}

```

```go {github="https://github.com/langwan/chigo/blob/main/Template/Variables/main.go"}

```

```output
x = 10
y = 20
```

### 函数

可以注册自定义函数，hugo 就是这样注册进去的。

```go {github="https://github.com/langwan/chigo/blob/main/Template/Functions/main.go"}

```

```output
4
```

### 防注入

```go {github="https://github.com/langwan/chigo/blob/main/Template/InjectionSafe/main.go"}

```

```output
Hello, <script>alert('you have been pwned')</script>
!Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!
```

### 转义

```go {github="https://github.com/langwan/chigo/blob/main/Template/Escape/main.go"}

```

```output
&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
&#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
&#34;Fran &amp; Freddie&#39;s Diner&#34;32&lt;tasty@example.com&gt;
\"Fran \u0026 Freddie\'s Diner\" \u003Ctasty@example.com\u003E
\"Fran \u0026 Freddie\'s Diner\" \u003Ctasty@example.com\u003E
\"Fran \u0026 Freddie\'s Diner\"32\u003Ctasty@example.com\u003E
%22Fran+%26+Freddie%27s+Diner%2232%3Ctasty%40example.com%3E
```

## 视频

{{<bilibili "BV1H84y1s7dg" >}}
