from dataclasses import dataclass


@dataclass(frozen=True)
class GetProductsRequest:
    page: int
    per_page: int
