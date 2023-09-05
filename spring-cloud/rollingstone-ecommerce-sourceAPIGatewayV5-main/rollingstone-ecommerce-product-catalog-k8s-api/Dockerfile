FROM adoptopenjdk/openjdk11
VOLUME /tmp
ARG JAR_FILE=build/libs/*.jar
COPY ${JAR_FILE} ./
ENTRYPOINT ["java","-jar","/rollingstone-ecommerce-product-catalog-k8s-api-1.0.jar"]