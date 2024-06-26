В корневой папке проекта запустите:
`docker-compose -f docker/docker-compose.yml up --build`

В ./docker/.env возможно поменять тип сохранения информации, изменив переменную 'STORAGE_TYPE'

Для проверки работы subscriptions нужно использовать следующие запросы в gql playground:

Создать пользователя:
`mutation addUser {
addUser(username: "jabri", password: "jabri", email: "jabri.muhi@yandex.ru")
}`

Создать пост:

`mutation addPost {
addPost(title: "new test", content: "new text", commentsAllowed: true, userId: 0)
}`

Создать комментарий:

`mutation addComment {
addComment(postId: 0, content: "new content", userId: 0)
}`

В окне 2:

`subscription commentAdded {
commentAdded(postID: 0){
content
}
}`

В окне 1:

`mutation addComment {
addComment(postId: 0, content: "new content2", userId: 0)
}`

Проверить результат обновления в окне 2...

`{
"data": {
"commentAdded": {
"content": "new content2"
}
}
}`

============================================================

Добавить ответ на комментарий:

`mutation addReply {
addReply(postId: 0, parentCommentId: 0, userId: 0, content: "new reply")
}`

Продолжить ветку ответов:

`mutation addReply {
addReply(postId: 0, parentCommentId: 2, userId: 0, content: "new reply reply")
}`

Получение комментариев поста:

`query getPostComments {
getPostComments(postId: 0, startLevel: 0, lastLevel: 2, limit: 5){
content
}
}`

В сигнатуру метода включена пагинация:
* startLevel, endLevel - глубина ветки с которой по котороую показываются комментарии
* limit - МАКСИМАЛЬНОЕ количество комментариев при запросе

Получение ответов на комментарии (может быть использовано в качестве кнопки "Показать ответы"):

`query getChildrenComments {
getChildrenComments(parentCommentId: 0, startLevel: 0, lastLevel: 2, limit: 5){
content
}
}`

Пагинация аналогична методу getPostComments.