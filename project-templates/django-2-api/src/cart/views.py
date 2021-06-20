from django.shortcuts import render, get_object_or_404, redirect
from django.utils.decorators import method_decorator
from django.views import View
from django.views.decorators.http import require_http_methods
from django.views.generic import FormView, TemplateView

from cart.cart import Cart
from cart.forms import CartAddProductForm
from coupons.forms import CouponApplyForm
from shop.models import Product


@method_decorator(require_http_methods(["POST"]), name='dispatch')
class AddCartView(FormView):
    form_class = CartAddProductForm
    cart = None
    product = None

    def dispatch(self, request, *args, **kwargs):
        self.cart = Cart(request)
        self.product = get_object_or_404(Product, id=kwargs['product_id'])
        return super(AddCartView, self).dispatch(request, *args, **kwargs)

    def form_valid(self, form):
        cd = form.cleaned_data
        self.cart.add(product=self.product,
                      quantity=cd['quantity'],
                      update_quantity=cd['update'])

        return redirect('cart:cart_detail')


class RemoveCartView(View):

    def get(self, request, *args, **kwargs):
        cart = Cart(request)
        product = get_object_or_404(Product, id=kwargs['product_id'])
        cart.remove(product)
        return redirect('cart:cart_detail')


class CartDetailView(TemplateView):
    template_name = 'cart/detail.html'

    def get_context_data(self, **kwargs):
        cart = Cart(self.request)
        for item in cart:
            item['update_quantity_form'] = CartAddProductForm(initial={
                'quantity': item['quantity'], 'update': True
            })

        coupon_apply_form = CouponApplyForm()

        context = super(CartDetailView, self).get_context_data(**kwargs)
        context['cart'] = cart
        context['coupon_apply_form'] = coupon_apply_form
        return context
