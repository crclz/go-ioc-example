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


## dependency locator (WIP)

Dependency locator is an anti pattern. However, some projects use it to simplify development. For example, get a scoped service from context.Context in a singleton service.

directory structure:
- locator: library for dependency locator
- locator/example: example
