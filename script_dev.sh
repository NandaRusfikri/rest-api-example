#!/bin/bash
sudo fuser -n tcp -k 54321
nohup ./rest-api > api.log 2>&1 &