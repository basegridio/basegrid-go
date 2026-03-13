# BaseGrid Go SDK

The official Go client for [BaseGrid](https://basegrid.io), the memory infrastructure for AI agents.

## Installation

```bash
go get github.com/basegrid-io/basegrid-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/basegrid-io/basegrid-go"
)

func main() {
    client := basegrid.New("bg_your_api_key")

    // Add a memory
    memory, err := client.Add(basegrid.AddParams{
        AgentID: "agent-123",
        Content: "User prefers dark mode",
        Metadata: map[string]interface{}{
            "preference": "dark_mode",
        },
    })
    if err != nil {
        panic(err)
    }

    // Search memories
    results, err := client.Search(basegrid.SearchParams{
        AgentID: "agent-123",
        Query:   "What are the user preferences?",
    })
    if err != nil {
        panic(err)
    }

    for _, m := range results {
        fmt.Println(m.Content)
    }
}
```

## Configuration

```go
client := basegrid.New("bg_your_api_key")
// Default BaseURL: https://basegrid-production.up.railway.app

// Or customize:
client.BaseURL = "https://your-custom-endpoint.com"
```

## Troubleshooting

### Connection Refused Error
If you see connection refused errors, ensure you're using the latest SDK version:
```bash
go get -u github.com/basegrid-io/basegrid-go
```

Earlier versions may have used an incorrect default API endpoint.

## Documentation

Full documentation: https://basegrid.io/docs
