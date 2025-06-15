#!/bin/bash

# 定义项目根目录和IDL文件路径
PROJECT_ROOT=$(cd "$(dirname "$0")" && pwd)
IDL_FILE="$PROJECT_ROOT/idl/scraping_go/service.thrift"
echo idl file is $IDL_FILE
cd $PROJECT_ROOT

cd gen_go/scraping_go
kitex --module github.com/For-December/trader_idl/gen_go/scraping_go -service scraping_service  ../../idl/scraping_go/service.thrift 
cd ../../
cd gen_py/brain_py
thrift --gen py -out . ../../idl/scraping_go/service.thrift
