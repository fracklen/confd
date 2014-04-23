#!/bin/sh

cd config
go test
cd ..

cd etcd/etcdtest
go test
cd ../..

cd etcd/etcdutil
go test
cd ../..

cd resource/template
go test
cd ../..
