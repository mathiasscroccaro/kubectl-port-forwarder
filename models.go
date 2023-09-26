package main

import (
	"fmt"
)

type Services struct {
	Services []Service `json:"services"`
}

type Service struct {
	ContainerPort int    `json:"containerPort"`
	LocalhostPort int    `json:"localhostPort"`
	ServiceName   string `json:"serviceName"`
	Namespace     string `json:"namespace"`
}

func (s Services) validate() error {
	for _, service := range s.Services {
		if err := service.validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s Service) validate() error {
	if s.ContainerPort == 0 || s.LocalhostPort == 0 || s.ServiceName == "" || s.Namespace == "" {
		return fmt.Errorf("Invalid configuration for service %v", s)
	}
	return nil
}
