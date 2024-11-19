# ai-commit-gen

An AI-powered commit message generator using OpenAI or Ollama models.

## Features

- Generates meaningful commit messages using AI
- Supports both OpenAI and local Ollama models
- Secure configuration management
- Interactive spinner during generation
- Easy installation via Homebrew

## Installation

```bash
brew install aryan1306/tap/ai-commit-gen
```

## Configuration

Before running the program, you need to set up the OpenAI API key in the config file in the $HOME directory eg: `~/.config/commit-gen/config.json`.

The structure of the config file should be something like this:

```json
{
  "openai_key": "your-api-key",
  "ollama_server": "http://localhost:11434/api/chat",
  "default_model": "deepseek-coder:latest"
}
```

## Usage

```bash
# Generate commit message
commit-gen

# Show version
commit-gen --version
```

## Prerequisites

For Ollama models:
```bash
brew install ollama
ollama pull deepseek-coder
ollama serve
```

## License

MIT License - see LICENSE file for details