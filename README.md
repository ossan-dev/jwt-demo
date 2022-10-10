# jwt-demo
A simple demo about JWT token with Go.

## go jws - PEM
- `openssl req -newkey rsa:4096  -x509  -sha512  -days 365 -nodes -out certificate.pem -keyout privatekey.pem` => this generates the certificate together with the private key
- `openssl x509 -noout -in certificate.pem -text` => used to see in a text-form the content of the certificate
- `openssl x509 -pubkey -noout -in certificate.pem -out publicKey.pem` => to extract the public key from the certificate

## resources:
- jwt vs jws vs jwe vs jwk: https://medium.com/@goynikhil/what-is-jwt-jws-jwe-and-jwk-when-we-should-use-which-token-in-our-business-applications-74ae91f7c96b
- generate self-signed certificate: https://linuxconfig.org/how-to-generate-a-self-signed-ssl-certificate-on-linux

## go jws - RS 256 private key and public pair
- `openssl genrsa -out certs/id_rsa 4096` => this generate a private key
- `openssl rsa -in certs/id_rsa -pubout -out certs/id_rsa.pub` => generate public key based on the private key