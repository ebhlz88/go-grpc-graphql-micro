package main

type Server struct {
	// accountClient *account.Client
	// catalogClient *catalog.Client
	// orderClient *order.Client
}

func NewGraphQLServer(accountURL, catalogURL, orderURL string) (*Server, error) {
	// accountClient, err := account.NewClient(accountURL)
	// if err != nil {
	// 	return nil, err
	// }
	// catalogClient, err := catalog.NewClient(accountURL)
	// accountClient.Close()
	// if err != nil {
	// 	return nil, err
	// }
	// orderClient, err := order.NewClient(accountURL)
	// accountClient.Close()
	// orderClient.Close()
	// if err != nil {
	// 	return nil, err
	// }
	return &Server{
		// accountClient,
		// catalogClient,
		// orderClient,
	}, nil
}

// func (s *Server) Mutation() MutationResolver {
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver {
// 	return &queryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Account() AccountResolver {
// 	return &accountResolver{
// 		server: s,
// 	}
// }

// func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
// 	return NewExecutableSchema(Config{
// 		Resolvers: s,
// 	})
// }
