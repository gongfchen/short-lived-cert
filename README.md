# short-lived-cert

## Setup

```
$ pip3 install merklelib
$ cd CA/script
$ go mod download
```

## Build

```
$ cd <CA or middle-daemon or website-daemon>/script
$ go build
```

## Run

### CA

Starting running CA

```
$ ./CA
```

### Middle daemon

Start running Midle Daemon

```
./Middle-daemon
```

To request short-lived certificates for a domain (default to be 365 certificates). Need domain public keys available.

```
request [domain_name]
```

To request daily decryption key for a domain. Need to provide the day number.

```
key [domain_name] [day_number]
```

### Website daemon: 

To verify a certificate, run the following:

```
./Website-daemon
...
Listening to request...
[domain_name] [num_of_day]
```
