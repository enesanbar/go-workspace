from django.urls import path

from blog.feeds import LatestPostsFeed
from blog.views import PostListView, PostDetailView, PostShareView

app_name = 'blog'

urlpatterns = [
    # Accepts both /blog and /blog/tag/<tag_slug>/
    path('', PostListView.as_view(), name='post_list'),
    path('tag/<slug:tag_slug>/', PostListView.as_view(), name='post_list_by_tag'),
    path('<int:year>/<int:month>/<int:day>/<slug:slug>/', PostDetailView.as_view(), name='post_detail'),
    path('<int:post_id>/share/', PostShareView.as_view(), name='post_share'),
    path('feed/', LatestPostsFeed(), name='post_feed'),
]
