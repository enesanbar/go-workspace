from django.conf import settings
from django.contrib.admin.views.decorators import staff_member_required
from django.http import HttpResponse
from django.shortcuts import render, get_object_or_404
from django.template.loader import render_to_string
from django.utils.decorators import method_decorator
from django.views.generic import FormView, TemplateView, DetailView

from cart.cart import Cart
from orders.forms import OrderCreateForm
from orders.models import OrderItem, Order
from .tasks import order_created
import weasyprint


class OrderCreateView(FormView):
    """
    Displays Checkout form and order information from the cart.
    Note: Cart is already in a context processor, so there's no need to pass it to context again.
    """
    form_class = OrderCreateForm
    template_name = 'orders/order/create.html'
    cart = None

    def dispatch(self, request, *args, **kwargs):
        self.cart = Cart(request)
        return super(OrderCreateView, self).dispatch(request, *args, **kwargs)

    def form_valid(self, form):
        order = form.save(commit=False)
        if self.cart.coupon:
            order.coupon = self.cart.coupon
            order.discount = self.cart.coupon.discount
        order.save()

        for item in self.cart:
            OrderItem.objects.create(order=order, product=item['product'],
                                     price=item['price'], quantity=item['quantity'])

        # clear the cart
        self.cart.clear()

        # launch asynchronous task
        order_created.delay(order.id)

        return render(self.request, 'orders/order/created.html', {'order': order})


@staff_member_required
def admin_order_detail(request, order_id):
    order = get_object_or_404(Order, id=order_id)
    return render(request, 'admin/orders/order/detail.html', {'order': order})


@method_decorator(staff_member_required, name='dispatch')
class AdminOrderDetailView(DetailView):
    model = Order
    template_name = 'admin/orders/order/detail.html'


@staff_member_required
def admin_order_pdf(request, order_id):
    order = get_object_or_404(Order, id=order_id)
    html = render_to_string('orders/order/pdf.html', {'order': order})
    response = HttpResponse(content_type='application/pdf')
    response['Content-Disposition'] = 'filename=order_{}.pdf"'.format(order.id)
    weasyprint.HTML(string=html)\
        .write_pdf(response, stylesheets=[weasyprint.CSS(settings.STATIC_ROOT + 'css/pdf.css')])
    return response
