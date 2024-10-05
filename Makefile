.PHONY: all commit push tag release

TAG_VERSION := v1.5.1

all: commit tag push

commit:
	git add .
	git commit -m "Обновление библиотеки"

# Удаление старого тега и создание нового
tag:
	git tag $(TAG_VERSION)

push:
	git push origin $(TAG_VERSION)
	git push origin main