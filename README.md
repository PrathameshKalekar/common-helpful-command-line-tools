

# PraLex CLI Tool

`PraLex` is a versatile command-line tool written in Go using the Cobra library. It offers several useful commands for managing files and directories, and integrates seamlessly with Visual Studio Code (VS Code).

## Commands

### `pralex make [name]`

Creates a file or directory based on the provided name.

- If no file extension is provided, it creates a directory.
- If a file extension is provided, it creates a file.

**Examples:**

```sh
pralex make my_directory
pralex make my_file.txt
```

### `pralex remove [name]`

Removes the specified file or directory.

**Examples:**

```sh
pralex remove my_directory
pralex remove my_file.txt
```

### `pralex open`

Opens the current directory in VS Code.

**Examples:**

```sh
pralex open
```

### `pralex open -u [file.zip]`

Unzips the specified file and opens the unzipped folder in VS Code.

**Examples:**

```sh
pralex open -u my_folder.zip
```

### `pralex rename [old_name] [new_name]`

Renames the specified file or directory to the new name.

**Examples:**

```sh
pralex rename old_file.txt new_file.txt
pralex rename old_directory new_directory
```

## Installation

To install `Pralex`, ensure you have Go installed and then use the following commands:

1. Clone the repository:

    ```sh
    git clone https://github.com/PrathameshKalekar/common-helpful-command-line-tools.git
    ```

2. Navigate to the project directory:

    ```sh
    cd common-helpful-command-line-tools
    ```

3. Build and install:

    ```sh
    go install
    ```

After installation, you can use the `pralex` command from anywhere in your terminal.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests on the [GitHub repository](https://github.com/PrathameshKalekar/common-helpful-command-line-tools).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please contact [Email](kalekarprathamesh19@gmail.com) or open an issue on the [GitHub repository](https://github.com/PrathameshKalekar/common-helpful-command-line-tools).