package main

import (
    "context"
    "flag"
    "fmt"
    "os"

    "dagger.io/dagger"
)

func main() {
    ctx := context.Background()
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
    if err != nil {
        panic(err)
    }
    defer client.Close()

    flag.Parse()
    DOCKER_HUB_PAT := flag.Args()[0]
    privateRegistryHost := "https://hub.docker.com"
    username := "kawa1"
    name := "github-actions-test"
    addr, err := client.Host().Directory("./docker").
        DockerBuild().
        WithRegistryAuth(privateRegistryHost, username, client.SetSecret("dockerhub-secret", DOCKER_HUB_PAT)).
        // Publish(ctx, "taniaitest.azurecr.io/myexample:with-dockerfile")
        Publish(ctx, fmt.Sprintf("%s/%s:%s", username, name, "latest"))
    if err != nil {
        panic(err)
    }
    fmt.Printf("Published to %s", addr)
}
