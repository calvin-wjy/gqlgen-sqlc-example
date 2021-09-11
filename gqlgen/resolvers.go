package gqlgen

import (
	"context"

	"github.com/fwojciec/gqlgen-sqlc-example/pg"
)

// Resolver connects individual resolvers with the datalayer.
type Resolver struct {
	Repository pg.Repository
}

// Student returns an implementation of the StudentResolver interface.
func (r *Resolver) Student() StudentResolver {
	return &studentResolver{}
}

// Mutation returns an implementation of the MutationResolver interface.
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns an implementation of the QueryResolver interface.
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type studentResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateStudent(ctx context.Context, data StudentInput) (*pg.Student, error) {
	student, err := r.Repository.CreateStudent(ctx, pg.CreateStudentParams{
		Name: data.Name,
		Nim:  data.Nim,
	})
	if err != nil {
		return nil, err
	}
	return &student, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Students(ctx context.Context) ([]pg.Student, error) {
	return r.Repository.ListStudents(ctx)
}
