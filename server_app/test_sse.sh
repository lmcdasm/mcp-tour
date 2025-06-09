#!/bin/bash

SSE_URL="http://localhost:10010/api/stream"

echo "Connection to $SSE_URL ... "
curl -N -H "Accept: text/event-stream" "$SSE_URL"
