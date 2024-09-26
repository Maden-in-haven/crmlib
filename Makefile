.PHONY: all commit push tag release

# Переменная для версии тега
TAG_VERSION := v1.1.0

# Обновление библиотеки, коммит, тег и пуш
all: commit tag push

# Коммит изменений
commit:
	git add .
	git commit -m "Обновление библиотеки"

# Добавление тега
tag:
	git tag $(TAG_VERSION)

# Пуш изменений и тега
push:
	git push origin $(TAG_VERSION)
	git push origin main # Если основная ветка называется main, поменяйте на master, если нужно
