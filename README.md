# Template for cobra

input in term

```bash
go build -v && ./t4cobra --config.file=app.yaml child --child.str=changed
```

out

```text
INFO[2019-01-17T18:38:09+08:00] reading from 127.0.0.1:2379
WARN[2019-01-17T18:38:09+08:00] current log level: info
INFO[2019-01-17T18:38:09+08:00] settings on child: {Config:{File:app.yaml} Child:{Bool:true Test:false Str:changed}}
```
