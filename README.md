# Go tennis-game

Código que simula uma partida de tênis com dois jogadores, seguindo as especificações passadas na disciplina de Programação Concorrente.

Cada jogador pode ter dois possíveis estados:
1. esperando para receber a bola ou 
2. mandando a bola de volta para o adversário. 

O ponto é marcado caso a bola toque no solo da quadra adversária e o adversário não consiga devolver a bola antes do segundo toque ou que a devolva para fora dos limites da outra meia-quadra.

No caso do nosso projeto, abstraímos os casos específicos que definem a marcação da pontuação, e distribuímos os pontos de forma randômica. Ou seja, quando um jogador realiza a jogada, computamos se ele perdeu ou não. Caso tenha perdido (não importa o motivo), damos ponto para 
o jogador adversário. Caso ele tenha ganho, recebe o ponto.
    
A regra de pontuação do jogo de tênis consiste na sequência game–set–match.

**game**: sequência de pontos marcados por um jogador
* ganha quem tiver feito pelo menos quatro pontos no total e pelo menos dois pontos a mais que o adversário 

**set**: conjunto de *games*
* ganha quem tiver ganho pelo menos seis games e pelo menos dois games a mais que o oponente 

**match**: sequência de *sets* 
* sendo vencedor o jogador que ganhar três de cinco sets. 

No código implementado, defini a constante **P** para determinar o valor de pontos necessários para fechar o *game*, outra constante **G** para o valor de games para fechar um *set*, e uma constante
**S** para o valor de *sets* para marcar o *match*.
Também defini uma constante **vantege**, para determinar qual o valor da vantagem que o jogador deve ter sobre a pontuação do adversário para marcar um *game*, *set* ou *match*.