Used for first learning full CRUD RESTful APIs in Golang, then applied my GraphQL experience to make a simple read API template with posts, users, and comments, which is the most common type I've been asked to make.

My Process

Before, I was learning to build with gqlgen, which seems to be a popular methodology for applying graphql in a Go backend. However like Django in Python, a lot is templated for you, but the API is much more fleshed out than Django and can be rerun to rebuild the template to your schema objects.

I decided to move forward with learning GraphQL-Go instead as it allowed me more freedom to build something to scratch and learn the feel of it better. I only learned Go yesterday, and though I'm brand new to the language, I'm used to this pace as a bootcamp grad and actually use the momentum positively. I'm generally a mission-based person and so building a project for someone or something else (like getting a job) tends to be a lot more motivating than learning through constant passion projects or learning in a tutorial-based fashion. 

It should be noted that I normally use forum-style schemas to new learn backend tech, but I likely wouldn't use GraphQL for this use case normally. Rather, I think a REST server would serve the purpose of a forum in which CRUD is it's main feature, given some conditions. I'd probably rather want to use GraphQL to aggregate data from different sources in a single location for different clients. Nonetheless, I don't believe GraphQL or gRPC or REST servers are superior to one another, but have different use cases to one another (though I could be wrong as I'm only junior and have A LOT to learn... so my mind's still open).

Next Steps

I love GraphQL for a lot of different reasons. I didn't use subscriptions or mutations for this example, but I'm choosing to move forward to learning gRPC in Go, because I may have a job interview in this soon, I understand graphQL principles, and need to repetition in different use cases to get used to the Go language/syntax.

I'm also looking forward to getting used to gRPC. After a fair amount of research into the theory, I think it'll help me to understand/use polyglot microservices, which is a personal goal of mine. Using GraphQL with Node servers and React frontends with the past, I've always loved the customizability og GraphQL, but sometimes I worry about sacrificing cacheability or overcommunicating with the user. I'm mostly excited to understand HTTP/2 and the differences between gRPC and gRPC-web with a proxy to interface with normal HTTP calls. 