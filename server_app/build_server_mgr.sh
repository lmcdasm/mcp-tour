#!/bin/bash
# 
# build-server-manager.sh
#
# This builds a container image of the mcp-go-server-manager application 

app=mcp-go-server-manager
version=latest 
docker build -t ${app}:${version}  .
