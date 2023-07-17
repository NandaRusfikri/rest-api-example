#!/bin/bash
fuser -n tcp -k 54321
mv rest-api /root/rest-api/rest-api
nohup /root/rest-api/rest-api > api.log 2>&1 &