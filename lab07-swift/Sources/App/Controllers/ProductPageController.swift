import Fluent
import Vapor

struct ProductIndexContext: Encodable {
    let products: [Product]
}

struct ProductPageController: RouteCollection {
    func boot(routes: RoutesBuilder) throws {
        routes.get("products", use: index)
    }

    func index(req: Request) throws -> EventLoopFuture<View> {
        Product.query(on: req.db).all().flatMap { products in
            req.view.render("Products/index", ProductIndexContext(products: products))
        }
    }
}
