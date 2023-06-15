<h1 align="center">Automated-AWS</h1>
<br />


  <p align="center">  
    <img src="https://img.shields.io/github/license/AbhiAlest/Automated-AWS.svg" alt = "License" >
    <img src="https://img.shields.io/github/issues/AbhiAlest/Automated-AWS.svg" alt = "Issues" >
    <img src="https://img.shields.io/github/issues-pr/AbhiAlest/Automated-AWS.svg" alt = "Pull Requests" >
    <img src="https://img.shields.io/github/watchers/AbhiAlest/Automated-AWS.svg" alt = "Watchers" >
    <img src="https://img.shields.io/github/stars/AbhiAlest/Automated-AWS.svg" alt = "Stars" >
  </p>
  

   Automated-AWS is an open-source [Go](https://go.dev/) package that automates provisioning/management of AWS resources (ec2, elbv2, and autoscaling).


<br />  

<h2>Setting Up The Development Environment</h2>
To set up the development environment for the project, follow the steps below:

<h3>Prerequisites</h3>
Make sure you have the following prerequisites installed on your system:

* Go (version 1.16 or later)
* Docker (optional, if you want to run the application in a container)
* AWS CLI (to interact with AWS services)

<h3>Clone the Repository</h3>
Clone the project repository from the GitHub repository:

```bash
git clone https://github.com/AbhiAlest/Automated-AWS.git
```

<h3>Install Dependencies</h3>

Change to the project directory:

```bash
cd Automated-AWS
```

Install the Go dependencies:

```go
go mod download
```

<h3>Configure AWS Credentials (if applicable)</h3>
If your code interacts with AWS services and you want to run it locally, configure the AWS CLI with your credentials:

```
aws configure
```

<h3>Build and Run</h3>
To build and run the project, use the following command:

```go
go run main.go
```
If you want to run the application in a Docker container, build the Docker image first:

```bash
docker build -t myapp .
```

Then, run the Docker container:
```bash
docker run -p 8080:8080 myapp
```

<h3>Running Tests</h3>
To run the tests, use the following command:

```bash
go test ./...
```

<h3>Running via Command Line Interface (CLI)</h3>
Use the command line for desired actions. Here are a few examples to get started:

```bash
# Create VPC
go run main.go -create-vpc

# Create S3 bucket
go run main.go -create-s3-bucket

# Launch 2 EC2 instances
go run main.go -launch-ec2-instances=2

# Create ALB
go run main.go -create-alb
```

<h3>Development Workflow</h3>

1. Make the necessary code changes in the project files.
2. Run the application or tests to verify your changes.
3. Commit the changes to your local Git repository.
4. Push the changes to the remote Git repository (if applicable).

<h3>Additional Configuration</h3>
Additional configuration for this project may apply. If you need to configure any project-specific settings, refer to the config package or the relevant configuration files in the project. Please note that the above instructions are a general guideline and may need to be adjusted based on your specific project requirements.

<br/><br/>
That's it! You now have your development environment set up for the project. Happy coding!

<br />  

<h2>File Structure</h2>
  
  Rough overview of the file structure for this project. Please keep in mind that this file structure is not frequently updated (i.e changes have occurred). 
  ```
Automated-AWS
├── .github
│   └── CODEOWNERS
│   └── CONTRIBUTING.md
├── cmd
│   └── main.go
├── config
│   └── aws.go
│   └── config.yaml
├── mocks
|   ├── elbv2_mock.go
├── pkg
│   ├── autoscaling
│   │   └── autoscaling.go
│   ├── ec2
│   │   └── ec2.go
│   └── elbv2
│       └── elbv2.go
├── test
│   ├── autoscaling
│   │   └── autoscaling_test.go
│   ├── ec2
│   │   └── ec2_test.go
│   └── elbv2
│       └── elbv2_test.go
├── Dockerfile
├── README.md
└── LICENSE

  ```
   
<br />
