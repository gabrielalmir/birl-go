# Documentação BIRL-GO 🏋️‍♂️

Bem-vindo à documentação oficial do **BIRL-GO**, a implementação em Go da linguagem que não faz código, faz **músculos**.

## O que é BIRL?
BIRL (**Bambam's "It's show time" Recursive Language**) é uma linguagem de programação inspirada no lendário Kleber Bambam. Esta implementação em Go foca em performance e robustez, permitindo que você saia do sedentarismo digital.

## Instalação Rápida

1. Baixe o Go (1.21+).
2. Clone o repo: `git clone https://github.com/gabrielalmir/birl-go.git`.
3. Compile: `go build -o birl main.go`.

## Guia de Sintaxe

### Estrutura do Programa
Todo programa monstro começa com `HORA DO SHOW` e termina com `BIRL`.

```birl
HORA DO SHOW
    // Seu treino aqui
BIRL
```

### Declaração de Variáveis
- `MONSTRO`: Inteiro (64 bits)
- `TRAPÉZIO`: Ponto Flutuante (64 bits)
- `FRANGO`: Texto/String

```birl
MONSTRO peso = 100;
TRAPÉZIO bf = 7.5;
FRANGO grito = "SAI DE CASA COMI PRA CARALHO!";
```

### Estruturas de Decisão
BIRL não aceita desculpas. Ou é, ou não é.

- `ELE QUE A GENTE QUER? (cond)`: IF
- `QUE NUM VAI DAR O QUE? (cond)`: ELSE IF
- `NÃO VAI DAR NÃO`: ELSE

```birl
ELE QUE A GENTE QUER? (peso > 80)
    CE QUER VER ESSA PORRA?("TA FICANDO MONSTRO!");
BIRL
NÃO VAI DAR NÃO
    CE QUER VER ESSA PORRA?("VAI TREINAR, FRANGO!");
BIRL
```

### Repetição
- `NEGATIVA BAMBAM (cond)`: WHILE
- `MAIS QUERO MAIS (init; cond; inc)`: FOR
- `SAI FILHO DA PUTA`: BREAK
- `VAMO MONSTRO`: CONTINUE

```birl
NEGATIVA BAMBAM (i < 10)
    i = i + 1;
BIRL
```

### Funções
- `OH O HOMEM AI PO`: Declaração
- `AJUDA O MALUCO TA DOENTE`: Chamada
- `BORA CUMPADE`: Retorno

```birl
OH O HOMEM AI PO MONSTRO DOBRO(MONSTRO n)
    BORA CUMPADE n * 2;
BIRL

HORA DO SHOW
    MONSTRO x = AJUDA O MALUCO TA DOENTE DOBRO(7);
BIRL
```

## Recursos Avançados

### Arrays e Fibras (Dicionários)
```birl
MONSTRO series = [10, 12, 15];
MONSTRO ficha = {"peito": "supino", "series": 4};
```

### Ponteiros
```birl
MONSTRO biceps = 45;
MONSTRO p = &biceps;
CE QUER VER ESSA PORRA?(*p);
```

### Concorrência (Go-routines)
```birl
BORA DIVIDIR O PESO AJUDA O MALUCO TA DOENTE TREINO_PESADO();
```

## Suplementação (Builtins)
- `TAMANHO(obj)`: Retorna o tamanho de um array ou string.
- `CONVERTE_MONSTRO(str)`: Converte string para inteiro.
- `IMC(peso, altura)`: Calcula o IMC.
- `DESCANSO(ms)`: Pausa a execução por N milissegundos.

## Contribuindo
Quer ajudar a deixar o BIRL-GO ainda mais pesado? Siga as instruções no [GitHub](https://github.com/gabrielalmir/birl-go).

---
*"AQUI É BODYBUILDER, PORRA!"* 🦍
