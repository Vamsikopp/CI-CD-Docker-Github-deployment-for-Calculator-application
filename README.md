# CI-CD-Docker-Github-deployment-for-Calculator-application



/****Containerize the source code using a Dockerfile****/



1.Install docker application

To containerize code using docker file:

Create a new file named “Dockerfile” in the same directory as your application file.

In Dockerfile, write a base image. If we want to build a Docker image for a  code. 







COPY
 WORKDIR /app
Copy the requirement file of your application from our local machine into the working directory.


COPY

COPY
 COPY requirements.txt .
Install the dependencies which are present in the requirement file.




COPY
 RUN pip install --no-cache-dir -r requirement.txt
In requirement.txt, we only need the Flask for this application.

Copy our application from our local/host machine into the working directory.


COPY

COPY
 COPY 
Define CMD as the startup command. Specify the command that should be executed when a container is launched from the image.








Build the Docker image
Now build the Docker image in the same directory where the Dockerfile and the calculator application file are located using the below command:


COPY

COPY
docker build -t calculator .
Above command instructed Docker to build a new image using the Dockerfile in the current directory and name the image as -calculator.



Verify the Docker image
Using the docker images command, you can check the Image ID of the image. Every image has a unique Image ID.


COPY

COPY
docker images


Run Docker Container
Once the image has been built, then you can run and create a container using the below command:


COPY

COPY
docker run -d -p 8000:5000 calculator
This above command runs a new Docker container from the calculator image.



See the Docker container
With the help of the docker ps command, we can see all the running containers.


COPY

COPY
docker ps -a
Verify it using the curl command
With the help of the curl command, we can verify it.


After open Backoffice and run on it.
 http://localhost:8000/add/3/2


