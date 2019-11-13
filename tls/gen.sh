#!/bin/sh

# Create Root signing Key
openssl genrsa -out ca.key 2048

# Generate self-signed Root certificate
openssl req -new -x509 -key ca.key -sha256 -subj "/C=JP/ST=TKO/O=CA, Inc." -days 3650 -out ca.crt

# ---

# Create a Key certificate for the Server
openssl genrsa -out service.key 4096

# Create a signing CSR
openssl req -new -key service.key -out service.csr

############################################################################
# You are about to be asked to enter information that will be incorporated
# into your certificate request.
# What you are about to enter is what is called a Distinguished Name or a DN.
# There are quite a few fields but you can leave some blank
# For some fields there will be a default value,
# If you enter '.', the field will be left blank.
# -----
# Country Name (2 letter code) []:JP
# State or Province Name (full name) []:TKO
# Locality Name (eg, city) []:Minato
# Organization Name (eg, company) []:Test, Inc.
# Organizational Unit Name (eg, section) []:
# Common Name (eg, fully qualified host name) []:localhost
# Email Address []:
#
# Please enter the following 'extra' attributes
# to be sent with your certificate request
# A challenge password []:
############################################################################

# Generate a certificate for the Server
openssl x509 -req -in service.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out service.pem -days 3650 -sha256
