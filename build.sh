cd gen_go/scraping_go
kitex --module github.com/For-December/trader_idl/gen_go/scraping_go -service scraping_service  ../../idl/scraping_go/service.thrift 

cd gen-py/brain_py
thriftgo -r -out . --gen py ../../idl/scraping_go/service.thrift

thrift --gen py -out . ../../idl/scraping_go/service.thrift
