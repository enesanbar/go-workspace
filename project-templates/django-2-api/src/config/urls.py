from django.conf import settings
from django.conf.urls import include
from django.conf.urls.i18n import i18n_patterns
from django.conf.urls.static import static
from django.contrib import admin
from django.contrib.sitemaps.views import sitemap
from django.urls import path
from django.utils.translation import gettext_lazy as _

from blog.sitemaps import PostSitemap

sitemaps = {
    'posts': PostSitemap,
}

urlpatterns = i18n_patterns(
    path('admin/', admin.site.urls),
    path(_('account/'), include('account.urls')),
    path(_('blog/'), include('blog.urls', namespace='blog')),
    path(_('images/'), include('images.urls', namespace='images')),
    path(_('shop/'), include('shop.urls', namespace='shop')),
    path(_('cart/'), include('cart.urls', namespace='cart')),
    path(_('orders/'), include('orders.urls', namespace='orders')),
    path(_('coupons/'), include('coupons.urls', namespace='coupons')),
    path(_('courses/'), include('courses.urls')),
    # url(r'^rosetta/', include('rosetta.urls')),
    path('sitemap.xml', sitemap, {'sitemaps': sitemaps}, name='django.contrib.sitemaps.views.sitemap'),
)

urlpatterns += [
    path('api/', include('courses.api.urls', namespace='api')),
]

if settings.DEBUG:
    import debug_toolbar

    urlpatterns += i18n_patterns(path('__debug__/', include(debug_toolbar.urls)))
    urlpatterns += static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)
