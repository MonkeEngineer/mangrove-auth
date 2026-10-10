# Mangrove Auth

Self-hosted authentication server built with Go and used as identity layer for the Open Admin ecosystem.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install): 1.26 or higher

### Setup instructions

1. Clone the repository:

   ```bash
   git clone https://github.com/ecosystem-admin/mangrove-auth.git
   cd mangrove-auth
   ```

2. Run the server:

   ```bash
   go run ./cmd/mangrove-auth/
   ```

   The server listens on port `8080` by default. To use a different port, set `MANGROVE_PORT`:

   ```bash
   MANGROVE_PORT=9090 go run ./cmd/mangrove-auth/
   ```

3. Verify it is running:

   ```bash
   curl http://localhost:8080/healthz
   ```

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) to learn how to get involved.
