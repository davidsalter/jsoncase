# JSON Case Converter

This application converts the keys of a JSON file to uppercase or lowercase based on the provided flag. With no flag, the default is to convert to lower case.

## Usage

```sh
jsoncase [flags] <input.json>
```

For example:

```sh
jsoncase example/input.json
```

or

```sh
jsoncase -upper example/input.json
```

### Flags

* -h show help
* -upper convert to uppercase

## Building

```sh
make build
```

## Running

```sh
make run
```

## Testing

```sh
make test
```

## Cleaning

```sh
make clean
```