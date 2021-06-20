from django.http import HttpResponseRedirect, HttpResponseBadRequest
from django.urls import reverse_lazy


def anonymous_required(view_function, redirect_to=None):
    return AnonymousRequired(view_function, redirect_to)


class AnonymousRequired(object):
    def __init__(self, view_function, redirect_to):
        self.view_function = view_function
        self.redirect_to = redirect_to

    def __call__(self, request, *args, **kwargs):
        if request.user is not None and request.user.is_authenticated:
            if self.redirect_to is None:
                self.redirect_to = reverse_lazy('dashboard')
            return HttpResponseRedirect(self.redirect_to)
        return self.view_function(request, *args, **kwargs)


def ajax_required(f):
    def wrap(request, *args, **kwargs):
        if not request.is_ajax():
            return HttpResponseBadRequest()
        return f(request, *args, **kwargs)

    wrap.__doc__ = f.__doc__
    wrap.__name__ = f.__name__
    return wrap
