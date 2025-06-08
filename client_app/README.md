# MCP CLIENT Implementations

Each of the directories has a client implementation using a different transport.  

STDIO - this using stdin/stdout as transport for the JSONRPC messages for MCP
IN-MEMORY - this was tinkering around with a single (Client/Server) test example from the Examples in the tools/internal/mcp from GoogleSource

APP - this is our main SSE/HTTP based client implementation "runner" and we will in fact use to have a MCP Client Manager for a Multi Client Approach overall.  this is based from the http example in the source repo for transport reference.



