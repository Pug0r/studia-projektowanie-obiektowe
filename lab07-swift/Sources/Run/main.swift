import App
import Vapor

let app = Application()

defer {
    app.shutdown()
}

try configure(app)
try app.run()
