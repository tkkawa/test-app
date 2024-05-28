// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	"dagger.io/dagger"
// )

// func main() {
// 	if err := build(); err != nil {
// 		fmt.Println(err)
// 	}
// }

// func build() error {
// 	ctx := context.Background()

// 	// クライアントを初期化して Dagger Engine に接続する
// 	// dagger.WithLogOutput でログの出力先を指定できる
// 	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer client.Close()

// 	// Docker イメージを取得する
// 	container := client.Container().From("golang:1.19")

// 	// カレントディレクトリをコンテナにマウントする
// 	src := client.Host().Directory(".")
// 	container = container.
// 		WithMountedDirectory("/workspace", src).
// 		WithWorkdir("/workspace")

// 	// テストを実行
// 	// container = container.WithExec([]string{"go", "test", "-v", "./..."})

// 	path := "build/"
// 	outpath := filepath.Join(".", path)
// 	err = os.MkdirAll(outpath, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}
// 	// container = container.Exec(dagger.ContainerExecOpts{
// 	// 	Args: []string{"go", "build", "-o", "build/"},
// 	// })

// 	// output, err := container.Directory(path).ID(ctx)
// 	// if err != nil {
// 	// 	return err
// 	// }


// 	container = container.WithExec([]string{"go", "build", "-o", outpath})

// 	outputs := client.Directory()
// 	outputs = outputs.WithDirectory(outpath, container.Directory(outpath))

// 	// パイプラインを実行する
// 	// Export でビルド先のディレクトリをホストに書き込む
// 	if _, err := outputs.Export(ctx, "."); err != nil {
// 		panic(err)
// 	}

// 	return nil
// }
