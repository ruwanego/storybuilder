# StoryBuilder
---
StoryBuilder API

## API Documentation

This project uses [Huma](https://huma.rocks/) to automatically generate live OpenAPI documentation based on the strongly-typed request and response structs defined in the controllers.

The interactive documentation is automatically served by the running application:
- **Swagger UI / Elements / Redoc**: Available at `/docs`
- **OpenAPI 3.1 JSON Specification**: Available at `/openapi.json`
- **OpenAPI 3.1 YAML Specification**: Available at `/openapi.yaml`

There is no need to manually run any code generation commands when updating API routes or models.
