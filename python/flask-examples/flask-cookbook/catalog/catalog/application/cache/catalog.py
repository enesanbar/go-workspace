import abc
from typing import Union, Awaitable, Any


class CatalogCache(abc.ABC):

    @abc.abstractmethod
    def set(self, key, value, expiration) -> None:
        pass

    @abc.abstractmethod
    def get(self, key) -> Union[Awaitable, Any]:
        pass

    @abc.abstractmethod
    def keys(self, pattern) -> Union[Awaitable, Any]:
        pass
