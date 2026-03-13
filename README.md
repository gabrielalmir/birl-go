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
| `QUE QUE CE QUER MONSTRÃO?` | `scanf` | Entrada de dados |
| `ELE QUE A GENTE QUER?` | `if` | Estrutura condicional |
| `NÃO VAI DAR NÃO` | `else` | Alternativa condicional |
| `NEGATIVA BAMBAM` | `while` | Estrutura de repetição |
| `BORA CUMPADE` | `return` | Retorno de valor |
| `OH O HOMEM AI PO` | `function` | Declaração de função |
| `AJUDA O MALUCO TA DOENTE`| `call` | Chamada de função |

*Nota: Atualmente o interpretador suporta declarações de variáveis, operações matemáticas, estruturas de controle e funções.*

## 🛠️ Como Executar

Certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina.

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/gabrielalmir/birl-go.git
   cd birl-go
   ```

2. **Execute um arquivo `.birl`:**
   ```bash
   go run main.go tests/scripts/soma.birl
   ```

## 📝 Exemplo de Código

Veja como é um código "monstro" em `tests/scripts/funcoes.birl`:

```birl
OH O HOMEM AI PO MONSTRO SOMA(MONSTRO a, MONSTRO b)
    BORA CUMPADE a + b;
BIRL

HORA DO SHOW
    MONSTRO x = AJUDA O MALUCO TA DOENTE SOMA(10, 20);
    CE QUER VER ESSA PORRA?(x);
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

Este projeto agora suporta as principais características da [especificação oficial do BIRL](https://birl-language.github.io/).

**Próximos passos:**
- [ ] Implementar `QUE NUM VAI DAR O QUE?` (else if).
- [ ] Suporte a arrays e ponteiros.
- [ ] Melhorar o tratamento de erros e mensagens.


---
*"É 13 PORRA! BORA CUMPADE!"* 🦍
