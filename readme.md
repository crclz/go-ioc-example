# go-ioc-example

## wire

examples that use https://github.com/google/wire

```bash
go install github.com/google/wire/cmd/wire@latest

wire ./wire/...
wire ./wire/singleton
wire ./wire/scoped
```

directory structure:
- wire/singleton: basic usage of wire
- wire/scoped: use wire for singleton and scoped dependencies
