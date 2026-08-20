# BIRL-GO 🏋️‍♂️💪

Uma implementação em Go (Golang) da linguagem **BIRL** (*Bambam's "It's show time" Recursive Language*). Siga a filosofia "TREZE PORRA!" e "AQUI É BODYBUILDER" diretamente no seu terminal.

Este projeto foi inspirado na [linguagem original BIRL](https://github.com/birl-language/birl-language.github.io).

Esta implementação suporta a maioria das funcionalidades da especificação original com adições regionais (acentuação) e recursos como ponteiros e arrays.

## 🚀 Como Começar

### Pré-requisitos
- [Go](https://golang.org/dl/) 1.21 ou superior instalado.

### Instalação
**Windows:** [Baixar executável (v0.0.4)](https://github.com/gabrielalmir/birl-go/releases/download/v0.0.4/birl-windows-amd64.exe)

**Via Código (Linux/macOS):**
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
| `MAIS QUERO MAIS (init; cond; inc)`| `for` |
| `SAI FILHO DA PUTA` | `break` |
| `VAMO MONSTRO` | `continue` |
| `BORA DIVIDIR O PESO [call]` | `go routine` |

### Operadores Lógicos
- `&&` (AND lógico)
- `||` (OR lógico)

### Funções
- **Declaração:** `OH O HOMEM AI PO [tipo] [nome]([params])`
- **Chamada:** `AJUDA O MALUCO TA DOENTE [nome]([args])`
- **Retorno:** `BORA CUMPADE [valor];`
- **Funções Nativas (Builtins):** `TAMANHO(obj)`, `CONVERTE_MONSTRO(str)`, `IMC(peso, altura)`, `DESCANSO(ms)`.

### Tipos Complexos
- **Ponteiros:** `&variavel` (endereço) e `*ponteiro` (desreferência)
- **Arrays:** `MONSTRO lista = [1, 2, 3];` e `lista[0]`
- **Dicionários (Fibras):** `MONSTRO ficha = {"musculo": "Peito", "series": 4};` e `ficha["musculo"]`

### Comentários
- `//` para comentários de linha
- `/* */` para comentários de bloco

---

## 🛠️ Exemplos

### 1. Loop Monstro (FOR e Controle)
```birl
HORA DO SHOW
    MAIS QUERO MAIS (MONSTRO i = 0; i < 10; i = i + 1)
        ELE QUE A GENTE QUER? (i == 5)
            CE QUER VER ESSA PORRA?("Fadigou, parando na 5!");
            SAI FILHO DA PUTA; // Break
        BIRL
        CE QUER VER ESSA PORRA?(i);
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

## 🔐 Segurança dos artefatos

O CI testa e compila todas as plataformas suportadas, executa `go vet` e
`govulncheck`, procura conteúdo de prompt injection nos arquivos textuais e usa
ClamAV para verificar o código e os binários contra malware. Uma correspondência
em qualquer verificação bloqueia a entrega.

Em tags `v*`, cada binário é assinado sem chave pelo Cosign usando a identidade
OIDC do GitHub Actions. A release contém o binário, um bundle
`.sigstore.json` (assinatura, certificado e prova de transparência) e o arquivo
`SHA256SUMS`. Para verificar, instale o Cosign e execute, substituindo a tag:

```bash
cosign verify-blob birl-linux-amd64 \
  --bundle birl-linux-amd64.sigstore.json \
  --certificate-identity-regexp='https://github.com/gabrielalmir/birl-go/.github/workflows/release.yml@refs/tags/v.*' \
  --certificate-oidc-issuer='https://token.actions.githubusercontent.com'
sha256sum --check SHA256SUMS
```

O detector de prompt injection é uma barreira defensiva baseada em padrões; ele
não substitui revisão humana de mudanças que adicionem prompts ou dados não
confiáveis. Execute-o localmente com
`python3 scripts/security/scan_prompt_injection.py .`.

---
*"É 13 PORRA! SAÍ DAQUI QUE É HORA DO SHOW!"* 🦍
