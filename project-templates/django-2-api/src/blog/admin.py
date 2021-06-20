from django.contrib import admin

from blog.models import Post, Comment


@admin.register(Post)
class PostAdmin(admin.ModelAdmin):
    # columns to display in /admin/blog/post/
    list_display = ('title', 'slug', 'author', 'publish', 'status')

    # fields to filter the list by
    list_filter = ('status', 'created', 'publish', 'author')

    # the search bar appears at the top and perform search against these fields.
    search_fields = ('title', 'body')

    # slug field will be prepopulated as we type in the title in /admin/blog/post/add/
    prepopulated_fields = {'slug': ('title',)}

    # enter id instead of selecting from dropdown menu in /admin/blog/post/add/
    raw_id_fields = ('author',)

    # adds a data-based navigation to /admin/blog/post
    date_hierarchy = 'publish'

    ordering = ['status', 'publish']


@admin.register(Comment)
class CommentAdmin(admin.ModelAdmin):
    list_display = ('name', 'email', 'post', 'created', 'active')
    list_filter = ('active', 'created', 'updated')
    search_fields = ('name', 'email', 'body')
