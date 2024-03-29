The first thing we need is a handler. If you’re coming from an MVC-background, you can think of handlers as being a bit like controllers. They’re responsible for executing your application logic and for writing HTTP response headers and bodies.

The second component is a router (or servemux in Go terminology). This stores a mapping between the URL patterns for your application and the corresponding handlers. Usually you have one servemux for your application containing all your routes.

The last thing we need is a web server. One of the great things about Go is that you can establish a web server and listen for incoming requests as part of your application itself. You don’t need an external third-party server like Nginx or Apache.
