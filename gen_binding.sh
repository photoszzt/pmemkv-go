#!/bin/bash
c-for-go -ccdefs -ccincl -out . pmemkv.yml
sed -i 's/) byte/) int8/g' pmemkv_sys/pmemkv_sys.go
sed -i 's/(byte)/(int8)/g' pmemkv_sys/pmemkv_sys.go
