#!/usr/bin/env bash
set -ex
rm -rf ./tmp
mkdir -p ./tmp/v2lbsyun
scp -r ../v2lbsyun/Dockerfile ./tmp/
scp -r ../v2lbsyun/docs.py ./tmp/
scp -r ../v2lbsyun/modules ./tmp/v2lbsyun/
scp -r ../v2lbsyun/public ./tmp/v2lbsyun/
scp -r ../v2lbsyun/static ./tmp/v2lbsyun/
scp -r ../v2lbsyun/__init__.py ./tmp/v2lbsyun/
scp -r ../v2lbsyun/jsapiAPI.json ./tmp/v2lbsyun/
scp -r ../v2lbsyun/main.py ./tmp/v2lbsyun/
scp -r ../v2lbsyun/requirements.txt ./tmp/v2lbsyun/
scp -r ../v2lbsyun/README.md ./tmp/v2lbsyun/

cd ./tmp

docker build -t v2lbsyun .
sleep 2
rm -rf ../tmp


