#!/bin/bash
# run_server_manager.sh 

# This runs a locally built docker instance of the MCP-GO-Server-Manager and allows your to test your built package.
# This runs very privileged (host network mode, privileged, etc) and its not meant for production. Use the K8s Envelope for more 'production' ready packaging and runtime goodness

app=mcp-go-server-manager
version=latest 
docker run -d --name mcp-go-server-manager \
	--privileged \
	--restart=always \
	--network=host \
	${app}:${version}
