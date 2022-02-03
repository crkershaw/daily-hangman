# tufferina

# Running normally



# Running on docker

```
docker build -t tufferina .

docker run --rm -p 80:80 tufferina
```

# AWS

## Internet Gateway
tufferina-ig # Attached to tufferina VPC

## VPC
Tufferina

# Route tables
tufferina-rt-public (routes to 10.1.0.0/16 - local, and 0.0.0.0/0 - internet gateway)


# Subnets (tied to VPC and route table)
tufferina-subnet-public-2b (linked to tufferina-rt-public)
Xtufferina-subnet-public-2c (linked to tufferina-rt-public)
Xtufferina-subnet-private-2c (routes to 10.1.0.0/16 - local, and 0.0.0.0/0 - nat-pub1)

# NAT gateways (tied to subnets)
Xtufferina-nat-pub1 (tied to tufferina-subnet-public-2c)
Xtufferina-nat-pub2 (tied to tufferina-subnet-public-2b)
tufferina-nat-pub (tied to tufferina-subnet-public-2b)

## Security groups
tufferina-sg-eb
tufferina-sg-lb

# Load balancer
Xtufferina-lb (mapped to subnet-public-2b and subnet-public-2c, security group tufferina-sg-lb, listeners http 80 and forwards to tufferina-tg-http)

# Target group
tufferina-tg-http (http port 80, associated with tufferina-lb)

# Instance
Linux
Tufferina VPC
subnet-public-2b
Default vpc security group

# Deploying to AWS

1. Ensure AWS CLI is installed
2. [Link docker to Amazon ECR repository](https://docs.aws.amazon.com/AmazonECR/latest/userguide/docker-push-ecr-image.html)
```
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 919994557194.dkr.ecr.us-east-2.amazonaws.com/tufferina
```

```
docker images
docker tag tufferina 919994557194.dkr.ecr.us-east-2.amazonaws.com/tufferina
docker push 919994557194.dkr.ecr.us-east-2.amazonaws.com/tufferina
```







# Previous docker notes

```
tufferina>docker run -it --rm -p 8010:8010 -v $PWD/src:/go/src/tufferina tufferina
```

The docker run command is used to run a container from an image,
The -it flag starts the container in an interactive mode (tie it to the current shell),
The --rm flag cleans out the container after it shuts down,
The --name mathapp-instance names the container mathapp-instance,
The -p 8010:8010 flag allows the container to be accessed at port 8010,
The -v $PWD/src:/go/src/mathapp is more involved. It maps the src/ directory from the machine to /go/src/mathapp in the container. This makes the development files available inside and outside the container, and
The mathapp part specifies the image name to use in the container.

