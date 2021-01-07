FROM rust:1.49.0

ENV CARGO_TERM_COLOR always
RUN apt-get update && apt-get install -y libpq-dev cmake

# create empty project for caching dependencies
RUN USER=root cargo new --bin /movie-service/docker-build
COPY /common-utils/ ./movie-service/common-utils/
WORKDIR /movie-service/docker-build
COPY /Cargo.lock ./
COPY /movie-service/Cargo.toml ./
# cache dependencies
RUN cargo install --path . --locked

COPY /movie-service/ ./
RUN cargo install --path . --locked

FROM debian:buster-slim
RUN apt-get update && apt-get install -y libpq-dev curl
COPY --from=0 /usr/local/cargo/bin/movie-service /usr/local/bin/movie-service
CMD ["movie-service"]
EXPOSE 9092
