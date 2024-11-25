# gofp
A simple tool for quickly finding projects.

<img alt="Demo" src="examples/demo.gif" width="600" />

## Installation
```sh
go install github.com/gabefiori/gofp/cmd/gofp@latest
```

Once the installation is complete, you can use the `gofp` command along with other commands in your shell.

### Examples with `cd`:

<details>
<summary>Bash</summary>

> Add to your `~/.bashrc` file:
>
> ```sh
> alias fp='cd "$(gofp)"'
> ```

</details>

<details>
<summary>Zsh</summary>

> Add to your `~/.zshrc` file:
>
> ```sh
> alias fp='cd "$(gofp)"'
> ```

</details>

<details>
<summary>Fish</summary>

> Add to your `~/config.fish` file:
>
> ```fish
> alias fp "cd (gofp)"
> ```

</details>

### Using with tmux
You can utilize this [script](/scripts/gofp-tmux.sh), which enables you to easily attach to or switch between Tmux sessions using the `gofp` command for selection.

<details>
<summary>Install</summary>

>```sh
>sudo wget -O /usr/local/bin/tms https://raw.githubusercontent.com/gabefiori/gofp/refs/heads/main/scripts/gofp-tmux.sh
>sudo chmod +x /usr/local/bin/tms
>```

</details>

## Configuration
Create a configuration file at `~/.config/gofp/config.json`:

```json
{
  "expand_output": true,
  "selector": "fzf",
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

> - `"expand_output"` is optional and defaults to `true`. 
> - `"selector"` is optional and defaults to `fzf`. Available options are `fzf` and `fzy`.

## CLI Options
```sh
--config file, -c file      Load configuration from file (default: "~/.config/gofp/config.json")
--selector value, -s value  Selector for displaying the projects (available: "fzf", "fzy") (default: "fzf")
--expand-output, --eo       Expand output (default: true)
--measure, -m               Measure performance (time taken and number of items) (default: false)
--help, -h                  show help
```
