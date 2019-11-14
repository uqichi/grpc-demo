#!/bin/sh

# https://itnext.io/practical-guide-to-securing-grpc-connections-with-go-and-tls-part-1-f63058e9d6d1

# ---
# オレオレ認証局

# Create Root signing Key
openssl genrsa -out ca.key 2048

# Generate self-signed Root certificate (証明書署名要求の作成と証明書の作成をまとめてやる)
openssl req -new -x509 -key ca.key -sha256 -subj "/C=JP/ST=Tokyo/O=CA, Inc./CN=rootca" -days 3650 -out ca.crt

# ---
# オレオレ証明書

# Create a Key certificate for the Server
openssl genrsa -out service.key 2048

# Create a signing CSR
openssl req -new -key service.key  -subj "/C=JP/ST=Tokyo/O=Test, Inc./CN=test.jp" -out service.csr

# Generate a certificate for the Server
openssl x509 -req -in service.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out service.pem -days 3650 -sha256

# Look inside the cert
openssl x509 -text -noout -in service.pem