import Fluent

struct CreateProduct: Migration {
    func prepare(on database: Database) -> EventLoopFuture<Void> {
        database.schema(Product.schema)
            .id()
            .field("name", .string, .required)
            .field("price", .double, .required)
            .create()
    }

    func revert(on database: Database) -> EventLoopFuture<Void> {
        database.schema(Product.schema).delete()
    }
}
