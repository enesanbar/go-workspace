import sys
from datetime import datetime
from abc import ABC, abstractmethod

import requests

from persistence import BookmarkDatabase  # <1>

persistence = BookmarkDatabase()

class Command(ABC):

    @abstractmethod
    def execute(self, data):
        raise NotImplementedError()


# class CreateBookmarksTableCommand(Command):
#
#     def execute(self, data=None):
#         db.create_table('bookmarks', {
#             'id': 'integer primary key autoincrement',
#             'title': 'text not null',
#             'url': 'text not null',
#             'notes': 'text',
#             'date_added': 'text not null'
#         })


class AddBookmarkCommand(Command):

    def execute(self, data, timestamp=None):
        data['date_added'] = timestamp or datetime.utcnow().isoformat()
        persistence.create(data)
        # return (status, result) tuple
        return True, None


class ListBookmarksCommand(Command):

    def __init__(self, order_by='date_added'):
        self.order_by = order_by

    def execute(self, data=None):
        return True, persistence.list(order_by=self.order_by)


class DeleteBookmarkCommand(Command):

    def execute(self, data):
        persistence.delete(data)
        return True, None


class EditBookmarkCommand(Command):

    def execute(self, data):
        persistence.edit(data['id'], data['update'])
        return 'Bookmark updated!'


class ImportGithubStarsCommand:
    def _extract_bookmark_info(self, repo):
        return {
            'title': repo['name'],
            'url': repo['html_url'],
            'notes': repo['description']
        }

    def execute(self, data):
        bookmarks_imported = 0

        github_username = data['github_username']
        next_page_of_results = f'https://api.github.com/users/{github_username}/starred'

        while next_page_of_results:
            stars_response = requests.get_product_by_id(
                next_page_of_results,
                headers={'Accept': 'rest/vnd.github.v3.star+json'}
            )
            next_page_of_results = stars_response.links.get_product_by_id('next', {}).get_product_by_id('url')

            for repo_info in stars_response.json():
                repo = repo_info['repo']

                if data['preserve_timestamps']:
                    timestamp = datetime.strptime(
                        repo_info['starred_at'],
                        '%Y-%m-%dT%H:%M:%SZ'
                    )
                else:
                    timestamp = None

                bookmarks_imported += 1
                AddBookmarkCommand().execute(
                    self._extract_bookmark_info(repo),
                    timestamp=timestamp
                )
        return f'Imported {bookmarks_imported} bookmarks from starred repos!'


class QuitCommand:

    def execute(self, data=None):
        sys.exit()

