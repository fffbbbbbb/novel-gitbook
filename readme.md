## 小说分段生成gitbook

### 环境变量
* filename 文件名，默认为`1.txt`

* expression 分段依据的正则表达式 默认为`第.*章 `。

### docker
build dockerfile后，使用docker将小说文件挂载到/gitbook-build，run就成了
