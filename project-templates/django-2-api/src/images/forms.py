from urllib import request

import itertools
from urllib.request import Request

from django import forms
from django.core.files.base import ContentFile
from django.utils.text import slugify

from images.models import Image


class ImageCreateForm(forms.ModelForm):

    class Meta:
        model = Image
        fields = ('title', 'url', 'description')
        widgets = {
            'url': forms.HiddenInput,
        }

    def clean_url(self):
        url = self.cleaned_data['url']
        valid_extensions = ['jpg', 'jpeg']
        extension = url.rsplit('.', 1)[1].lower()
        if extension not in valid_extensions:
            raise forms.ValidationError('The given URL does not match valid image extensions')
        return url

    def save(self, commit=True):
        image = super(ImageCreateForm, self).save(commit=False)

        image_url = self.cleaned_data['url']
        image.slug = initial = slugify(image.title)

        for x in itertools.count(1):
            if not Image.objects.filter(slug=image.slug).exists():
                break
            image.slug = '%s-%d' % (initial, x)

        image_name = '{}.{}'.format(image.slug, image_url.rsplit('.', 1)[1].lower())

        # download the image from the given url
        req = Request(image_url, headers={'User-Agent': 'Mozilla/5.0'})
        response = request.urlopen(req)

        image.image.save(image_name, ContentFile(response.read()), save=False)

        if commit:
            image.save()
        return image
