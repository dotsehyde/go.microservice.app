entCmd = go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert,sql/modifier
schemaPath = /ent/schema
entcm = go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

run:
	go run cmd/main.go
restart:
	@echo Stopping down Docker images if running...
	docker compose -f compose.yml down -v
	@echo Restarting Docker images in watch mode...
	docker compose -f compose.yml up --build
	@echo Docker images started!
watch:
	@echo Starting Docker images in watch mode...
	docker compose -f compose.yml up
	@echo Docker images started!
watch-d:
	@echo Starting Docker images in detatch mode...
	docker compose -f compose.yml up -d
	@echo Docker images started!
watch-down:
	@echo Stopping Docker images ...
	docker compose -f compose.yml down -v
	@echo Docker images stopped!
prune-image:
	@echo Removing all images without at least one container associated to them...
	docker image prune -a
	@echo Docker images pruned!
prune-system:
	@echo Removing all images without at least one container associated to them...
	docker system prune -a
	@echo Docker images pruned!
ent:
	@echo Creating ent Model ...
	cd ../organization && go run entgo.io/ent/cmd/ent new $(name)
	@echo Model created successfully!
ent-landlord:
	@echo Creating ent Model ...
	cd ../admin && go run entgo.io/ent/cmd/ent new $(name)
	@echo Model created successfully!
gen-ent-organization:
	@echo Gnerating ent Model ...
	cd ../organization && $(entCmd) .$(schemaPath) \
	&& cp -rp ../organization$(schemaPath) ../worker/ent && cd ../worker && $(entCmd) .$(schemaPath) \
	&& cp -rp ../organization$(schemaPath) ../launch/ent && cd ../launch && $(entCmd) .$(schemaPath) \
	# && cp -rp ../organization$(schemaPath) ../spap6$(schemaPath) && cd ../spap6 && $(entCmd) .$(schemaPath) \
	# && cp -rp ../organization$(schemaPath) ../tech$(schemaPath) && cd ../tech && $(entCmd) ./ent/schema \
	# && cp -rp ../organization$(schemaPath) ../agency$(schemaPath) && cd ../agency && $(entCmd) .$(schemaPath) \
	# && cp -rp ../organization$(schemaPath) ../growth$(schemaPath) && cd ../growth && $(entCmd) .$(schemaPath) \
	@echo Model generated successfully!
gen-ent-launch:
	@echo Gnerating ent Model ...
	cd ../organization && $(entCmd) .$(schemaPath) \
	&& cp -rp ../organization$(schemaPath) ../worker$(schemaPath) && cd ../worker && $(entCmd) .$(schemaPath) \
	&& cp -rp ../launch$(schemaPath) ../launch$(schemaPath) && cd ../launch && $(entCmd) .$(schemaPath) \
	# && cp -rp ../spap6$(schemaPath) ../spap6$(schemaPath) && cd ../spap6 && $(entCmd) .$(schemaPath) \
	# && cp -rp ../tech$(schemaPath) ../tech$(schemaPath) && cd ../tech && $(entCmd) ./ent/schema \
	# && cp -rp ../agency$(schemaPath) ../agency$(schemaPath) && cd ../agency && $(entCmd) .$(schemaPath) \
	# && cp -rp ../growth$(schemaPath) ../growth$(schemaPath) && cd ../growth && $(entCmd) .$(schemaPath) \
	@echo Model generated successfully!

gen-ent-admin:
	@echo Gnerating ent Model ...
	cd ../admin && $(entCmd) .$(schemaPath) \
	&& cp -rp ../admin/ent/schema ../worker/entlandlord && cd ../worker && $(entCmd) ./entlandlord/schema
	@echo Model generated successfully!
print-schema:
	cd ../$(dir) && go run -mod=mod entgo.io/ent/cmd/ent describe ./$(name)/schema
docker-install:
	@echo Installing docker...
	sudo apt-get update
	sudo apt-get install ca-certificates curl
	sudo install -m 0755 -d /etc/apt/keyrings
	sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
	sudo chmod a+r /etc/apt/keyrings/docker.asc
	echo \
	"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
	$(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
	tee /etc/apt/sources.list.d/docker.list > /dev/null
	sudo apt-get update
	sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
	@echo Done installing docker!