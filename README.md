# ai-commit-gen

An AI-powered commit message generator using OpenAI or Ollama models.

## Features

- Generates meaningful commit messages using AI
- Supports both OpenAI and local Ollama models
- Secure configuration management
- Interactive spinner during generation
- Easy installation via Homebrew

## Installation

Using Go you can install this using the `go get` command
```bash
go get github.com/aryan1306/ai-commit-gen
```
Or, using Homebrew on macOS/Linux devices
```bash
brew install ai-commit-gen
```

## Configuration

Before running the program, you need to set up the OpenAI API key in the config file in the $HOME directory eg: `~/.config/commit-gen/config.json`.

The structure of the config file should be something like this:

```js
{
  "openai_key": "your-api-key", // leave it blank if you are not using it
  "ollama_server": "http://localhost:11434/api/chat",
  "default_model": "qwen2.5-coder:3b"
}
```

## Usage

```bash
# Generate commit message
ai-commit-gen

# Show version
ai-commit-gen --version

# Select different model other than one mentioned in config
ai-commit-gen -model=gemma2:latest
```

## Prerequisites

For Ollama models:
```bash
brew install ollama
ollama pull qwen2.5-coder:3b # or any other Ollama model you like
ollama serve
```

## License

MIT License - see LICENSE file for details
