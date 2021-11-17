# QR Code Generator

<aside>
💡 María Camila Lenis Restrepo A00351598

</aside>

# How to run the app?

## Prerrequisites

- Docker installed and running
- Docker compose installed

## Let's run this

1. Clone this repository on your machine
2. Go to the reverse-proxy folder and run `bash generate_keys.sh` . This will generate the pair of keys for the ssl certificate 
3. Get back to the root of the repository and build the app running `docker-compose build`
4. Then upload the app and let this running with this command `docker-compose up`

## Let's play with it

Once the app is running you can access to it with [`https://localhost/qrcodes`](https://localhost/qrcodes)

The purpose of the API is to convert a string into a Base64 QR code that you can paste in a [Base64 to Image Converter](https://codebeautify.org/base64-to-image-converter) to scan it. 

This are the Endpoints that you can use to play 

### GET qrcodes → List all the QR codes in the system (limit 20)

```json
[
  {
    "id": "1",
    "text_value": "saved string ",
    "encoded_qr": "Base 64 QR Code"
  }
]
```

### GET qrcodes/{id} →Gets a specific QR Code information

```json
{
    "id": "1",
    "text_value": "saved string ",
    "encoded_qr": "Base 64 QR Code"
}
```

### POST qrcodes → Add a QR code to the system

Send the body with a JSON like this

```json
{
	"text_value": "The text that I want to convert to QR Code"
}
```

It must returns this

```json
{
  "QR ID": "QRID",
  "Status": "Created"
}
```

### PUT qrcodes/{id} → Updates the text_value and generate the Base 64 code

```json
{
  "ID": "QRID",
  "Status": "Updated"
}
```

### DELETE qrcodes/{id} → Deletes an specific QR code

```json
{
  "ID": "QRID",
  "Status": "Deleted"
}
```

### GET qrcodes/health → View the general state of the app

```json
{
  "BD Connection": "Running",
  "QR Codes Saved": "1"
}
```

# How the app is build?

This is an app to generate and manage QR codes from a text. The API is build in Golang and it's connected to a MySQL database. With the run of the Docker Compose the application will upload the following architecture:

![docker (1).png](QR%20Code%20Generator%20e621196341ab40a79d26f6686320f2ea/docker_(1).png)

Where you can see there are three Docker containers: One that works as a reverse proxy, the API container and the database. 

## Docker Images

[For the reverse proxy](https://hub.docker.com/r/camilaleniss/proxy)

[For the API](https://hub.docker.com/r/camilaleniss/web)

[For the Database](https://hub.docker.com/r/camilaleniss/mysql)

You can see the Dockerfile inside each folder in this repo. 

# Next steps

## Production deploy

- To run this service in a production environment, we must deploy the containers in a Cloud service. We could use a service like AWS ECS to run Docker containers.
- To be able to connect with the service, the container nginx container must have a public IP.  To access to it easier, we could use a DNS service to give it a domain name.

## Improvements

- Include a bash script that makes the build and deploy of the app, including functional tests
- Makefile with those phases
- Refactor of the code app
- Include replicas with Docker swarm to make it more reliable
- Environment variables