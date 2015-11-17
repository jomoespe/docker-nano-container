FROM scratch
MAINTAINER José Moreno
LABEL Description="Minimal docker container for a simple rest service" 

ADD rest-server cert.pem key.pem /
CMD ["/rest-server"]
