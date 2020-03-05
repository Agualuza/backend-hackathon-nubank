# Bank Gamefy API

<div class="app-desc">Bank Gamefy API</div>

<div class="license-info">Apache 2.0</div>

<div class="license-url">http://www.apache.org/licenses/LICENSE-2.0.html</div>

## Access

### Table of Contents

#### [Default](#Default)

*   https://nubank-gamefy.herokuapp.com/login
*   https://nubank-gamefy.herokuapp.com/maketransaction
*   https://nubank-gamefy.herokuapp.com/register
*   https://nubank-gamefy.herokuapp.com/safebuy

# <a name="Default">Default</a>

<div class="method"><a name="loginGet"></a>

<div class="method-path"><a name="loginGet"></a>

    get /login

</div>

<div class="method-summary">Return user logged (<span class="nickname">loginGet</span>)</div>

### Query parameters

<div class="field-items">

<div class="param">email (optional)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User email</div>

<div class="param">password (optional)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User password</div>

<div class="param">token (optional)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User logged token</div>

</div>

### Return type

<div class="return-type">[inline_response_200](#inline_response_200)</div>

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

### Produces

This API call produces the following media types according to the <span class="header">Accept</span> request header; the media type will be conveyed by the <span class="header">Content-Type</span> response header.

*   `application/json`

### Responses

#### 200

OK [inline_response_200](#inline_response_200)</div>

* * *

<div class="method"><a name="maketransactionGet"></a>

<div class="method-path"><a name="maketransactionGet"></a>

    get /maketransaction

</div>

<div class="method-summary">Make bank transaction (<span class="nickname">maketransactionGet</span>)</div>

### Query parameters

<div class="field-items">

<div class="param">token (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User logged token</div>

<div class="param">type (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — Transaction type. D - Debit (sub), C - Credit (sum)</div>

<div class="param">amount (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — Amount of money in the transaction</div>

</div>

### Return type

<div class="return-type">[inline_response_200_1](#inline_response_200_1)</div>

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

    {
      "response" : "response",
      "message" : "Success",
      "status" : "OK"
    }

### Produces

This API call produces the following media types according to the <span class="header">Accept</span> request header; the media type will be conveyed by the <span class="header">Content-Type</span> response header.

*   `application/json`

### Responses

#### 200

OK [inline_response_200_1](#inline_response_200_1)

#### 409

NOK [inline_response_409_1](#inline_response_409_1)</div>

* * *

<div class="method"><a name="registerGet"></a>

<div class="method-path"><a name="registerGet"></a>

    get /register

</div>

<div class="method-summary">Register new user (<span class="nickname">registerGet</span>)</div>

### Query parameters

<div class="field-items">

<div class="param">email (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User email</div>

<div class="param">password (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User password</div>

<div class="param">name (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User full name</div>

</div>

### Return type

<div class="return-type">[inline_response_200_1](#inline_response_200_1)</div>

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

    {
      "response" : "response",
      "message" : "Success",
      "status" : "OK"
    }

### Produces

This API call produces the following media types according to the <span class="header">Accept</span> request header; the media type will be conveyed by the <span class="header">Content-Type</span> response header.

*   `application/json`

### Responses

#### 200

OK [inline_response_200_1](#inline_response_200_1)

#### 409

NOK [inline_response_409](#inline_response_409)</div>

* * *

<div class="method"><a name="safebuyGet"></a>

<div class="method-path"><a name="safebuyGet"></a>

    get /safebuy

</div>

<div class="method-summary">Return risk analysis rate (<span class="nickname">safebuyGet</span>)</div>

### Query parameters

<div class="field-items">

<div class="param">profile_id (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User profile id</div>

<div class="param">product_price (required)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — Product price</div>

<div class="param">token (optional)</div>

<div class="param-desc"><span class="param-type">Query Parameter</span> — User logged token</div>

</div>

### Return type

<div class="return-type">[inline_response_200_2](#inline_response_200_2)</div>

### Example data

<div class="example-data-content-type">Content-Type: application/json</div>

    {
      "rate" : "0.77",
      "user_profile" : {
        "payment" : 2000.0,
        "id" : 3
      },
      "message" : "Success",
      "status" : "OK"
    }

### Produces

This API call produces the following media types according to the <span class="header">Accept</span> request header; the media type will be conveyed by the <span class="header">Content-Type</span> response header.

*   `application/json`

### Responses

#### 200

OK [inline_response_200_2](#inline_response_200_2)</div>

* * *
