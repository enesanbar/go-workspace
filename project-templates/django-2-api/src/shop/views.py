from django.shortcuts import render, get_object_or_404
from django.views.generic import TemplateView, DetailView

from cart.forms import CartAddProductForm
from shop.models import Category, Product


class ProductListView(TemplateView):
    category = None
    template_name = 'shop/product/list.html'

    def dispatch(self, request, *args, **kwargs):
        if 'category_slug' in kwargs:
            self.category = get_object_or_404(Category, slug=kwargs['category_slug'])
        return super(ProductListView, self).dispatch(request, *args, **kwargs)

    def get_context_data(self, **kwargs):
        products = Product.objects.filter(available=True)
        if self.category:
            products = products.filter(category=self.category)
        context = super(ProductListView, self).get_context_data(**kwargs)
        context['category'] = self.category
        context['categories'] = Category.objects.all()
        context['products'] = products
        return context


class ProductDetailView(DetailView):
    model = Product
    template_name = 'shop/product/detail.html'
    context_object_name = 'product'

    def get_object(self, queryset=None):
        return get_object_or_404(Product,
                                 id=self.kwargs['id'],
                                 slug=self.kwargs['slug'],
                                 available=True)

    def get_context_data(self, **kwargs):
        context = super(ProductDetailView, self).get_context_data(**kwargs)
        context['cart_product_form'] = CartAddProductForm()
        return context
