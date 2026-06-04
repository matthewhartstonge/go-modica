# go-modica #
[![build](https://github.com/matthewhartstonge/go-modica/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/matthewhartstonge/go-modica/actions/workflows/build.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewhartstonge/go-modica)](https://goreportcard.com/report/github.com/matthewhartstonge/go-modica)

`go-modica` is a Go Client library for accessing [Modicagroup's RESTful APIs.][modica api uri]

Requires Go version 1.20 or greater.

[modica api uri]: https://confluence.modicagroup.com/display/DC/Modica+API+Documentation

## Usage

```go
import "github.com/matthewhartstonge/go-modica"
```

### Authentication

The `go-modica` library handles client basic authentication for you and asks for
these when creating a new api client. You can find your client credentials under:

[Omni Dashboard][omnidashboard] > Applications >  Mobile Gateway (API)

* `ClientID`: is the `Application Name`
* `ClientSecret`: is the configured `password`

```go
client := modica.NewClient("ClientID", "ClientSecret", nil)
```

[omnidashboard]: https://omni.modicagroup.com

### GetMessage

Get a message by ID.

```go
// Get a message that has been sent
client := modica.NewClient("clientID", "clientSecret", nil)
message, err := client.MobileGateway.GetMessage(123456789)
if err != nil {
    panic(err)
}

fmt.Printf("message: %#+v\n", message)
```

### CreateMessage

Send a single SMS message with create message.

```go
newMessage := &modica.Message{
    Destination: "+642123456789",
    Content:     "Hello, this is a test message!",
}

client := modica.NewClient("clientID", "clientSecret", nil)
msgID, err := client.MobileGateway.CreateMessage(newMessage)
if err != nil {
    panic(err)
}

fmt.Printf("message: %#+v\n", msgID)
```

### CreateBroadcastMessage

Send a broadcast SMS message to multiple destinations:

```go
newBroadcastMessage := &modica.BroadcastMessage{
    Destinations: []string{"+642123456789", "+64987654321"},
    Message: modica.Message{
        Content: "Hello, this is a test message!",
    },
}

client := modica.NewClient("clientID", "clientSecret", nil)
messageStatuses, err := client.MobileGateway.CreateBroadcastMessage(newBroadcastMessage)
if err != nil {
    panic(err)
}

for _, status := range messageStatuses {
    switch status.Status {
    case modica.MessageStatusSubmitted:
        fmt.Printf("Message %d has been %s\n", status.ID, status.Status)
    case modica.MessageStatusRejected, modica.MessageStatusFailed, modica.MessageStatusDead, modica.MessageStatusExpired:
        fmt.Printf("Message %d has %s\n", status.ID, status.Status)
    }
}
```

## Versioning
In general, go-modica follows [semver](https://semver.org/) as closely as we can 
for tagging releases of the package.

* The major version is incremented with any incompatible change to the exported 
	Go API.
* The minor version is incremented with any backwards-compatible changes to 
	functionality, including new features.
* The patch version is incremented with any backwards-compatible bug fixes.

## License

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
