# Hello World Go Docker Image
Este projeto fornece um Dockerfile para construir uma imagem Docker mínima que exibe a mensagem "hello world" na tela.

## Requisitos
Para executar este projeto, você precisará ter o Docker instalado em sua máquina.

## Como usar
Para criar a imagem Docker, abra o terminal e navegue até o diretório raiz deste projeto. Em seguida, execute o seguinte comando:

```docker build -t hello-world-image .```

Esse comando constrói a imagem Docker e a nomeia como "hello-world-image". O ponto no final do comando indica que o Dockerfile está no diretório atual.

Depois que a imagem for criada, você poderá executá-la com o seguinte comando:

```docker run hello-world-image```

Isso exibirá a mensagem "hello world" na tela.

## Como funciona

O Dockerfile neste projeto usa a imagem Scratch como base, que é conhecida por ser extremamente leve e compacta. Em seguida, ele adiciona um único comando para imprimir a mensagem "hello world" na tela. Esse comando é executado no momento em que a imagem é iniciada.

## Contribuindo

Se você quiser contribuir para este projeto, sinta-se à vontade para abrir um pull request ou discutir possíveis melhorias na seção de problemas. =D
