# mcp-tour

A Tour of Model Context Protocol (MCP) using Go-Lang implementation

!! Please ensure to look in LICENSES to find all the deserved credits (MIT) to Google and Anthropic. !!

## Goal

To experiement with the currently under-development MCP GoSDK and at the same time learn a bit about MCP and the implementations.

In the process, we will also setup a boilerplate to allow us to explore the Roadmapped items from the MCP homepage.

### Project Layout

- client_app - this is a go based MCP client application that will connect, using MCP, to a MCP Server
- server_app - this is a go-based MCP Server application that talks MCP to clients and does things (LLM and/or calls to services, etc.)
- mcp_src - this is a copy/paste out of https://go.googlesource.com/tools/internal/mcp


### How to Use

- Clone project and decide if you want to build the client or the server first and "cd" into the appropriate driectory
- run "go build -o mcp-client" (or mcp-server if you are doing the server") and build the binary
- run the binary with "./mcp-client" (or ./mcp-server") to start the app.

Note: You should generally start the server "first" so the client has something to connect "to" 

### Other Libraries "hacked-in" 

The MCP internal project uses other libraries from within googlesource.com/tools/internal packages.  The LICENSES directory has the BSD-License that is used with internal projects from Google, all libraries have had only their import statements changes to function in most cases.

In all cases, the appropriate slogans/headers and other markups are all left intact.
The LICENSES directory contains all liceneses inherited from 3rd parties 


