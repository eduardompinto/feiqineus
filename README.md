# Feiqineus (Fake news)

A ideia desse projeto é bem simples. 
Criar um bots e um endpoint http para checagem de fatos. 
Até agora ele consiste em dois apps um em Go e um em Python.
O app em Python é apenas para utilizar as bibliotecas do nltk para manipular texto, ou seja, 
remover stopwords, reduzir as palavras a sua raíz e remover acentuação. 
A ideia é diminuir a quantidade de entradas na tabela de mensagens salvando apenas se a entrada normalizada não existir 
na tabela. 

Por exemplo:

`Gatos são extraterrestres` -> `extraterrestr gat sao`
`Gatas são extraterrestre` -> `extraterrestr gat sao`

```
Existe um imovel de 400 m², instalado em um terreno de 720 m desocupado, na Rua Paula Ney, 446,
Vila Mariana-SP, avaliado em mais de R$ 2.800.000. A construção, da década de 70, faz parte do 
espólio do pai de um rico médico infectologista. Esse médico, por acaso, e o Dr. Marcos Boulos,
pai do ‘sem-teto’ Guilherme Boulos
```

```
2800000 400 446 70 720 acas avali boul construca dec desocup dr espoli exist faz guilherm imovel
infectolog instal m m2 marc marianasp medic ney pai part paul r ric rua semtet terren vil
```

A partir disso, eu faço a buscando usando `pg_trgm`(https://www.postgresql.org/docs/current/pgtrgm.html) com um limiar
de 0.6 (a similaridade vai de 0 a 1).

O app em Go recebe as mensagens, faz a chamada no app em python e responde as mensagens. 
Não existe um motivo em específico para ele ser em Go, apesar da lib do telegram em Go ser maravilhosa.
Escolhi Go porque preciso praticar. 
