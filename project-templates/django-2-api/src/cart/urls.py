from django.urls import path
from django.utils.translation import gettext_lazy as _

from cart.views import CartDetailView, AddCartView, RemoveCartView

app_name = 'cart'

urlpatterns = [
    path('', CartDetailView.as_view(), name='cart_detail'),
    path(_('add/<int:product_id>/'), AddCartView.as_view(), name='cart_add'),
    path(_('remove/<int:product_id>/'), RemoveCartView.as_view(), name='cart_remove'),
]
