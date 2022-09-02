# FileServerTCPF
Es un servidor que permite transferir archivos entre 2 o más clientes usando un custom protocol (protocolo no estándar) basado en TCP, utilizando Golang y telnet para la linea de comandos.

### ¿Cómo funciona?⚙️
Este es un protocolo simple de comunicación entre el cliente y el servidor, el objetivo es el cliente pueda recibir y enviar archivos. Asimismo, permitir al cliente suscribirse a canales y también enviar archivos a
canales específicos.
Para esto el cliente cuenta con los siguientes comandos:

* `/name <name>` - configurar el nombre del cliente.
* `/subscribe <name>` - Subscribirse a un canal especifico, solo uno a la vez.
* `/channels` - Lista de los canales disponibles.
* `/sendfile <filename>` - enviar archivo.
* `/quit` - desconectarse.
 ## 
 ### Testing :hammer_and_wrench:
 Para empezar ejecutar el archivo FileServe.exe ese será nuestro servidor, también se debe tener 2 o más terminales para cada uno de los clientes.
 ```
 Cliente 1:
 telnet localhost 8888
  /name Valentina
> Hello Valentina
/subscribe #friends
> welcome to #friends
> Juan joined the room
> Juan send a file example.txt 
  ¡Hola, esto es un ejemplo!
```
 ```
 Cliente 2 :
 telnet localhost 8888
  /name Juan
> Hello Juan
/subscribe #friends
> welcome to #friends
/sendfile example.txt
```
 ## Creado por
Valentina Paredes B.
