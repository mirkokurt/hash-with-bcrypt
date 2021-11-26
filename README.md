# hash-with-bcrypt
A simple showcase program that take two parameters, username and password, and produce a file with the hashed password, ideally to be stored in a database.

The hashing of the password is done with bcrypt using a standard cost of 10 (default cost). 

`
bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
`


When, in the future, the hashing cost of a password system needs to be increased in order to adjust for greater computational power, it could be easily done changing the parameter passed to GenerateFromPassword function.


