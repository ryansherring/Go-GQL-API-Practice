package main


import( 
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

//Schema types (remember we're not in nodeJS anymore, toto)

type Post struct {
	ID int
	Title string
	User User
	Comments []Comment 
}
type User struct {
	Name string
	Posts []int
}

type Comment struct {
	Body string
}

func populate() []Post {
	user := &User{Name: "Ryan Sherring", Posts: []int {1}}
	post := Post{
		ID: 1,
		Title: "Go GraphQL Server",
		User: *user,
		Comments: []Comment{
			Comment{Body: "First Comment"},
		},
	}
	var posts []Post
	posts = append(posts, post)

	return posts
}


func main() {
	fmt.Println("GraphQL server up and running")

	posts := populate()
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	var commentType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"body": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var userType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Posts": &graphql.Field{
					Type: graphql.NewList(graphql.Int),
				},
			},
		},
	)

	var postType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Post",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"user": &graphql.Field{
					Type: userType,
				},
				"comments": &graphql.Field{
					Type: commentType,
				},
			},
		},
	)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	fields := graphql.Fields{
		"post": &graphql.Field{
			Type: postType,
			Description: "Get Post by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, post := range posts {
						if int(post.ID) == id {
							return post, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type: graphql.NewList(postType),
			Description: "Get List of Posts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return posts, nil
			},
		},
	}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	

	// defines object config
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	
	// defines schema config
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	
	// creates the schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new GQL Schema, err %v", err)
	}

	query := `
		{
			post(id:1) {
				title
				user {
					Name
					Posts
				}
			}
		}
	`
	// check the terminal. The query should be in a data object. Run a test fetch in python and react to verify

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, errors %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}