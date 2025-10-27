# CipherTracey

**Local-only, open-source AML intelligence and analytics for cryptocurrency flows.**

## Abstract

CipherTracey is a command-line tool for analyzing blockchain transactions and identifying suspicious patterns in cryptocurrency flows. It runs entirely locally, visualizing transaction graphs in your browser without sending data to third parties.

## Motivation

Anti-Money Laundering (AML) analysis for cryptocurrency requires powerful tools that respect privacy. CipherTracey provides investigators, compliance teams, and researchers with a self-hosted solution for:

- Ingesting blockchain transaction data
- Scoring addresses based on risk heuristics
- Visualizing complex transaction flows
- Exporting analysis reports

All processing happens on your machine, ensuring complete data sovereignty.

## Quick Start

### Installation
```bash
# Clone the repository
git clone https://github.com/Oblivionsage/ciphertracey.git
cd ciphertracey

# Build the binary
make build

# Run it
./bin/ciphertracey --version
```

### Basic Usage
```bash
# Show help
./bin/ciphertracey help

# Ingest transaction data (stub)
./bin/ciphertracey ingest --source bitcoin.json

# Calculate risk scores (stub)
./bin/ciphertracey score --address 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

# Visualize transaction graph (stub)
./bin/ciphertracey graph --address 1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa

# Export analysis results (stub)
./bin/ciphertracey export --format json

# Start interactive server (stub)
./bin/ciphertracey run --port 8080
```

## Development
```bash
# Run tests
make test

# Format code
make fmt

# Vet code
make vet

# Clean build artifacts
make clean
```

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.

## Status

**Early Development** - Core functionality is being implemented.
