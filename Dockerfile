FROM ubuntu:latest

ADD ./share /root/share

# install dependencies
RUN apt-get -y update
RUN apt-get -y install git 
RUN apt-get -y install vim
RUN apt-get -y install golang-go

# Expose port 3000 to real world
EXPOSE 3000
