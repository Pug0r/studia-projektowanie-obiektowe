import Vapor

public func routes(_ app: Application) throws {
    try app.register(collection: ProductController())
    try app.register(collection: ProductPageController())
}
