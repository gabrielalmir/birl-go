# BIRL-GO 🏋️‍♂️💪

Uma implementação em Go (Golang) da linguagem **BIRL** (*Bambam's "It's show time" Recursive Language*). Siga a filosofia "TREZE PORRA!" e "AQUI É BODYBUILDER" diretamente no seu terminal.

Esta implementação suporta a maioria das funcionalidades da especificação original com adições regionais (acentuação) e recursos como ponteiros e arrays.

## 🚀 Como Começar

### Pré-requisitos
- [Go](https://golang.org/dl/) 1.21 ou superior instalado.

### Instalação
```bash
git clone https://github.com/gabrielalmir/birl-go.git
cd birl-go
go build -o birl main.go
```

### Uso
Para rodar um arquivo `.birl`:
```bash
./birl tests/scripts/soma.birl
```

---

## 📚 Referência de Sintaxe

### Estrutura do Programa
| BIRL | C / Equivalente | Descrição |
| :--- | :--- | :--- |
| `HORA DO SHOW` | `int main() {` | Início do programa principal |
| `BIRL` | `}` | Fim de qualquer bloco |

### Tipos de Dados
| BIRL | Tipo |
| :--- | :--- |
| `MONSTRO` | Inteiro (`int64`) |
| `TRAPÉZIO` | Ponto flutuante (`float64`) |
| `FRANGO` | Caractere / String |

### Entrada e Saída
| BIRL | Função |
| :--- | :--- |
| `CE QUER VER ESSA PORRA?(exp);` | `printf` (Saída) |
| `QUE QUE CE QUER MONSTRÃO?(var);` | `scanf` (Entrada) |

### Controle de Fluxo
| BIRL | Equivalente |
| :--- | :--- |
| `ELE QUE A GENTE QUER? (cond)` | `if` |
| `QUE NUM VAI DAR O QUE? (cond)` | `else if` |
| `NÃO VAI DAR NÃO` | `else` |
| `NEGATIVA BAMBAM (cond)` | `while` |

### Funções
- **Declaração:** `OH O HOMEM AI PO [tipo] [nome]([params])`
- **Chamada:** `AJUDA O MALUCO TA DOENTE [nome]([args])`
- **Retorno:** `BORA CUMPADE [valor];`

### Ponteiros e Arrays
- **Endereço:** `&variavel`
- **Desreferência:** `*ponteiro`
- **Arrays:** `MONSTRO lista = [1, 2, 3];`
- **Acesso:** `lista[0]`

---

## 🛠️ Exemplos

### 1. Loop Monstro
```birl
HORA DO SHOW
    MONSTRO i = 0;
    NEGATIVA BAMBAM (i < 3)
        CE QUER VER ESSA PORRA?("CONTANDO...");
        CE QUER VER ESSA PORRA?(i);
        i = i + 1;
    BIRL
BIRL
```

### 2. Função com Retorno
```birl
OH O HOMEM AI PO MONSTRO DOBRO(MONSTRO n)
    BORA CUMPADE n * 2;
BIRL

HORA DO SHOW
    MONSTRO x = AJUDA O MALUCO TA DOENTE DOBRO(7);
    CE QUER VER ESSA PORRA?(x); // Imprime 14
BIRL
```

### 3. Ponteiros (AQUI É BODYBUILDER!)
```birl
HORA DO SHOW
    MONSTRO x = 13;
    MONSTRO p = &x;
    CE QUER VER ESSA PORRA?(*p); // Imprime 13
BIRL
```

---

## 📂 Documentação Online
A documentação detalhada está disponível em: [https://gabrielalmir.github.io/birl-go/](https://gabrielalmir.github.io/birl-go/)

## 🤝 Contribuição
1. Faça um Fork do projeto
2. Crie uma Branch para sua Feature (`git checkout -b feature/TreinoPesado`)
3. Commit suas mudanças (`git commit -m 'Add: Novo Exercício'`)
4. Push para a Branch (`git push origin feature/TreinoPesado`)
5. Abra um Pull Request

---
*"É 13 PORRA! SAÍ DAQUI QUE É HORA DO SHOW!"* 🦍
