package main

import "fmt"

func main() {
  
  exibeIntroducao()
  exibeMenu()
  comando := leComando()
  

  if comando == 1{
    fmt.Println("Monitorando...")
  } else if comando == 2{
    fmt.Println("Exibindo Logs...")
  } else if comando == 0{
    fmt.Println("Saindo do Programa")
  } else {
    fmt.Println("Não conheço este comando.")
  }
}

func exibeIntroducao() {
nome := "Jaqueline"
  versao := 1.1
  fmt.Println("Olá,", nome)
  fmt.Println("Este programa está na versão", versao)
}

func exibeMenu(){
fmt.Println("1 - Iniciar Monitoramento")
fmt.Println("2 - Exibir Logs")
fmt.Println("0 - Sair do Programa")
}

func leComando() int{

  var comandoLido int
  fmt.Scanf("%d,", &comandoLido)
  fmt.Println("O comando escolhido foi", comandoLido)

  return comandoLido
}