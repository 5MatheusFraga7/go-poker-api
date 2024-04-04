Temos: 
  - Gerador de deck de cartas
  - Gerador de carta na mão dos players e da mesa
  - Sistema que confere combinações


O que precisamos:

  []=> Gerar uma rodada 

    -> Uma rodada é um json contendo: 
      {
        players: [
          {
            name: 'Player 1', 
            id: 1, 
            hand: {
              {value: '', suit: ''}, {value: '', suit: ''}
            }
          }
        ],
        tableCards: [
          {value: '', suit: ''}, 
          {value: '', suit: ''},
          {value: '', suit: ''}, 
          {value: '', suit: ''},
          {value: '', suit: ''}
        ]
      }
    -> Passar por param o número de players

  []=> Gerar a verificação do vencedor baseado nos jogadores na mesa
      {
        winner: [
          player:{
              name: 'Player 1', 
              id: 1, 
              hand: {
                {value: '', suit: ''}, {value: '', suit: ''}
              }
            }
          ],
        combination: [
          {value: '', suit: ''}, 
          {value: '', suit: ''},
          {value: '', suit: ''}, 
          {value: '', suit: ''},
          {value: '', suit: ''}
        ]
      }
    -> Passar param json:
        {
          players: [
            {
              name: 'Player 1', 
              id: 1, 
              hand: {
                {value: '', suit: ''}, {value: '', suit: ''}
              }
            }
          ]
          tableCards: [
            {value: '', suit: ''}, 
            {value: '', suit: ''},
            {value: '', suit: ''}, 
            {value: '', suit: ''},
            {value: '', suit: ''}
          ]
        }

