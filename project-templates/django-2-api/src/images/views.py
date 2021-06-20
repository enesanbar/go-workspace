import redis
from django.conf import settings
from django.contrib import messages
from django.http import JsonResponse
from django.shortcuts import render, redirect
from django.views import View
from django.views.decorators.http import require_http_methods
from django.views.generic import FormView, DetailView, ListView
from django.utils.decorators import method_decorator
from django.contrib.auth.decorators import login_required

from actions.utils import create_action
from images.forms import ImageCreateForm
from images.models import Image
from utils.decorators import ajax_required

# connect to redis
r = redis.StrictRedis(host=settings.REDIS_HOST, port=settings.REDIS_PORT, db=settings.REDIS_DB)


class ImageCreateView(FormView):
    form_class = ImageCreateForm
    template_name = 'images/image/create.html'

    def form_valid(self, form):
        form = self.get_form()

        new_item = form.save(commit=False)
        # assign current user to the item
        new_item.user = self.request.user
        new_item.save()
        create_action(self.request.user, 'bookmarked image', new_item)

        messages.success(self.request, 'Image added successfully')

        # redirect to new created item detail view
        return redirect(new_item.get_absolute_url())

    def get_context_data(self, **kwargs):
        context = super(ImageCreateView, self).get_context_data(**kwargs)
        context['section'] = 'images'

        if self.request.method == 'GET':
            context['form'] = self.form_class(self.request.GET)
        return context


class ImageListView(ListView):
    model = Image
    context_object_name = 'images'
    paginate_by = 8
    template_name = 'images/image/list.html'

    def get(self, request, *args, **kwargs):
        if request.is_ajax():
            self.template_name = 'images/image/list_ajax.html'
        return super(ImageListView, self).get(request, *args, **kwargs)

    def get_context_data(self, **kwargs):
        context = super(ImageListView, self).get_context_data(**kwargs)
        context['section'] = 'images'
        return context


class ImageDetailView(DetailView):
    model = Image
    template_name = 'images/image/detail.html'

    def get_context_data(self, **kwargs):
        # increment total image views by 1
        total_views = r.incr('image:{}:views'.format(self.get_object().id))
        context = super(ImageDetailView, self).get_context_data(**kwargs)
        context['section'] = 'images'
        context['total_views'] = total_views
        return context


@method_decorator(ajax_required, name='dispatch')
@method_decorator(require_http_methods(["POST"]), name='dispatch')
@method_decorator(login_required, name='dispatch')
class ImageLikeView(View):

    def post(self, request, *args, **kwargs):
        image_id = kwargs['id']
        action = request.POST.get('action')
        if image_id and action:
            try:
                image = Image.objects.get(id=image_id)
                if action == 'like':
                    image.users_like.add(request.user)
                    create_action(request.user, 'likes', image)
                else:
                    image.users_like.remove(request.user)
                return JsonResponse({'status': 'ok'})
            except:
                pass
        return JsonResponse({'status': 'ko'})
