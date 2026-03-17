# StoryBuilder
---
StoryBuilder API

## Generating API Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) to generate Swagger/OpenAPI documentation into the `docs/` folder.

To regenerate the docs after updating handlers or comments:

```sh
# Install the swag CLI (once)
go install github.com/swaggo/swag/cmd/swag@latest

# Regenerate docs
go generate ./...
```
