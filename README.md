# bankproject

These are the following end point host by this service 


It takes 2 query param email and password, it will return the 
Authorization token which we can use for subsequent requeset

http://localhost:5000/v1/user/v1/user/new -d '{ "email":" ", "password":"" }'


It take email is as a query param and return the token 

http://localhost:5000/v1/user/gettoken  -d '{ "email":" "}



If token is expired, it will return the new token, takes email as query params

http://localhost:5000/v1/user/updatetoken -d '{ "email":" "}



It return all the branch and city , it takes 2 optional params limit and offset 

http://localhost:5000/api/banks/{branch}/{city} -d '{"limit":10,"offset":2}'


It return the detail for the bank for the given ifsc code 

http://localhost:5000/api/bank/{ifsc}


