FROM adoptopenjdk/openjdk11
VOLUME /tmp
ARG JAR_FILE=build/libs/*.jar
COPY ${JAR_FILE} ./
ENTRYPOINT ["java","-jar","/rollingstone-ecommerce-category-k8s-api-1.jar"]