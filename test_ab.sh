#! /bin/bash

# 使用 ab 测试工具
ab -n 10000 -c 3000 \
-p postfile \
-T 'application\json' \
'http://127.0.0.1:8080/community/page/post' 