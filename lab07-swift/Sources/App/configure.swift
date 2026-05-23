import Fluent
import FluentSQLiteDriver
import Vapor

public func configure(_ app: Application) throws {
    app.http.server.configuration.hostname = "0.0.0.0"
    app.databases.use(.sqlite(.file("db.sqlite")), as: .sqlite)

    app.migrations.add(CreateProduct())

    try app.autoMigrate().wait()
    try routes(app)
}
