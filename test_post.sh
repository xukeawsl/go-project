#! /bin/bash

curl -v -X POST 'http://127.0.0.1:8080/community/page/post' \
-H 'Content-Type: application/json' \
-d '{"topic_id":3,"content":"不错哦"}'