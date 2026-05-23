import Vapor

struct ProductInput: Content {
    let name: String
    let price: Double
}
