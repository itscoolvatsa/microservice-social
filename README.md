The app is where we have the main.go to bootstrap all the applications (basically where we import the incoming and
outgoing packages and some from the internal if needed).

The cmd is where we define the command line tools we need for the service. In this case we have a cli that interacts
with our service.

The docs is where we put the relevant documentation for this project.

As I mentioned before the endpoints are exposed via grpc so that folder is responsible for implementing the interfaces
provided by the protobuffer compilation.

The internal is where we have the implementation of the interfaces we use in the project but are not part of the domain;
basically the infrastructure things like the database, queues, external services (Slack), the logger…

The k8s one is referred to the Kubernetes config files for the deployment.

As mentioned before the incoming and outgoing are our domain packages so that is where all the magic happens. I will
describe a bit of one of them because the other is similar but with the particularities of that domain.

auth
├───app
├───cmd
├───docs
├───internal
│ ├───controller
│ ├───handler
│ └───repository
├───k8s
└───pkg
└───model

docker-compose.test.yml
go.mod
go.sum
Makefile
README.md

// To build the docker file from parent directory

> docker build -t auth -f ./auth/Dockerfile .

// TO create a new token in kubernetes dashboard

> kubectl create token default

// To forward port for localhost access

> kubectl port-forward auth-depl-7c4589b9c7-qfzhc 8080:8080

// To create a new deployment from yaml

> kubectl apply -f <path/\*.yaml>

// Ubuntu

> eval $(minikube docker-env)

// Minikube docker configuration for internal cache
PowerShell

> & minikube -p minikube docker-env --shell powershell | Invoke-Expression

cmd

> @FOR /f "tokens=\*" %i IN ('minikube -p minikube docker-env --shell cmd') DO @%i

`
{
    "status": "success",
    "data": {
        "id": 1,
        "name": "John Smith",
        "email": "john.smith@example.com",
        "age": 35,
        "address": {
            "street": "123 Main St",
            "city": "Anytown",
            "state": "CA",
            "zip": "12345"
        },
        "phoneNumbers": [
            {
                "type": "home",
                "number": "555-555-1234"
            },
            {
                "type": "work",
                "number": "555-555-5678"
            }
        ],
        "active": true
    },
    "message": "User retrieved successfully"
}

`
