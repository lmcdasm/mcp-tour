= MCP Tour
:toc: left
:icons: font
:sectnums:
:source-highlighter: highlight.js
:imagesdir: resources
:docdir: resources
:banner-caption: A Go-Lang Exploration of Model Context Protocol (MCP)
:docfile: docs/mcp-tour-architecture.pptx

[.lead]
A Tour of the Model Context Protocol (MCP) using a Go-based implementation.

image::mcp-architecture.png[Architecture Diagram, width=800]

[NOTE]
====
This project builds on top of MCP's early concepts, including a Go-based SDK.
Please consult the `LICENSES` directory for full license attributions (MIT/BSD) for Google and Anthropic work.
====

link:docs/mcp-tour-architecture.pptx[Download the Architecture Overview (PPT)]

image::ux-preview.png[UX Preview, width=700]

== Goal

To experiment with the under-development MCP Go SDK and learn more about MCP’s structure and implementation patterns.

This project also serves as a boilerplate and sandbox to explore roadmap concepts from the official MCP homepage.

== Project Layout

* `client_app/` – Go-based MCP client that connects to an MCP server and interacts using the protocol.
* `server_app/` – Go-based MCP server that handles requests from clients, optionally routes to LLMs or external services.
* `mcp_src/` – Sourced from https://go.googlesource.com/tools/internal/mcp – adjusted for standalone use.

== How to Use

. Clone the repository.
. Build the following components:
  * `server_app`
  * `client_app`
  * `registry_app`
  * `explorer_ux`
. Run all four applications.

=== Accessing Applications

Access via the following endpoints once running (adjust for your IP/domain):

* *Explorer UX*: https://<your-ip>:9000/
* *Server Manager*: https://<your-ip>:10010/swagger/index.html
* *Client Manager*: https://<your-ip>:10011/swagger/index.html
* *Registry Manager*: https://<your-ip>:10012/swagger/index.html

== External Libraries

The MCP source tree (`mcp_src/`) includes code from Google's internal toolchain at `googlesource.com/tools/internal`.

Adjustments were limited to import paths and compatibility fixes. All original licensing headers and notices have been preserved.

Refer to the `LICENSES/` directory for BSD and MIT attributions used across this project.

== Resources

* link:docs/mcp-tour-architecture.pptx[MCP Architecture Overview (PPT)]
* image::ux-preview.png[Main Explorer UX Screenshot]

