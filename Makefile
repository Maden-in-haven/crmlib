.PHONY: all commit push tag release

# Извлекаем последний тег и увеличиваем его версию
TAG_VERSION := $(shell git describe --tags --abbrev=0)
NEW_VERSION := $(shell echo $(TAG_VERSION) | awk -F. -v OFS=. '{$$NF++; print}')

# Основные цели
all: commit tag push

# Коммит изменений
commit:
	git add .
	git commit -m "Обновление библиотеки"

# Удаление старого тега и создание нового
tag:
	-git tag -d $(TAG_VERSION) # Удаляем локальный тег, если он существует
	-git push origin --delete $(TAG_VERSION) # Удаляем удалённый тег, если он существует
	git tag $(NEW_VERSION)

# Пуш изменений и тега
push:
	git push origin $(NEW_VERSION)
	git push origin main