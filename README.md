# Eventurely Backend

Eventurely is an application that allows users to create and manage events. It is created using Go, Connect RPC, and Flutter. Created as a learning exercise whilst (loosely) following the book "Let's Go" by Alex Edwards.

## Dependencies

- Go v1.21
- Air v1.49 (optional, for hot reloading)
- Buf v1.29 (optional, for generating Protobuf files)

## Notes
- This project is a work in progress and is not yet ready for production use.
- The project leverages the `buf` CLI for generating the Protobuf files for the gRPC server. This can be done using `buf generate api/proto` from the root of the project. This will place the generated files in the `gen` directory.