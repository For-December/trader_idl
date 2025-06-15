cd gen_go/scaping_go
kitex --module github.com/For-December/trader_idl/gen_go/scaping_go -service caping_service  ../../idl/scaping_go/service.thrift 

cd gen-py/brain_py
thriftgo -r -out . --gen py ../../idl/scaping_go/service.thrift