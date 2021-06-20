from django.urls import path
from django.utils.translation import gettext_lazy as _

from orders import views
from orders.views import OrderCreateView, AdminOrderDetailView

app_name = 'orders'

urlpatterns = [
    path(_('create/'), OrderCreateView.as_view(), name='order_create'),
    path('admin/order/<int:pk>/', AdminOrderDetailView.as_view(), name='admin_order_detail'),
    path('admin/order/<int:order_id>/pdf/', views.admin_order_pdf, name='admin_order_pdf'),
]
