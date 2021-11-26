package main

// no coderunner não funcionará, pois ele executa apenas este arquivo (que não tem 'imports'). Contudo, no terminal, utilizando o comando 'go run *go' (que equivale a pegar todos os arquivos com final .go, no caso, 'go run funcoes.go outroarquivo.go'), pegará todos os arquivos da pasta como se fosse um único:
func main() {
	result := sum(1, 2)
	printing(result)
}
