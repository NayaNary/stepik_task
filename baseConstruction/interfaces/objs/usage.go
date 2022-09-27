package objs

import "stepik_task/baseConstruction/interfaces"

type usage struct {
	service string
	interfaces.Counter
}

func NewUsage(service string) *usage {
    return &usage{service, &counter{}}
}

func (u *usage)Service()string{
	return u.service
}