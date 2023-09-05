# Builder stage
FROM maven:3.8.3-openjdk-11-slim AS builder
WORKDIR /opt
COPY . /opt/
RUN mvn clean install -DskipTests
RUN mkdir -p target/dependency && (cd target/dependency; jar -xf ../*.jar)

# Final Stage
FROM openjdk:11-jre-slim AS final
RUN addgroup app-user && adduser app-user --gecos "" --disabled-password  --no-create-home --ingroup app-user
USER app-user
ARG DEPENDENCY=/opt/target/dependency
ENV JAVA_OPTS=""
WORKDIR /app
COPY --from=builder ${DEPENDENCY}/BOOT-INF/lib lib
COPY --from=builder  ${DEPENDENCY}/META-INF META-INF
COPY --from=builder  ${DEPENDENCY}/BOOT-INF/classes .
ENTRYPOINT ["sh","-c","java $JAVA_OPTS -cp /app:/app/lib/* com.demo.app.Application"]