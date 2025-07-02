# Spacelift Resources Scraper

A Go-based tool that scrapes and converts Spacelift Terraform provider documentation into markdown files.

## Overview

This tool automatically fetches documentation for Spacelift Terraform provider resources from the Terraform Registry and converts them to markdown format. It handles concurrent requests with retry logic and failure tracking.

## Features

- Concurrent documentation fetching (8 workers by default)
- Automatic retry on failures (3 attempts)
- HTML to Markdown conversion
- Failure tracking and reporting
- Resource list management

## Requirements

- Go 1.x
- Required dependencies:
  - github.com/JohannesKaufmann/html-to-markdown
  - github.com/PuerkitoBio/goquery

## Usage

1. Ensure your resource list is in `resources.txt` (one resource per line)
2. Run the program:
   ```sh
   go run main.go
   ```
3. Generated markdown files will be saved in the `docs/` directory

## Configuration

The following constants can be adjusted in `main.go`:

- `maxWorkers`: Number of concurrent workers (default: 8)
- `maxRetries`: Number of retry attempts (default: 3)
- `outDir`: Output directory for markdown files (default: "docs")

## Output

- Successfully converted files are saved to the `docs/` directory
- Failed conversions are logged in `failure.txt`
- Progress and errors are logged to stdout

## Project Structure

```
.
├── main.go           # Main application code
├── resources.txt     # List of Spacelift resources to process
├── go.mod           # Go module file
└── docs/            # Generated markdown files (created on run)
```
