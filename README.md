# calculate-temperature
Full Cycle Challenge

API:  https://calculate-temperature-xyd2xuzs7a-uc.a.run.app/calculate-temperature?zipcode=SEU_CEP_AQUI
Exemplo:   https://calculate-temperature-xyd2xuzs7a-uc.a.run.app/calculate-temperature?zipcode=29055370

Retorno:
{
    "temp_C": 28.7,
    "temp_F": 83.6,
    "temp_K": 301.7
}


Como rodar em dev:
1. Inicie baixando o projeto na sua máquina;
2. Instale as dependências;
3. Faça a criação de uma conta na API https://www.weatherapi.com/;
4. Na sua máquina faça uma cópia do arquivo .env.example e renomeia para .env;
5. Na sua conta feita nesse site: https://www.weatherapi.com/, procute pela Api Key e cole a mesma na variável de ambiente WEATHER_API_KEY;
6. Rode o docker compose utilizando o comando: docker compose up;
7. Acesse o container e inicie a aplicação com o comando: go run main.go;
8. Para rodar os testes basta usar o comando: go test ./...
