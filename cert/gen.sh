#!/bin/bash

rm *.pem

# Prompt for information
read -p "Country (2 letter code) [AU]: " country
read -p "State or Province Name (full name) [New South Wales]: " state
read -p "Locality Name (eg, city) [Sydney]: " locality
read -p "Organization Name (eg, company) [Internet Widgits Pty Ltd]: " organization
read -p "Organizational Unit Name [*.google.com]: " organizationalunit
read -p "Common Name (e.g. server FQDN or YOUR name) [Gary Sanford]: " commonname
read -p "Email Address [garysanford@amail.com]: " email

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=$country/ST=$state/L=$locality/O=$organization/OU=$organizationalunit/CN=$commonname/emailAddress=$email"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -new -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=$country/ST=$state/L=$locality/O=$organization/OU=$organizationalunit/CN=$commonname/emailAddress=$email"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate   # -extfile server-ext.cnf
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text
