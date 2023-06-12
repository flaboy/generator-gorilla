run: public/manifest.json
	revel run -a .

build: public/manifest.json build/app/server

public/manifest.json: frontend/node_modules
	cd frontend; npx vite build --base=/public/

frontend/node_modules:
	cd frontend; npm install

build/app/server:
	revel build -a . -t build/app/server