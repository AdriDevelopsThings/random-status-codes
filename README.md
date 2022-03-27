# random-status-codes
A webserver written in golang that returns random status codes without body or needed headers.

## Use it
### Docker
``docker run -p 80:80 ghcr.io/adridevelopsthings/random-status-codes:main``

### Raw

``go build . -t main``


You can configure a host in the environment variable ``HTTP_HOST`` (for example ``0.0.0.0:80``).


Start it with: `./main`.