# BIRL-GO 🏋️‍♂️💪

Uma implementação em Go (Golang) da linguagem **BIRL** (*Bambam's "It's show time" Recursive Language*). O interpretador segue a filosofia "TREZE PORRA!" e "AQUI É BODYBUILDER", permitindo que você escreva códigos com a energia do Kleber Bambam.

## 🚀 Sobre o Projeto

Este repositório contém um interpretador básico para a linguagem BIRL, implementando o Lexer, Parser, AST e Evaluator. Ele suporta um subconjunto da sintaxe original, com algumas adaptações regionais (acentuação em português).

### Sintaxe Suportada

| BIRL (Este Projeto) | C / Equivalente | Descrição |
| :--- | :--- | :--- |
| `HORA DO SHOW` | `int main() {` | Início do programa |
| `BIRL` | `}` | Fim do bloco |
| `MONSTRO` | `int` | Tipo Inteiro |
| `TRAPÉZIO` | `float` | Tipo Float |
| `FRANGO` | `char` / `string` | Tipo String |
| `CE QUER VER ESSA PORRA?` | `printf` | Saída de dados |
| `BORA CUMPADE` | `return` | Retorno (em desenvolvimento) |

*Nota: Atualmente o interpretador foca em declarações de variáveis, operações matemáticas básicas e exibição de dados.*

## 🛠️ Como Executar

Certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina.

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/gabrielalmir/birl-go.git
   cd birl-go
   ```

2. **Execute um arquivo `.birl`:**
   ```bash
   go run main.go teste.birl
   ```

## 📝 Exemplo de Código

Veja como é um código "monstro" em `teste.birl`:

```birl
HORA DO SHOW
    MONSTRO a = 13;
    MONSTRO b = 24;
    CE QUER VER ESSA PORRA?(a + b);
    CE QUER VER ESSA PORRA?("BORA, CUMPADE!");
BIRL
```

## 📂 Estrutura do Repositório

- `/lexer`: Analisador léxico que transforma o texto em tokens.
- `/parser`: Transforma os tokens em uma Árvore de Sintaxe Abstrata (AST).
- `/ast`: Definições dos nós da árvore de sintaxe.
- `/evaluator`: Onde a "mágica" acontece e o código é executado.
- `/object`: Sistema de tipos internos do interpretador.
- `/tests/scripts`: Exemplos de scripts BIRL para teste.

## 🚧 Status do Desenvolvimento

Este projeto é um subconjunto da [especificação oficial do BIRL](https://birl-language.github.io/).

**Próximos passos:**
- [ ] Implementar estruturas de controle (`ELE QUE A GENTE QUER?`, `NEGATIVA BAMBAM`).
- [ ] Adicionar suporte a funções (`OH O HOMEM AI PO`).
- [ ] Suporte para entrada de dados (`QUE QUE CE QUER MONSTRÃO?`).

---
*"É 13 PORRA! BORA CUMPADE!"* 🦍
