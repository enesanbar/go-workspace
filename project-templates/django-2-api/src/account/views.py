from django.contrib import messages
from django.contrib.auth.decorators import login_required
from django.contrib.auth.forms import UserCreationForm
from django.contrib.auth.models import User
from django.contrib.auth.views import LoginView
from django.http import JsonResponse
from django.shortcuts import render, redirect
from django.utils.decorators import method_decorator
from django.views import View
from django.views.decorators.http import require_http_methods
from django.views.generic import TemplateView, FormView, ListView, DetailView, CreateView

from account.forms import ProfileEditForm, UserEditForm
from account.models import Profile, Contact
from actions.models import Action
from actions.utils import create_action
from utils.decorators import anonymous_required, ajax_required


@method_decorator(anonymous_required, name='dispatch')
class CustomLoginView(LoginView):
    def dispatch(self, request, *args, **kwargs):
        if self.request.user.is_authenticated:
            messages.error(self.request, "You are already logged in!")
        return super(CustomLoginView, self).dispatch(request, *args, **kwargs)

    def form_valid(self, form):
        messages.success(self.request, "You have been successfully logged in.")
        return super(CustomLoginView, self).form_valid(form)


@method_decorator(login_required, name='dispatch')
class DashboardView(TemplateView):
    template_name = 'account/dashboard.html'

    def get_context_data(self, **kwargs):
        # Display all actions by default
        actions = Action.objects.exclude(user=self.request.user)
        following_ids = self.request.user.following.values_list('id', flat=True)

        # If user is following others, retrieve only their actions
        if following_ids:
            actions = actions.filter(user_id__in=following_ids)\
                            .select_related('user', 'user__profile')\
                            .prefetch_related('target')

        actions = actions[:10]

        context = super(DashboardView, self).get_context_data(**kwargs)
        context['section'] = 'dashboard'
        context['actions'] = actions
        return context


@method_decorator(anonymous_required, name='dispatch')
class SignupView(CreateView):
    form_class = UserCreationForm
    template_name = 'registration/register.html'

    def form_valid(self, form):
        # Create a new user object but avoid saving it yet
        new_user = self.get_form().save()

        # Create the user profile
        Profile.objects.create(user=new_user)

        create_action(new_user, 'has created an account')
        return render(self.request, 'registration/register_done.html', {'new_user': new_user})


@method_decorator(login_required, name='dispatch')
class EditView(View):
    template_name = 'account/edit.html'
    user_form = None
    profile_form = None

    def get(self, request):
        self.user_form = UserEditForm(instance=request.user)
        self.profile_form = ProfileEditForm(instance=request.user.profile)
        return self.render(request)

    def post(self, request):
        self.user_form = UserEditForm(instance=request.user, data=request.POST)
        self.profile_form = ProfileEditForm(instance=request.user.profile,
                                            data=request.POST, files=request.FILES)

        if self.user_form.is_valid() and self.profile_form.is_valid():
            self.user_form.save()
            self.profile_form.save()
            messages.success(request, 'Your profile has been updated successfully')
            return redirect('dashboard')
        else:
            messages.error(request, 'An error occurred updating your profile')
            return self.render(request)

    def render(self, request):
        return render(request, 'account/edit.html', {
            'user_form': self.user_form,
            'profile_form': self.profile_form
        })


class UserListView(ListView):
    queryset = User.objects.filter(is_active=True)
    context_object_name = 'users'
    template_name = 'account/user/list.html'

    def get_context_data(self, **kwargs):
        context = super(UserListView, self).get_context_data(**kwargs)
        context['section'] = 'people'
        return context


class UserDetailView(DetailView):
    template_name = 'account/user/detail.html'
    context_object_name = 'user'

    def get_object(self, queryset=None):
        return User.objects.get(username=self.kwargs['username'], is_active=True)

    def get_context_data(self, **kwargs):
        context = super(UserDetailView, self).get_context_data(**kwargs)
        context['section'] = 'people'
        return context


@method_decorator(ajax_required, name='dispatch')
@method_decorator(require_http_methods(["POST"]), name='dispatch')
@method_decorator(login_required, name='dispatch')
class UserFollowView(View):

    def post(self, request, *args, **kwargs):
        user_id = request.POST.get('id')
        action = request.POST.get('action')

        if user_id and action:
            try:
                user = User.objects.get(id=user_id)
                if action == 'follow':
                    Contact.objects.get_or_create(user_from=request.user, user_to=user)
                    create_action(request.user, 'is following', user)
                else:
                    Contact.objects.filter(user_from=request.user, user_to=user).delete()
                return JsonResponse({'status': 'ok'})

            except User.DoesNotExist:
                return JsonResponse({'status': 'ko'})

        return JsonResponse({'status': 'ko'})
