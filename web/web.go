package web

//go:generate npm run build
//go:generate go-bindata -fs -o web_gen.go -pkg web -prefix dist/ ./dist/...
