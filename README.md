# JinaAI Scraper

This repository contains a Go-based scraper tool for extracting content from URLs or sitemaps and returning it in various formats such as HTML, MDX, or plain text. The tool leverages the JinaAI reader API to perform content extraction.

## Features

- Scrape individual URLs or parse sitemaps to extract URLs.
- Support for multiple return formats: HTML, MDX, and plain text.
- Concurrent URL scraping (TODO).

## Prerequisites

- Go 1.16 or later
[JinaAI reader API Key](https://jina.ai/reader/)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/jinaai-scraper.git
    cd jinaai-scraper
    ```

2. Build the executable:

    ```sh
    go build -o jinaai-scraper
    ```

## Usage

### Command-line Flags

- `-url`: URL to scrape.
- `-sitemap`: URL of the sitemap.
- `-return`: Return type (e.g., `html`, `mdx`, `text`).
- `-apiKey`: JinaAI reader API key.

### Examples

#### Scrape a Single URL

To scrape a single URL and return the content as HTML:

```sh
./jinaai-scraper -url https://example.com -return html
```

To scrape a single URL and return the content as MDX:

```sh
./jinaai-scraper -url https://example.com -return mdx -apiKey your-api-key
```

#### Parse a Sitemap

To parse a sitemap and validate its format:

```sh
./jinaai-scraper -sitemap https://example.com/sitemap.xml
```

### Flag Details

- `apiKey`: The API key for JinaAI's reader service. Required for MDX and text return types.
- `sitemap`: The URL of the sitemap to parse. The scraper will validate if the provided URL is a valid sitemap.
- `return`: Specifies the format in which the scraped content should be returned. Valid values are `html`, `mdx`, and `text`.
- `url`: The URL to scrape. If provided, the scraper will fetch and return the content of this URL.

## Code Structure

- `main.go`: Contains the main function and command-line flag parsing.
- `scrapUrl`: Function to scrape content from a single URL.
- `isSiteMap`: Function to validate if a URL is a sitemap.
- `scrapSitemapUrls`: (TODO) Function to scrape URLs found within a sitemap concurrently.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes or enhancements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.