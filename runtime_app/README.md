# Runtime Manager

This is a component that runs and implements the southbound "hooks" towards whatever infrastrcture that you want to your MCP Servers to "run on".

### Supported Runtime Infrastructures (Target)

- docker - single host 
 - assumes that the runtime container is running on the same docker host and has privileged rights - really for development
- k8s 
 - this is our target, since we want to model the transmission between the server and the runtime manager via a K8s CRD ideally.
 - however, this assumes that we can somehow make a "on the fly CRD writer" that can wrap a abstracted Mcp Server Model, but that is the idea.

