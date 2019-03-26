go get https://github.com/jackdoe/go-baxx-cli

# baxx-put: upload stdin to baxx.dev

encrypt os.Stdin and upload it to baxx.dev

## Install

go install https://github.com/jackdoe/go-baxx-cli/cmd/baxx-put

## Usage
Usage baxx-put:
  -k string
        file containing the encryption key, key is sha256(content of the file)
Example:
  cat file | gzip | ./baxx-put -k pass-file https://baxx.dev/io/$BAXX_TOKEN/file




# baxx-get: download files from baxx.dev
download specific file from baxx.dev, decrypt it and print it
to os.Stdout

## Install

go install https://github.com/jackdoe/go-baxx-cli/cmd/baxx-get

## Usage

Usage baxx-get:
  -k string
        file containing the encryption key, key is sha256(content of the file)
Example:
  baxx-get -k pass-file https://baxx.dev/io/$BAXX_TOKEN/file > file.downloaded

