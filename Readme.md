# UML-Caddy

Uml-caddy is a tool for templating plantuml from known metadata

# TODO 

- Ablity to upload custom kubectl config
- AWS S3 mapping/compatibility
- Terraform State file compatibility
- .tf file compatibility
- k8s interface tests

## Running

The docker-compose file allows you to run standard docker commands to run the container locally for debugging

`docker-compose up --build`

## Process Layers

### Front


- API scanner interface
    - k8s
    - terraform
    - AWS
- Simple frontend webapp

### Middle

- Input Template model
- Input Template
- Package for common defaults
    - K8s (Models may be adoptable from k8s models)
        - Pod
        - Node
        - Deployment/ReplicaSet
        - PV/PVC
        - Secrets/Configmaps
    - Terraform
        - VPC/Subnets
        - EC2
        - S3
        - IAM (as data objects)
        - Databases
    - AWS
        - s3
        - ec2
        - VPC
        - Subnets
        - Route53
        - LBs
        - EKS
        - ECR
        - ECS
        - Aurora
        - Dynamodb
- Template Model processing
    - Gathering data from known sources

### Backend

- Template processing and translation to plantuml
- PlantUML generation/export

## Methodology

Input Infrastructure resources, output plantuml

Infrastructure resources should have a defined template utilizing 
    - [golang templating syntax](https://pkg.go.dev/text/template)
    - [sprig](https://github.com/Masterminds/sprig)

Other tooling
    - [k8s client](https://github.com/kubernetes/client-go)
    - [chi router](https://github.com/go-chi/chi)
