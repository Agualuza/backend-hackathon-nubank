# Bank Gamefy API

<div class="app-desc">Bank Gamefy API</div>

<div class="license-info">Apache 2.0</div>

<div class="license-url">http://www.apache.org/licenses/LICENSE-2.0.html</div>

### Endpoints

*   https://nubank-gamefy.herokuapp.com/login
*   https://nubank-gamefy.herokuapp.com/maketransaction
*   https://nubank-gamefy.herokuapp.com/register
*   https://nubank-gamefy.herokuapp.com/safebuy
*   https://nubank-gamefy.herokuapp.com/loadpersonas
*   https://nubank-gamefy.herokuapp.com/loadcategories
*   https://nubank-gamefy.herokuapp.com/persona
*   https://nubank-gamefy.herokuapp.com/balance

<div class="method"><a name="loginGet"></a>

### Login
<div class="method-path"><a name="loginGet"></a>

    get /login

</div>

<div class="method-summary">Returns user logged (<span class="nickname">loginGet</span>)</div>

### Query parameters

<b>email</b>: optional // user email
<b>password</b>: optional // user password
<b>token</b>: optional // user logged token

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

    {
      "response" : {
        "user" : {
          "password" : "******",
          "balance" : 250.0,
          "id" : 3,
          "email" : "teste@teste.com.br",
          "token" : "d44bc834b15c66926e5850d8c6fdad5a"
        }
      },
      "message" : "Success",
      "status" : "OK"
    }

* * *

<div class="method"><a name="maketransactionGet"></a>

### Make Transaction
<div class="method-path"><a name="maketransactionGet"></a>

    get /maketransaction

</div>

<div class="method-summary">Make bank transaction (<span class="nickname">maketransactionGet</span>)</div>

### Query parameters

<b>token</b>: required //user logged token
<b>type</b>: required //transaction type D - Debit ; C - Credit
<b>amount</b>: required // transaction money amount. Only positive float numbers.

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>
   
    {
      "response" : null,
      "message" : "Success",
      "status" : "OK"
    }
    
    Error :
    
     {
      "response" : null,
      "message" : "Error Message",
      "status" : "NOK"
    }



* * *

<div class="method"><a name="registerGet"></a>

### Register
<div class="method-path"><a name="registerGet"></a>

    get /register

</div>

<div class="method-summary">Register new user (<span class="nickname">registerGet</span>)</div>

### Query parameters

<b>email</b>: required //new user email
<b>password</b>: required //new user password
<b>name</b>: required //new user name


### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

    {
      "response" : null,
      "message" : "Success",
      "status" : "OK"
    }


* * *

<div class="method"><a name="safebuyGet"></a>

### Safe Buy
 
<div class="method-path"><a name="safebuyGet"></a>

    get /safebuy

</div>

<div class="method-summary">Returns risk analysis (<span class="nickname">safebuyGet</span>)</div>

### Query parameters
    
    <b>token</b>: required // user logged token. Save user historic.
    <b>persona_id</b>: required // persona id.
    <b>category_id</b>: required // product category
    <b>product_price</b>: required // product price
    

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

    {
    "status": "OK",
    "response": [
        {
            "evaluation": {
                "evaluation": "GE",
                "description": "Green End"
            },
            "blog": {
                "id": 3,
                "title": "Aumentando os riscos e os resultados",
                "post": "Lorem ipsum dolor sit amet...",
                "author": "Robson",
                "created_at": "08/03/2020"
            }
        }
    ],
    "message": "Success"
}

* * *

<div class="method"><a name="loadPersonasGet"></a>

### Load Personas
<div class="method-path"><a name="loadPersonasGet"></a>

    get /loadpersonas

</div>

<div class="method-summary">Load all personas. (<span class="nickname">loadPersonasGet</span>)</div>

### Query parameters

    no parameters

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

    {
    "status": "OK",
    "response": [
        [
            {
                "id": 1,
                "name": "João",
                "description": "Seu João 52 anos, trabalha desde os 12 anos, começou muito cedo na oficina do José que é o pai do João, aos 25 anos João começou o seu primeiro negócio, onde se apaixonou pelo mundo empresarial e a cada vez mais diversifica os seus investimentos.",
                "goal": "Aprofundar-se em financiamento",
                "photo": "",
                "factor": 1.2
            },
            {
                "id": 2,
                "name": "Maria",
                "description": "Maria 43 anos, formada em administração e ama o que faz. A Maria é gerente de uma loja do grupo x. Desde a faculdade Maria sabe que não se gasta tudo que ganha então sempre coloca um dinheiro na poupança. Sempre que pode gosta de presentear a família mas sempre escolhe presentes que não sejam superiores a sua renda.",
                "goal": "Aprofundar-se em investimento",
                "photo": "",
                "factor": 1.1
            },
            {
                "id": 3,
                "name": "Pedro",
                "description": "Pedro 35 anos, é concursado público, pedro sempre foi uma pessoa feliz pois curte a vida no hoje, faz sempre longas viagens e gosta de comprar bens de lazer, pedro não acredita que deva deixar dinheiro na poupança pois sempre fala, será que vou está aqui amanha?",
                "goal": "Aprender sobre como poupar dinheiro",
                "photo": "",
                "factor": 1
            },
            {
                "id": 4,
                "name": "Joana",
                "description": "Joana 32 anos, é professora do ensino infantil, sempre gostou de trabalhar com crianças. A Joana de vez em quando gosta de fazer umas comprinhas para a garotada porém Joana não gosta de calcular para ver se ta no momento certo. A Joana hoje entende que precisa economizar um pouco pois sempre visita o vermelho.",
                "goal": "Aprender como fazer compras saudáveis",
                "photo": "",
                "factor": 0.9
            },
            {
                "id": 5,
                "name": "Carlos",
                "description": "Carlos 22 anos, começou a trabalhar na fábrica x e está gostando muito pois conseguiu comprar um carro financiado e aproveita o carro lindo que conseguiu para sair sempre no fim de semana com amigos. O carlos diz que a bagunça dele está organizada pois sempre utiliza o crédito para pagar uma dívida mas só paga a dívida se for interessante para ele se não ele diz que não se importa em pagar no futuro com um pouco mais de juros.",
                "goal": "Aprender sobre finanças, como gastar e economizar",
                "photo": "",
                "factor": 0.8
            }
        ]
    ],
    "message": "Success"
}

* * *

<div class="method"><a name="loadCategoriasGet"></a>

### Load Categories
<div class="method-path"><a name="loadCategoriasGet"></a>

    get /loadcategories

</div>

<div class="method-summary">Load all products categories. (<span class="nickname">loadCategoriasGet</span>)</div>

### Query parameters

    no parameters

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>
   
   {
    "status": "OK",
    "response": [
        [
            {
                "id": 1,
                "name": "Casa",
                "type": "L"
            },
            {
                "id": 2,
                "name": "Carro",
                "type": "L"
            },
            {
                "id": 3,
                "name": "Viagens",
                "type": "L"
            },
            {
                "id": 4,
                "name": "TV",
                "type": "F"
            },
            {
                "id": 5,
                "name": "Videogame",
                "type": "F"
            },
            {
                "id": 6,
                "name": "Eletrônicos",
                "type": "F"
            },
            {
                "id": 7,
                "name": "Roupas",
                "type": "N"
            },
            {
                "id": 8,
                "name": "Livros",
                "type": "N"
            },
            {
                "id": 9,
                "name": "Cursos",
                "type": "N"
            },
            {
                "id": 10,
                "name": "Comida",
                "type": "P"
            },
            {
                "id": 11,
                "name": "Remédio",
                "type": "P"
            }
        ]
    ],
    "message": "Success"
}

* * *

<div class="method"><a name="personaGet"></a>

### Persona
<div class="method-path"><a name="personaGet"></a>

    get /persona

</div>

<div class="method-summary">Get persona based on the form. (<span class="nickname">personaGet</span>)</div>

### Query parameters

<b>token</b>: required //user logged token
<b>q1</b>: required // Answer to question 1. Values 0 or 1.
<b>q2</b>: required // Answer to question 2. Values 0 or 1.
<b>q3</b>: required // Answer to question 3. Values 0 or 1.
<b>q4</b>: required // Answer to question 4. Values 0 or 1.
<b>q5</b>: required // Answer to question 5. Values 0 or 1.

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

{
    "status": "OK",
    "response": [
        {
            "id": 2,
            "name": "Maria",
            "description": "Maria 43 anos, formada em administração e ama o que faz. A Maria é gerente de uma loja do grupo x. Desde a faculdade Maria sabe que não se gasta tudo que ganha então sempre coloca um dinheiro na poupança. Sempre que pode gosta de presentear a família mas sempre escolhe presentes que não sejam superiores a sua renda.",
            "goal": "Aprofundar-se em investimento",
            "photo": "",
            "factor": 1.1
        }
    ],
    "message": "Success"
}


* * *

<div class="method"><a name="balanceGet"></a>

### Balance
<div class="method-path"><a name="balanceGet"></a>

    get /balance

</div>

<div class="method-summary">Get user balance (<span class="nickname">balanceGet</span>)</div>

### Query parameters

<b>token</b>: required //user logged token

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

{
    "status": "OK",
    "response": [
        750
    ],
    "message": "Success"
}


* * *
