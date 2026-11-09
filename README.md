# Registro de usuarios con Cognito

Lambda en Go que se ejecuta como **trigger Post Confirmation de Amazon Cognito**: cuando una persona confirma su cuenta con el código del correo, esta función guarda su correo y su identificador en la tabla `users` de Amazon RDS para MySQL. Así, la API puede relacionar sus direcciones y pedidos con ese usuario.

> El proyecto completo se compone de tres repositorios:

| Repositorio | Contenido |
|---|---|
| [santa-maria-ecommerce](https://github.com/javierPerez900/santa-maria-ecommerce) | Frontend en Next.js |
| [santaMaria](https://github.com/javierPerez900/santaMaria) | API REST en Go (Lambda + API Gateway) |
| **santaMariaUser** (este) | Lambda en Go que guarda en la base de datos a cada usuario que confirma su cuenta en Cognito |
