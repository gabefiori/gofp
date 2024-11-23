# gofp
A simple tool for quickly finding projects.

## Installation
```sh
go install github.com/gabefiori/gofp/cmd/gofp@latest
```

## Configuration
Create a configuration file at `~/.config/gofp/config.json`:

```json
{
  "sources": [
    {
      "path": "~/your/path",
      "depth": 1
    },
    {
      "path": "/home/you/your_other/path",
      "depth": 3
    }
  ]
}
```

## Integrate with your shell
The shell integration is necessary because, due to POSIX standards, a program cannot change the current working directory of the shell that invoked it.
When you run a command in a shell, it operates in its own process, and any changes to the working directory made by that process do not affect the parent shell.

### bash/zsh
To set up the integration in bash or zsh, add the following function to your `.bashrc` or `.zshrc` file:

```sh
function fp() {
    cd "$(gofp)"
}
```

### fish
For users of the Fish shell, you can achieve the same functionality by adding the following function to your `config.fish` file:

```fish
function fp
    cd (gofp)
end
```
