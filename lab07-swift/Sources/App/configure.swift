import Fluent
import FluentSQLiteDriver
import Leaf
import Vapor

public func configure(_ app: Application) throws {
    app.http.server.configuration.hostname = "0.0.0.0"
    app.views.use(.leaf)
    app.databases.use(.sqlite(.file("db.sqlite")), as: .sqlite)

    app.migrations.add(CreateProduct())

    try app.autoMigrate().wait()
    try seedProductsIfNeeded(app)
    try routes(app)
}

private func seedProductsIfNeeded(_ app: Application) throws {
    let count = try Product.query(on: app.db).count().wait()
    guard count == 0 else {
        return
    }

    let samples: [Product] = [
        Product(name: "Kawa", price: 12.5),
        Product(name: "Herbata", price: 8.0),
        Product(name: "Czekolada", price: 6.5)
    ]

    try samples.create(on: app.db).wait()
}
