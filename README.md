# gosched-simulator
this repository is experiment tool to visualize Go scheduler.

## Development

### goverlay

layers:
  - from: /path/to/go/src/runtime/proc.go
    patch: /path/to/gosched-simulator/runtime/_proc.go
    dist: /path/to/gosched-simulator/runtime/proc.go
  - from: /path/to/go/go/src/runtime/runtime2.go
    patch: /path/to/gosched-simulator/runtime/_runtime2.go
    dist: /path/to/gosched-simulator/runtime/runtime2.go

### overla

{
  "Replace":{
    "/path/to/go/src/runtime/proc.go":"./runtime/proc.go",
    "/path/to/go/src/runtime/runtime2.go":"./runtime/runtime2.go"
  }
}

