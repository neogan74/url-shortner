# go url shortener app

## Список инкрементов

<details>
  <summary>Инкремент 1</summary>

Напишите сервис для сокращения длинных URL. Требования:
* Сервер должен быть доступен по адресу: `http://localhost:8080`.
* Сервер должен предоставлять два эндпоинта: `POST /` и `GET /{id}`.
* Эндпоинт `POST /` принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом `201` и сокращённым URL в виде текстовой строки в теле.
* Эндпоинт `GET /{id}` принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом `307` и оригинальным URL в HTTP-заголовке Location.
* Нужно учесть некорректные запросы и возвращать для них ответ с кодом `400`.

</details>

### Other

<details>
<summary>Запуск автотестов</summary>

Для успешного запуска автотестов вам необходимо давать вашим веткам названия вида `iter<number>`, где `<number>` -
порядковый номер итерации.

Например в ветке с названием `iter4` запустятся автотесты для итераций с первой по четвертую.

При мерже ветки с итерацией в основную ветку (`main`) будут запускаться все автотесты.
</details>