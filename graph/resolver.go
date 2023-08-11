package graph

import "github.com/qomaindo-dev/go-bookstore/app/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BookRepository   repository.BookRepository
	ResolverEmployee repository.EmployeeRepository
	ResolverRole     repository.RoleRepository
}

// type ResolverEmployee struct {
// }

// type ResolverRole struct {
// }
