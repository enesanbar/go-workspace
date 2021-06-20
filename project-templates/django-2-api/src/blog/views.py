from django.core.mail import send_mail
from django.db.models import Count
from django.shortcuts import render, get_object_or_404
from django.views import View
from django.views.generic import ListView, DetailView
from taggit.models import Tag

from blog.forms import EmailForm, CommentForm
from blog.models import Post


class PostListView(ListView):
    context_object_name = 'posts'
    template_name = 'blog/post/list.html'
    queryset = Post.published.all()
    paginate_by = 3
    tag = None

    def get_queryset(self):
        queryset = self.queryset
        if 'tag_slug' in self.kwargs and self.kwargs['tag_slug']:
            self.tag = get_object_or_404(Tag, slug=self.kwargs['tag_slug'])
            queryset = self.queryset.filter(tags__in=[self.tag])

        return queryset

    def get_context_data(self, **kwargs):
        context = super(PostListView, self).get_context_data(**kwargs)
        context['tag'] = self.tag
        return context


class PostDetailView(DetailView):
    context_object_name = 'post'
    template_name = 'blog/post/detail.html'
    comment_form = CommentForm
    object = None
    comments = None

    def get_object(self, queryset=None):
        return Post.published.get(slug=self.kwargs['slug'],
                                  publish__year=self.kwargs['year'],
                                  publish__month=self.kwargs['month'],
                                  publish__day=self.kwargs['day'])

    def dispatch(self, request, *args, **kwargs):
        # List of active comments for this post
        self.object = self.get_object()
        self.comments = self.object.comments.filter(active=True)
        return super(PostDetailView, self).dispatch(request, *args, **kwargs)

    def get(self, request, *args, **kwargs):
        comment_form = self.comment_form()
        return self.send_view(request, comment_form)

    def post(self, request, *args, **kwargs):
        comment_form = self.comment_form(request.POST)

        if comment_form.is_valid():
            new_comment = comment_form.save(commit=False)
            # Assign the current post to the comment
            new_comment.post = self.object
            # Save the comment to the database
            new_comment.save()

        return self.send_view(request, comment_form)

    def get_context_data(self, **kwargs):
        post_tags_ids = self.object.tags.values_list('id', flat=True)
        similar_posts = Post.published.filter(tags__in=post_tags_ids).exclude(id=self.object.id)
        similar_posts = similar_posts.annotate(same_tags=Count('tags')).order_by('-same_tags', '-publish')[:4]

        context = super(PostDetailView, self).get_context_data(object=self.object, **kwargs)
        context['similar_posts'] = similar_posts
        context['comments'] = self.comments
        return context

    def send_view(self, request, comment_form):
        context = self.get_context_data()
        context['comment_form'] = comment_form
        return render(request, 'blog/post/detail.html', context)


class PostShareView(View):
    form_class = EmailForm
    template_name = 'blog/post/share.html'
    blog_post = None
    sent = False

    def dispatch(self, request, *args, **kwargs):
        self.blog_post = get_object_or_404(Post, id=kwargs['post_id'], status='published')
        return super(PostShareView, self).dispatch(request, *args, **kwargs)

    def get(self, request, *args, **kwargs):
        form = self.form_class()
        return render(request, self.template_name, {'form': form, 'post': self.blog_post, 'sent': self.sent})

    def post(self, request, *args, **kwargs):
        form = self.form_class(request.POST)
        if form.is_valid():
            cd = form.cleaned_data

            post_url = request.build_absolute_uri(self.blog_post.get_absolute_url())
            subject = '{} ({}) recommends you reading "{}"'.format(cd['name'], cd['email'], self.blog_post.title)
            message = 'Read "{}" at {}\n\n{}\'s comments: {}'.format(self.blog_post.title, post_url, cd['name'], cd['comments'])
            send_mail(subject, message, 'admin@myblog.com', [cd['to']])
            self.sent = True
        return render(request, self.template_name, {'form': form, 'post': self.blog_post, 'sent': self.sent})
