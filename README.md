<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-30%25-brightgreen.svg?longCache=true&style=flat)</a>
# Issue Maker
Проект создан для автоматического добавления задач ([issues](https://docs.gitlab.com/ee/user/project/issues/)) для gitlab

### Как это работает
* В папке с исполняемым файлом создайте документ `issues.yaml`.  
Узнать подробнее о разметке yaml можно [здесь](https://yaml.org/start.html).
* Внутри добавьте содержимое согласно следующему примеру:
```yaml
token: <токен доступа>
project_id: <id проекта>
milestone: <milestone> (необязательное поле. Применяется, если у конкретной задачи на задано соответствующее поле)
issues:
  - title: <заголовок задачи>
    id: <id задачи> (необязательное поле. Применяется, если надо редактиваровать задачу) 
    description: <описание задачи>
    labels:
      - <метка 1>
      - <метка 2>
    milestone: <milestone>
    weight: <вес задачи> (необязательное поле)
```
* Рекомендуется создать новый [токен доступа](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html), который будет использоваться приложением.
`project_id` своего приложения можно узнать в `Settings > General > Project ID`.
* После создания документа запустите приложение, выполнив в терминале `./issue_maker`.
Далее приложение проанализирует файл `issues.yaml` и если всё указано правильно, выведит в консоле предполаемые действия.
* Если вы не согласаны с тем, что собирается делать `issue_maker` введите `n`,`no`, либо любой другой символ.
В противном случае введите `y`, `yes`. Поле этого приложение отправит запросы в гитлаб на создание задач.  
* По окончанию выполнения программы в этой же директории появится файл `done_<текущая_дата>.yaml`, содержимое которого можно вставить в `issues.yaml` для исправления допущенных ошибок.

### Создание новой задачи или редактиварование ранее созданной
* Каждая issues в файле `issue.yaml` содержит необязательное поле `id`, если это поле не задано, то приложение создаёт новую задачу.  
* Если это поле заполнено, приложение понимает это, как редактирование задачи с соответствующим `id`.
Обратите внимание, что данный атрибут никак не связан с номером задаче, а также адресом в url.
Поэтому настоятельно рекомендуется брать данный атрибут из файла `done_<дата_создания_задачи>.yaml`.  
В случае редактирования все поля необходимо заносить заново, иначе приложение будет считать, что этот атрибут должен быть пустым.
---
###### Skriplenok Alexey skriplenok88@mail.ru