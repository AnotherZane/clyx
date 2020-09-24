# Clyx.ml

#### Simple file uploader implementation in different programming languages.

### Purpose of this project?

This project was started by [@pepsi](https://github.com/pepsi) in 2019. It's purpose was to help learn some basics of programming languages by making a simple file upload web server. I tend to do the same.

### Some rules:

1. All implementations should have the following routes:
    - ``/`` - Serve static files from ``public/`` folder
    - ``/upload`` - Upload endpoint
    - ``/i/:file`` - Serve uploaded files

2. The response from ``/upload`` should be in json with the following:
    - ``success`` - Boolean value
    - ``url`` - File url
    - ``error`` - Any errors

3. Must have a check for the key from ``config.json`` on upload.

4. Functional upload button on homepage. (Optional)

### List of current implementations:

- [Go](https://github.com/AnotherZane/clyx-ml/tree/go) (Old)


### List of planned implementations:

- Node
- Rust
- Go (New)
- Elixir
- Python
- C#
- Java
- Php?