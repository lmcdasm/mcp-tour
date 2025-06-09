#!/bin/bash
# 
# build-server-manager.sh
#
# This logs the mcp-go-server-manager (assumed built from the build-server-manager and ran from run-server-manager scripts) 

docker logs -f mcp-go-server-manager
