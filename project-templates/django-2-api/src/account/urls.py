from django.conf.urls import url
from django.contrib.auth.views import LogoutView, password_change, password_change_done, password_reset, \
    password_reset_done, password_reset_confirm, password_reset_complete
from django.urls import path
from django.utils.translation import gettext_lazy as _

from account.views import DashboardView, SignupView, EditView, UserListView, UserDetailView, CustomLoginView, \
    UserFollowView

urlpatterns = [
    path(_('users/'), UserListView.as_view(), name='user_list'),
    path(_('users/follow/'), UserFollowView.as_view(), name='user_follow'),
    path(_('users/<slug:username>/'), UserDetailView.as_view(), name='user_detail'),

    path(_('register/'), SignupView.as_view(), name='register'),
    path(_('login/'), CustomLoginView.as_view(), name='login'),
    path(_('logout/'), LogoutView.as_view(), name='logout'),
    path(_('dashboard/'), DashboardView.as_view(), name='dashboard'),
    path(_('edit/'), EditView.as_view(), name='edit'),

    # change password urls
    path(_('password-change/'), password_change, name='password_change'),
    path(_('password-change/done/'), password_change_done, name='password_change_done'),

    # restore password urls
    path(_('password-reset/'), password_reset, name='password_reset'),
    path(_('password-reset/done/'), password_reset_done, name='password_reset_done'),
    path(_('password-reset/confirm/<slug:uidb64>/<slug:token>/'), password_reset_confirm, name='password_reset_confirm'),
    path(_('password-reset/complete/'), password_reset_complete, name='password_reset_complete')
]
