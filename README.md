# notes

An small API that stores and returns notes from cache. It will work with the JSON representation of the notes.

A note has two fields: message, and an optional tag. Here are two examples:
```json
{ “message”: “Buy apples and oranges”, “tag”: “groceries” }
```
```json
{ “message”: “Feed the cat”}
```

- The tag field is optional
- The tag may be used to retrieve all notes with this tag
- A note may not have multiple tags
- The message field is mandatory and cannot be empty

## API

The api has only one endpoint which implements several methods:

- __GET__ /notes
- __POST__ /notes
- __DELETE__ /notes

> For more details, see the postman collection

## Usage

```bash
# go version go1.18.5 linux/amd64

# build and start the server with the default port (3000)
go run main.go

# build and start the server with a custom port using environment
NOTE_PORT=4000 go run main.go
```

## Dependencies

This project is a simple vanilla golang api, without any dependency because "Scaling and performance are not a concern", so *less is more*.
