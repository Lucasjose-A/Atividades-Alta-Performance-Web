package main

import "fmt"

/*
RESPOSTAS DO DIAGNÓSTICO:

1. Falha A: Falha A: O site permaneceu ativo, mas o tempo
de resposta para finalizar compras saltou de 2s
para 45s sob carga, gerando desistência massiva
de usuários.

[Escalabilidade]: É Escalabilidade visto que Por causa do crescimento da utilização do sistema pelos usuários ele sobrecarregou e aumentou o seu tempo de resposta
 pois não conseguiu distribuir a carga posta a ele.

2. Falha B: Durante o pico de acessos, um erro de
concorrência no código expôs dados de pagamento
de um cliente para outro que estava navegando
simultaneamente.

[Confiabilidade]: É Confiabilidade porque o sistema deveria funcionar corretamente mesmo diante a problemas como excesso de acessos ao sistemas,
 perdendo assim a moral com os usuários pelo mal funcionamento do sistema.

3. Falha C: A equipe tentou corrigir um bug urgente,
mas descobriu que o sistema era tão complexo e
mal documentado que qualquer alteração levava
dias para ser testada e homologada, impedindo
uma solução rápida durante o evento.

[Manutenibilidade]: É Manutenibilidade pois a equipe deveria saber mexer no sistema sem nenhuma dificuldade,
 sem demorar para fazer alterações simples solucionando os problemas do sistema rapidamente para melhor eficiência.

*/
func main() {

	fmt.Println("Ambiente configurado por: Profº. Marcos André")

	fmt.Println("Partiu Carnaval, mas com o mindset de Alta Performance!")

}
