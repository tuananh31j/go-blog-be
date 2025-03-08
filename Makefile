.PHONY: redis-server

redis-server:
	docker start redis-server || docker run --name redis-server -d -p 6379:6379 redis --requirepass "1122"


redis-cli:
	docker exec -it redis-server redis-cli


