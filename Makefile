dev:
# Windows
	start air && start tailwindcss -i ./tailwind.css -o ./assets/styles.css --watch
# Unix
# air &! tailwindcss -i ./tailwind.css -o ./assets/styles.css --watch
.PHONY: dev