# OnlyOne Smart Contract

## Descripción
OnlyOne Smart Contract es el contrato desarrollado en go para la interacción de los usuarios de la wallet de este mismo con BLion y
todos los beneficios que este último puede ofrecer como la seguridad de los datos. Las transacciones que se emiten por este Smart Contract
son conocidos como `credenciales` y el usuario final puede tener acceso a esta credencial solo en modo lectura.

Para entender un poco más sobre OnlyOne le recomendamos ver el siguiente video explicativo [OnlyOne](https://www.youtube.com/watch?v=qosQwcuYLwM)


## Modelos de identidad

![Modelos de Identidad](https://www.bjungle.net/assets/img/banners/banner-modelos-identidad.jpeg)

## Alcance

El presente servicio está diseñado y construido para ser utilizado en la red de BLion mediante su BLion Virtual Smart en relación con el 
aplicativo mobile, no tiene alcance para todo público y está restringido a entidades certificadoras y emisoras que son las únicas que pueden
emitir credenciales.

## Casos de uso

* Se requiere para la creación de credenciales mediante la información que el ente emisor y/o certificador provea al servicio
* Se requiere poder obtener las credenciales por categoria
* Se requiere poder obtener todas las credenciales de un usuario
* Se requiere poder crear categorias con las que se podrá organizar las credenciales
* Se requiere poder crear un usuario
* Se requiere poder iniciar sesión con un usuario ya existente
* Se requiere poder obtener los datos de un usuario (Nombres, apellidos, número de identification, foto, etc.)
* Se requiere poder guardas la información de customisación de una credencial con respecto a lo que un ente emisor/certificador envie o requiera

## Casos de uso no soportados

* Se requiere poder cambiar o modificar los datos de una credencial en caso de error a la hora de crearlo
* Se requiere poder cambiar o actualizar la foto de perfil del usuario
* Se requiere poder eliminar una credencial

## Arquitectura

![Arquitectura BLion](https://lh3.googleusercontent.com/fife/AAWUweWqgKXdZdpzQmQ5aKjeDtD9I6gNNmItwG8SqAs_94SIIo40f7ndPup6YL92THAMLBgW4riBpxwGhAJdy6ajILyOqE-bPO-mtArb-TC_g0esL-0dk3GWB5uZCc3cJWh3G8RbnKo2uNxpVpVrxOGvRBCDjPU33rEAfJdEA2DFWTTpI_45-UtYsKaQf40HlaqWRtrFf_7147PVtJqZtveY1w3JRmtE8kAqX751RWtMfRVYRgeox9meIzr6PghGJ2N4vIW8fU-azuHVlpQqe_OM6rBWfCPlFFd_254AbPlLeNL7r-IW1wZPxjcAf6gT9Pe0jcK9Vu2-4p2PSsBTKBBdxlSLJBkQqiVRqzEy8zEilIcMpNrL42LspBUPH4CrZhKP11BuNHdDUFQRWMgRBuVPsOdWfVF--IV8S6cKnDt7DIo_PvI5S5XHwD9bbFzq8Nu4eavCPeQHrWqcKjqPVW9311r5Bju-5USBM2MgGM-ebooxz0S-Aysrnt1v7_WkAdu1y3jMkDtofLTxeEaXC-__DW1K2xYVGhscdEA-jXM8aEI-29sk8diJY5IntNUFw6aHHeOPUAm3zRclKx6FQjgpRPjUyNzgaeiIzwZu1vo1383CZkCI772kWECMUeAE1R430ZvsMBfgfHTtx5LaO5P2fqFImVjKi5DHSeh-rJCqL7EcIjChu3MwZ305EcSy7gZhIdSWcH1NtgeofeaMhxCa5cZ_PzdKGthoh-GpqeDeSlKKNzXHouZ2SBwCQHNpP-QlUSJKdtZkjjc6eDmUL7DInGfWRiiRyELYX_8B486KJJtPtVv-pIX_lu6-Fo0RAeIYqOcAQwTVz8XnmOQpJzBwQNOnjCNQRMsCvLLVa7B2BDqa8jAdS38fWpDlGwaHTyZuIl13dVMMbohaY6ivYEOHUO_yvIffh5NTD3BBG3lBGImjwywuGNCjS2W95GJJret4zCBjKCfp-hzEeVkJACNg7Kbsijdw8saoRH4uQ6GPifB9zjSfGuP8frV-SH1wmLosLB7wW5JtUM9L61scgYS2zlhUbKoxgac5ngpvqnfl3DstbevwMpJBBsQsVlgB3jqc_gngN-ub9COrJZDcPFG9zUXfYURk7N2dQB74i-m4BzeCF7PUTuh__3ViANK0HiL3cUZbHMpnLJfDLXH9ZbJiexCHPzyopW9aKdUl61n6wZwiQVuDf3S8vWlOlKi5_EsjO0PoVV9rITNcvXSch9ivOIYr9l5dwloucdOjLElWphWnu8DqBE7yLnU7nkaPC5bfvhwyjRXYMt75j42I37r6gFisUCrmtobbOodeeEikbEQr=w977-h937)

### Diagramas

### Modelos de datos
El presente diagrama muestra el modelo de datos que se utiliza para el servicio.

![Modelo Relacional](https://lh3.googleusercontent.com/fife/AAWUweUzIfkWY7dyA5jX-enrKmZmKuJewwNMqyq5vdfQsSqCpBcUUV5O4YJUJVKcED6VmGWoRuRFW5lDjKWIdYZ5MWOTZ4J7mKSNiMIsNAvJTCzzXgyXqZ6HssrO8XNYjFUP493sNEQUf9WaOtwXkLxpBYRtSoauGDZcoFQla6b_wCxMP33ndI_UNwQ89BrzHqcFeSfWalse3ZJheSW8-ubsj5QYoeaNJKrDyk6JuJWGUFS4ako6pi-TayEWu7qsXCnxJ_iA9N512iUUVAfrwKiOSn_yNLMAwnG2gIBmYQg-KSqrH0cVPtiQd0rtn5Z5rYWICeV-L9qjeMh0qhimfTqtbOTxEod6PvoGc5K5jBlEmQoYKE1EaNnfUzYN4DQMDWT9iH4k74etM6Fd18jUE4s86gUPOH2iwefozjJmTKjUQ8OoWBQreFW0kTHBoIXEcx7GzbFtDp5o0E2dHQYQ77r-WXYB2mvuBKlBWD5WZ8a6jkUGhHGx_3cWyEQWCEm5XH5zXGoaGdWeMIb6dwIbjGR4tN_niuEOy2pXSzXmWof0nDBiQMQuGvu5D3pIVg9uKwusfRP20fTulVxDs8GSs9bGrmVPFMyqrta5xrFFfbhCmOpkfrASBXq49TeqJ75UjfWt4kPCFPU9MG3xklstTnsI69eLwpfwj8I7La5bfNDMVCVGIMjU8v0lgqm7W29ZpB7iXlT5Suatp8w9RD2lA6XsGvlecrwEiEgSeMbf4p-M1STIEBz-ynojP5Sy-13bWRSoKvZBCA1UN5SSEuUp_4We-GopatXD69qjfmx3Ws9YBHjgF9bDGYaeGjfUE-59BKzxh_Tg0oPIds9gQ2St8jQxTg9ctBEaWlW0CdKafmn2veFM34voN3tJCKdmGx0JHbsgopRDD96CnpHl44rkgWIqUnIVG7HkW3OXzapJr5vkAK2yDP1w01JcAG5GJy8B5ay04OvJcspLlG_TexLMWa5qZjPXRNRzt93yXuE3RpRvubMFvZWamvIS2RJy55k3OszZsJPXnYzoilX5AkFSRzalDoC3sK6LDr0yecZOL3x_7aaObp3ozcEZ0mYtCwjnjluoc02sXICu3FlFIlDPlk9itzhop6gT9jwM2DmS0nHxiXRjR9ZZPDerRTGa_9fN3IqBMVict5NqmjUm-l7OSZ5Mo-JFgXk81_tiJAM3TMrmRJ_TTgVDTcmpmm1tqExXcsUf7SHTpcCscrr6dR1uo2wWbJ-Wq1560mXXpDpmz-F4ivz8vxaoP8b7N60cLQB_cvjhWNfO1622OLVgC3CI59OU9mURxhBCLdzGkouk-UmCJRol=w1920-h954)

## Limitaciones

Este servicio presenta las siguientes limitaciones con respecto a sus funciones y/o caracteristicas:

* Solo se Puede ejecutar en la BLion Virtual Smart

## Costos

La creación de una credencial por medio del Smart Contract de OnlyOne está regulado bajo los costos de BLion además de algunos costos adicionales como un backup de la información
o el envío de adjuntos de gran tamaño. Todos los costos y aproximaciones de este mismo se realizarán con el área correspondiente a ventas.  
Dichos costos o aproximados se pueden obtener de la siguiente en el simulador de la página de [BLion](https://www.bjungle.net).

## Instalación y Ejecución
Este servicio está construido para correr en la BLion Virtual Smart de BLion Blockchain por ende tiene que pasar por regulaciones y aprobaciones para poderse instalar y ejecutar 
en la red de BLion. Para Mayor información sobre los `Smart Contract` en BLion consulte la siguiente página [BLion](https://www.bjungle.net).



