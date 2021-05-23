# REST API для работы с RabbitMQ

Приложение для работы с брокером сообщений RabbitMQ через REST API. Основная мысль - что одиночные сообщения отправлять долго, поэтому делаем API, которое позволяет сначала сделать запрос на открытие соединения с брокером сообщений, а затем отправлять сообщения в эту очередь через API. После отправки всех сообщений соединение закрывается.