<h1 align="center">VIMc</h1>

VIMc is a very basic VimSript "interpreter" for running `.vim` files in the terminal. It is essentially just a wrapper around Vim that auto runs your `.vim` script in a headless mode then quits

<sub>What, did you expect me to write something which didn't need to be written? If something more than adequate can be written as a simple wrapper, why do any more?</sub>

## Usage
- help
    ```sh
    vimc -h
    ```
- run file

    (replace `/foo/bar/baz.vim` with your VimScript file)
    ```sh
    vimc /foo/bar/baz.vim
    ```

## Dependencies
- Vim

## Installation
- Download binary

    wcurl:
    ```sh
    wcurl https://raw.githubusercontent.com/Supraboy981322/vimc/master/build/vimc
    ```

- Move binary to path (eg: `/usr/bin`)
    
    (may require `su`)
    ```sh
    mv vimc /usr/bin/vimc
    ```
