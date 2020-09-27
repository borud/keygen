# keygen

This is just a simple random key generator for producing keys suitable
for use with AES or any other symmetric cipher.  The default key size
is 256 bits.

The random generator used is the `crypto/rand` package from Go, so if
you want to evaluate whether this is sufficiently high quality for
your requirements you should investigate this package.

As a convenience the key is output as base64 (RFC 4648), hex encoded
and as a C style array for your cutting and pasting pleasure.

## Installing

    go get -u github.com/borud/keygen/cmd/keygen
