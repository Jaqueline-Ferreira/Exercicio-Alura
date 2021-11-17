package main

import "fmt"
import "os"
import "net/http"

func main() {
  
  exibeIntroducao()
  
for {
    exibeMenu()

    comando := leComando()
    
    if comando == 1{
      iniciarMonitoramento()
    } else if comando == 2{
      fmt.Println("Exibindo Logs...")
    } else if comando == 0{
      fmt.Println("Saindo do Programa")
      os.Exit(0)
    } else {
      fmt.Println("Não conheço este comando.")
      os.Exit(-1)
    }
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

func iniciarMonitoramento() {
  fmt.Println("Monitorando...")
  site := "https://random-status-code.herokuapp.com"
  resp, _ := http.Get(site)

  if resp.StatusCode == 200 {
  fmt.Println("Site: ", site, "foi carregado com sucesso!")
  }else {
    fmt.Println("Site: ", site, "esta com problemas. Status Code:", resp.StatusCode)
  }

}