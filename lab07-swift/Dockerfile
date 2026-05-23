FROM swift:6.0

WORKDIR /app
COPY Package.swift ./
RUN swift package resolve

COPY . .
RUN swift build -c release

EXPOSE 8080
ENV PORT=8080

CMD [".build/release/Run"]
