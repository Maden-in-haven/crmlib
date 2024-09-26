.PHONY: all commit push tag release

TAG_VERSION := v1.1.3

all: commit tag push

commit:
	git add .
	git commit -m "Обновление библиотеки"

# Удаление старого тега и создание нового
tag:
	-git tag -d $(TAG_VERSION)  # Игнорируем ошибку, если тега нет
	-git push origin --delete $(TAG_VERSION)  # Удаляем удалённый тег, если он существует
	git tag $(TAG_VERSION)

push:
	git push origin $(TAG_VERSION)
	git push origin main