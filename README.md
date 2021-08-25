# SimpleLogging
This is a simple to use logging module that enable multi level and format logging without any unnecessary additions.

## How to use it
First import and initialize the module. 
Use the `Initialize` function to set lo *level*, *format* and (optionally) *output file* at once. 
Otherwise, use the *set* functions as shown in the code snip below.
```go
func main() {
  // Set the important configuration with one call
	slog.Initialize(
		slog.LEVEL_DEBUG,
		slog.FORMAT_JSON,
		slog.LogFiles{Path: "/tmp/logs/", "_my_app"},
	)

  // Set/Change the log level
  // Default is "LEVEL_INFO"
  slog.SetLogLevel(slog.LEVEL_INFO)

  // Set/Change the log format
  // Default is "FORMAT_JSON"
  slog.SetLogFormat(slog.FORMAT_JSON)

  // Set/Change the log file
  // Default is not using file logging
  slog.SetLogFiles(slog.LogFiles{Path: "", Prefix: ""})

  // Set/Change the log time format
  // Default is "2006-01-02 15:04:05" (YYYY-MM-DD HH:mm:ss)
  slog.SetTimeFormat("2006-01-02 15:04:05")
}
```

### Log level
| Value | Enum          | Meaning                                               |
|-------|---------------|-------------------------------------------------------|
| 0     | LEVEL_NONE    | Do no output at all                                   |
| 1     | LEVEL_FATAL   | Log only on fatal error                               |
| 2     | LEVEL_ERROR   | Log on error or above                                 |
| 3     | LEVEL_WARNING | Log on warnings or above                              |
| 4     | LEVEL_INFO    | Log on information or above                           |
| 5     | LEVEL_DEBUG   | Log on debug or above                                 |
| 6     | LEVEL_TRACE   | Log on trace or above (e.g. more detailed than debug) |

### Log format
| Value | Enum        | Meaning     |
|-------|-------------|-------------|
| 0     | FORMAT_JSON | Log as JSON |
| 1     | FORMAT_TEXT | Log as text |

## Log types
As show in the table **Log level**, there are multiple log levels.
All these log functions can be used in the same way.
They are fired based on the log level.
```go
slog.Fatal(slog.Entry{"event": "Hello world"})
slog.Error(slog.Entry{"event": "Hello world"})
slog.Warning(slog.Entry{"event": "Hello world"})
slog.Info(slog.Entry{"event": "Hello world"})
slog.Debug(slog.Entry{"event": "Hello world"})
slog.Trace(slog.Entry{"event": "Hello world"})
```

### Entry helper
There are also helper to create an entry.
E.g. `EventEntry` to define a entry as above.
You can see the available helper by using autocomplete after typing `slog.Entry`.
```go
// Both has the same output
slog.Debug(slog.EntryEvent("Hello world"))
slog.Debug(slog.Entry{"event": "Hello world"})

// Both has the same output
slog.Trace(slog.EntryMessage("Hello world"))
slog.Trace(slog.Entry{"message": "Hello world"})

err := fmt.Errorf("Hello world")
// Both has the same output
slog.Error(slog.EntryError(err))
slog.Error(slog.Entry{"error": err.Error()})
```