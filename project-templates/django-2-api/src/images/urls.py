from django.urls import path
from django.utils.translation import gettext_lazy as _

from images.views import ImageCreateView, ImageDetailView, ImageLikeView, ImageListView

app_name = 'images'

urlpatterns = [
    path('', ImageListView.as_view(), name='list'),
    path(_('create/'), ImageCreateView.as_view(), name='create'),
    path(_('detail/<int:id>/<slug:slug>/'), ImageDetailView.as_view(), name='detail'),
    path('<int:id>/like/', ImageLikeView.as_view(), name='like'),
]
