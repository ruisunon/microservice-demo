FROM openjdk:17
COPY target/*.jar application.jar
ENTRYPOINT ["sh" ,"-c", "java -jar /application.jar"]
LABEL maintainer="Hardik Singh Behl" email="behl.hardiksingh@gmail.com"
EXPOSE 8080