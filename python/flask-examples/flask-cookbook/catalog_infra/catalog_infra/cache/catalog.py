from typing import Union, Any, Awaitable

from redis import Redis

from catalog.application.cache.catalog import CatalogCache


class RedisCache(CatalogCache):

    def __init__(self) -> None:
        self._redis = Redis()

    def set(self, key, value, expiration=600) -> None:
        self._redis.set(key, value, ex=expiration)

    def get(self, key) -> Union[Awaitable, Any]:
        return self._redis.get(key)

    def keys(self, pattern) -> Union[Awaitable, Any]:
        return self._redis.keys(pattern)


