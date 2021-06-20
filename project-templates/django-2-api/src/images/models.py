from django.contrib.auth.models import User
from django.db import models
from django.urls import reverse


class Image(models.Model):
    user = models.ForeignKey(User, related_name='images_created', on_delete=models.CASCADE)
    title = models.CharField(max_length=200)
    slug = models.SlugField(max_length=200, blank=True)
    url = models.URLField()
    image = models.ImageField(upload_to='images/%Y/%m/%d')
    description = models.TextField(blank=True)
    created = models.DateField(auto_now_add=True, db_index=True)
    total_likes = models.PositiveIntegerField(db_index=True, default=0)

    users_like = models.ManyToManyField(User, related_name='images_liked', blank=True)

    def get_absolute_url(self):
        return reverse('images:detail', args=[self.id, self.slug])

    def __str__(self):
        return self.title