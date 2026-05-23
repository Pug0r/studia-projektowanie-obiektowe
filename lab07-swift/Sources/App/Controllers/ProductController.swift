import Fluent
import Vapor

struct ProductController: RouteCollection {
    func boot(routes: RoutesBuilder) throws {
        let products = routes.grouped("api", "products")
        products.get(use: index)
        products.post(use: create)

        products.group(":productID") { product in
            product.get(use: show)
            product.put(use: update)
            product.delete(use: delete)
        }
    }

    func index(req: Request) throws -> EventLoopFuture<[Product]> {
        Product.query(on: req.db).all()
    }

    func show(req: Request) throws -> EventLoopFuture<Product> {
        try findProduct(req)
    }

    func create(req: Request) throws -> EventLoopFuture<Product> {
        let input = try req.content.decode(ProductInput.self)
        let product = Product(name: input.name, price: input.price)
        return product.save(on: req.db).map { product }
    }

    func update(req: Request) throws -> EventLoopFuture<Product> {
        let input = try req.content.decode(ProductInput.self)
        return try findProduct(req).flatMap { product in
            product.name = input.name
            product.price = input.price
            return product.save(on: req.db).map { product }
        }
    }

    func delete(req: Request) throws -> EventLoopFuture<HTTPStatus> {
        return try findProduct(req).flatMap { product in
            product.delete(on: req.db).transform(to: .noContent)
        }
    }

    private func findProduct(_ req: Request) throws -> EventLoopFuture<Product> {
        guard let id = req.parameters.get("productID", as: UUID.self) else {
            throw Abort(.badRequest)
        }

        return Product.find(id, on: req.db)
            .unwrap(or: Abort(.notFound))
    }
}
