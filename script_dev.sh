#!/bin/bash
fuser -n tcp -k 54321
nohup /root/rest-api/rest-api > api.log 2>&1 &