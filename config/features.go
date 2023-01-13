package config

import (
    "github.com/gocondor/core"
)

// Features to control GoCondor's core features
var Features *core.Features = &core.Features{
    Database: false, // Database
    Cache:    false, // Cache
    GRPC:     false, // GRPC 
}