dev:
# Windows
# start air && start tailwindcss -i ./tailwind.css -o ./static/css/styles.css --watch
# Unix
	air &! tailwindcss -i ./tailwind.css -o ./static/css/styles.css --watch
.PHONY: dev